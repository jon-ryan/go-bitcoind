package bitcoind

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// A rpcClient represents a JSON RPC client (over HTTP(s)).
type rpcClient struct {
	serverAddr string
	user       string
	passwd     string
	httpClient *http.Client
}

// rpcRequest represent a RCP request
type rpcRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
}

// RPCErrorCode represents an error code to be used as a part of an RPCError
// which is in turn used in a JSON-RPC Response object.
//
// A specific type is used to help ensure the wrong errors aren't used.
type RPCErrorCode int

// RPCError represents an error that is used as a part of a JSON-RPC Response
// object.
type RPCError struct {
	Code    RPCErrorCode `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

// Guarantee RPCError satisfies the builtin error interface.
var _, _ error = RPCError{}, (*RPCError)(nil)

// Error returns a string describing the RPC error.  This satisfies the
// builtin error interface.
func (e RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type rpcResponse struct {
	Id     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
	Err    *RPCError       `json:"error"`
}

func newClient(host string, port int, user, passwd string, useSSL bool, timeout int) (c *rpcClient, err error) {
	if len(host) == 0 {
		err = errors.New("bad call missing argument host")
		return
	}
	var urlHttpPrefix string
	var httpClient *http.Client
	if useSSL {
		urlHttpPrefix = "https://"
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{
			Transport: t,
			Timeout:   time.Duration(timeout) * time.Second,
		}
	} else {
		urlHttpPrefix = "http://"
		httpClient = &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}
	}
	c = &rpcClient{
		serverAddr: fmt.Sprintf("%s%s:%d", urlHttpPrefix, host, port),
		user:       user,
		passwd:     passwd,
		httpClient: httpClient,
	}
	return
}

// call prepare & exec the request
func (c *rpcClient) call(method string, params interface{}) (rr rpcResponse, err error) {
	rpcR := rpcRequest{method, params, time.Now().UnixNano(), "1.0"}
	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	err = jsonEncoder.Encode(rpcR)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	// Auth ?
	if len(c.user) > 0 || len(c.passwd) > 0 {
		req.SetBasicAuth(c.user, c.passwd)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &rr)
	return
}
