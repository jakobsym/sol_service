package model

type Wallet struct {
	WalletAddress string           `json:"walletAddress"`
	Tokens        map[string]int64 `json:"tokens"`
}

type WalletFormatted struct {
	WalletAddress string `json:"walletAddress"`
	Tokens        map[TokenDetails]int64
}

// type TokenAddress string
type APIResponse struct {
	Data map[string]TokenDetails `json:"data"`
}

type TokenDetails struct {
	Symbol        string  `json:"mintSymbol"`
	TokenAddress  string  `json:"id"`
	VsTokenSymbol string  `json:"vsTokenSymbol,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Marketcap     string  `json:"marketcap,omitempty"`
	Supply        string  `json:"supply,omitempty"`
}
