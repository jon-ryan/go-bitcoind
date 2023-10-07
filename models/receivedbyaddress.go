package models

type ReceivedByAddress struct {
	InvolvesWatchOnly bool     `json:"involvesWatchonly"`
	Address           string   `json:"address"`
	Amount            float64  `json:"amount"`
	Confirmations     uint32   `json:"confirmations"`
	Label             string   `json:"label"`
	TxIds             []string `json:"txids"`
}
