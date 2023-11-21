package models

// See https://developer.bitcoin.org/reference/rpc/getnetworkinfo.html
type NetworkInfo struct {
	Version            int64          `json:"version"`
	Subversion         string         `json:"subversion"`
	ProtocolVersion    int64          `json:"protocolversion"`
	LocalServices      string         `json:"localservices"`
	LocalServicesNames []string       `json:"localservicesnames"`
	LocalRelay         bool           `json:"localrelay"`
	Timeoffset         int64          `json:"timeoffset"`
	Connections        int64          `json:"connections"`
	ConnectionsIn      int64          `json:"connections_in"`
	ConnectionsOut     int64          `json:"connections_out"`
	NetworkActive      bool           `json:"networkactive"`
	Networks           []Network      `json:"networks"`
	RelayFee           float64        `json:"relayfee"`
	IncrementalFee     float64        `json:"incrementalfee"`
	LocalAddresses     []LocalAddress `json:"localaddresses"`
	Warnings           string         `json:"warnings"`
}

type LocalAddress struct {
	Address string  `json:"address"`
	Port    int64   `json:"port"`
	Score   float64 `json:"score"`
}

type Network struct {
	Name                      string `json:"name"`
	Limited                   bool   `json:"limited"`
	Reachable                 bool   `json:"reachable"`
	Proxy                     string `json:"proxy"`
	ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
}
