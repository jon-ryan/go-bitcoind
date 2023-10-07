package models

type VerboseTx struct {
	VSize            uint64   `json:"vsize"`
	Time             uint64   `json:"time"`
	Height           uint64   `json:"height"`
	DescendantCount  uint64   `json:"descendantcount"`
	DescendantSize   uint64   `json:"descendantsize"`
	AncestorCount    uint64   `json:"ancestorcount"`
	AncestorSize     uint64   `json:"ancestorsize"`
	AncestorFees     uint64   `json:"ancestorfees"`
	WTxId            string   `json:"wtxid"`
	Fees             Fees     `json:"fees"`
	Depends          []string `json:"depends"`
	SpentBy          []string `json:"spentby"`
	Bip125Replacable bool     `json:"bip125-replaceable"`
	Unbroadcast      bool     `json:"unbroadcast"`
}

type Fees struct {
	Base      float64 `json:"base"`
	Modified  float64 `json:"modified"`
	Ancestor  float64 `json:"ancestor"`
	Decendant float64 `json:"decendant"`
}
