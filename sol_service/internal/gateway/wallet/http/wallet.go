package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jakobsym/sol_service/sol_service/internal/gateway"
	model "github.com/jakobsym/sol_service/sol_service/pkg/model"
)

type Gateway struct {
	addr string
}

func New(addr string) *Gateway {
	return &Gateway{addr: addr}
}

// get wallet content by wallet address
func (g *Gateway) Get(ctx context.Context, address string) (*model.Wallet, error) {
	fAddress := "http://" + g.addr
	req, err := http.NewRequest(http.MethodGet, fAddress+"/wallet", nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", address)
	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non: 2XX error: %v", err)
	}

	var w *model.Wallet
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return nil, err
	}
	return w, nil
}
