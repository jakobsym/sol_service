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

func (g *Gateway) Get(ctx context.Context, address string) (*model.TokenDetails, error) {
	return nil, nil
}
