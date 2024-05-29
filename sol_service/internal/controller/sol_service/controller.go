package sol_service

import (
	"context"
	"errors"

	"github.com/jakobsym/sol_service/sol_service/internal/gateway"
	"github.com/jakobsym/sol_service/sol_service/pkg/model"
	tokenmodel "github.com/jakobsym/sol_service/token/pkg"
)

var ErrNotFound = errors.New("wallet/token data not found")

// make req at specific endpoint recieve json and store into a Wallet model
type walletGateway interface {
	Get(ctx context.Context, address string) (*model.Wallet, error)
}
type tokenGateway interface {
	Get(ctx context.Context, address tokenmodel.TokenAddress) (*tokenmodel.TokenDetails, error)
}

type Controller struct {
	walletGateway walletGateway
	tokenGateway  tokenGateway
}

func New(walletGateway walletGateway) *Controller {
	return &Controller{walletGateway: walletGateway}
}

func (c *Controller) Get(ctx context.Context, address string) (*model.Wallet, error) {
	walletdata, err := c.walletGateway.Get(ctx, address) // sends GET request to wallet service
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	wallet := &model.Wallet{Tokens: walletdata.Tokens, WalletAddress: address}
	// call get tokendetails here, passing a Wallet
	// the token service willl iterate over the wallet contents assiging information to each token
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
