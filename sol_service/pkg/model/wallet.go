package model

type Wallet struct {
	WalletAddress string           `json:"walletAddress"`
	Tokens        map[string]int64 `json:"tokens"`
}

type WalletFormatted struct {
	WalletAddress string `json:"walletAddress"`
}

// Tokens may need to be map[string]string for handling response
