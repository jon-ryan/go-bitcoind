package models

// See https://developer.bitcoin.org/reference/rpc/getblockchaininfo.html
type BlockChainInfo struct {
	Chain                string                 `json:"chain"`
	Blocks               uint64                 `json:"blocks"`
	Headers              uint64                 `json:"headers"`
	BestBlockHash        string                 `json:"bestblockhash"`
	Difficulty           float64                `json:"difficulty"`
	Time                 uint64                 `json:"time"`
	Mediantime           uint64                 `json:"mediantime"`
	VerificationProgress float64                `json:"verificationprogress"`
	InitialBlockDownload bool                   `json:"initialblockdownload"`
	Chainwork            string                 `json:"chainwork"`
	SizeOnDisk           uint64                 `json:"size_on_disk"`
	Pruned               bool                   `json:"pruned"`
	PruneHeight          float64                `json:"pruneheight"`
	AutomaticPruning     bool                   `json:"automatic_pruning"`
	PruneTargetSize      float64                `json:"prune_target_size"`
	Softforks            map[string]interface{} `json:"softforks"`
	Warnings             string                 `json:"warnings"`
}
