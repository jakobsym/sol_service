package model

type Wallet struct {
	WalletAddress string           `json:"walletAddress"`
	Tokens        map[string]int64 `json:"tokens"`
}
