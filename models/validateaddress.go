package models

type ValidateAddress struct {
	IsValid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	ScriptPubKey string `json:"scriptPubKey"`
	IsScript     bool   `json:"isscript"`
	IsWitness    bool   `json:"iswitness"`
}
