package models

// See https://developer.bitcoin.org/reference/rpc/getnetworkinfo.html
type NetworkInfo struct {
	Version            uint64         `json:"version"`
	Subversion         string         `json:"subversion"`
	ProtocolVersion    uint64         `json:"protocolversion"`
	LocalServices      string         `json:"localservices"`
	LocalServicesNames []string       `json:"localservicesnames"`
	LocalRelay         bool           `json:"localrelay"`
	Timeoffset         uint64         `json:"timeoffset"`
	Connections        uint64         `json:"connections"`
	ConnectionsIn      uint64         `json:"connections_in"`
	ConnectionsOut     uint64         `json:"connections_out"`
	NetworkActive      bool           `json:"networkactive"`
	Networks           []Network      `json:"networks"`
	RelayFee           float64        `json:"relayfee"`
	IncrementalFee     float64        `json:"incrementalfee"`
	LocalAddresses     []LocalAddress `json:"localaddresses"`
	Warnings           string         `json:"warnings"`
}

type LocalAddress struct {
	Address string  `json:"address"`
	Port    uint64  `json:"port"`
	Score   float64 `json:"score"`
}

type Network struct {
	Name                      string `json:"name"`
	Limited                   bool   `json:"limited"`
	Reachable                 bool   `json:"reachable"`
	Proxy                     string `json:"proxy"`
	ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
}
