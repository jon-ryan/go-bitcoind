package models

// See https://developer.bitcoin.org/reference/rpc/getmempoolinfo.html
type MemPoolInfo struct {
	Loaded           bool    `json:"loaded"`
	Size             uint64  `json:"size"`
	Bytes            uint64  `json:"bytes"`
	Usage            uint64  `json:"usage"`
	TotalFee         float64 `json:"total_fee"`
	MaxMemPool       uint64  `json:"maxmempool"`
	MemPoolMinFee    float64 `json:"mempoolminfee"`
	MinRelayTxFee    float64 `json:"minrelaytxfee"`
	UnbroadcastCount uint64  `json:"unbroadcastcount"`
}
