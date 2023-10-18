package models

type MiningInfo struct {
	Blocks                     uint64  `json:"blocks"`
	Difficulty                 float64 `json:"difficulty"`
	NetworkHashPerSecondString string  `json:"networkhashps"`
	NetworkHashPerSecond       int64
	PooledtTx                  uint64 `json:"pooledtx"`
	Chain                      string `json:"chain"`
	Warnings                   string `json:"warnings"`
}
