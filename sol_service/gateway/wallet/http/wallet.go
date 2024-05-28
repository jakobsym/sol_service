package http

import (
	"context"

	model "github.com/jakobsym/sol_service/token/pkg"
)

type Gateway struct {
	addr string
}

func New(addr string) *Gateway {
	return &Gateway{addr: addr}
}

// get wallet content
func (g *Gateway) Get(ctx context.Context, address string) (*model.Wallet, error) {
	return nil, nil
}