package model

type TokenAddress string

type APIResponse struct {
	Data map[TokenAddress]TokenDetails `json:"data"`
}

// type TokenSymbol string
// type TokenSupply float64
type TokenDetails struct {
	Symbol        string       `json:"mintSymbol"`
	TokenAddress  TokenAddress `json:"id"`
	VsTokenSymbol string       `json:"VsTokenSymbol,omitempty"`
	Price         float64      `json:"price,omitempty"`
	Marketcap     string       `json:"marketcap,omitempty"`
	Supply        string       `json:"supply,omitempty"`
}
