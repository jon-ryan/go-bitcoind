package models

// See https://developer.bitcoin.org/reference/rpc/getwalletinfo.html
type WalletInfo struct {
	WalletName            string   `json:"walletname"`
	WalletVersion         uint64   `json:"walletversion"`
	Format                string   `json:"format"`
	TxCount               uint64   `json:"txcount"`
	KeyPoolOldest         uint64   `json:"keypoololdest"`
	KeyPoolSize           uint64   `json:"keypoolsize"`
	KeyPoolSizeHdInternal uint64   `json:"keypoolsize_hd_internal"`
	UnlockedUntil         uint64   `json:"unlocked_until"`
	PayTxFee              float64  `json:"paytxfee"`
	HdSeedId              string   `json:"hdseedid"`
	PrivateKeysEnabled    bool     `json:"private_keys_enabled"`
	AvoidReuse            bool     `json:"avoid_reuse"`
	Scanning              Scanning `json:"scanning"`
	Descriptors           bool     `json:"descriptors"`
}

type Scanning struct {
	Duration uint64  `json:"duration"`
	Progress float64 `json:"progress"`
}
