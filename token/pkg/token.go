package model

type TokenAddress string

type Token struct {
	TokenAddress  TokenAddress `json:"id"`
	MintSymbol    string       `json:"mintSymbol"`
	VsTokenSymbol string       `json:"VsTokenSymbol,omitempty"`
	Price         float64      `json:"price,omitempty"`
}

// type TokenSymbol string
// type TokenSupply float64
type TokenDetails struct {
	Symbol    string
	Price     float64
	Marketcap float64
	Supply    string
}
