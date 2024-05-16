package model

type Token struct {
	CoinAddress   string  `json:"id"`
	MintSymbol    string  `json:"mintSymbol"`
	VsTokenSymbol string  `json:"VsTokenSymbol,omitempty"`
	Price         float64 `json:"price,omitempty"`
}
