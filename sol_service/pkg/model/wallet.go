package model

type Wallet struct {
	WalletAddress string           `json:"walletAddress"`
	Tokens        map[string]int64 `json:"tokens"`
}

// Tokens may need to be map[string]string for handling response
