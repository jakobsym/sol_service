package token

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/jakobsym/sol_service/token/internal/repository"
	model "github.com/jakobsym/sol_service/token/pkg"
)

var ErrNotFound = errors.New("not found")
var Url = "https://price.jup.ag/v4/price"

type tokenRepo interface {
	Get(ctx context.Context, address model.TokenAddress) (*model.TokenDetails, error)
}

type Controller struct {
	repo tokenRepo
}

func New(repo tokenRepo) *Controller {
	return &Controller{repo: repo}
}

// TODO:
// - return all token information
// symbol, price, mc, supply, (unsure)holders?

func (c *Controller) Get(ctx context.Context, address model.TokenAddress) (*model.TokenDetails, error) {
	res, err := c.repo.Get(ctx, address) // check if token in database
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		res, err = c.Getjup(ctx, address)
		if err != nil {
			return nil, ErrNotFound
		}
		return res, fmt.Errorf("error getting token from mem./api req: %v\n", err)
	}
	return res, err
}

// calls Jupiter API to obtain token details
func (c *Controller) Getjup(ctx context.Context, address model.TokenAddress) (*model.TokenDetails, error) {
	var res *model.APIResponse

	req, err := http.NewRequest(http.MethodGet, Url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(address))
	req.URL.RawQuery = values.Encode()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// error check response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Token not found\n")
	}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	tokenDetails := &model.TokenDetails{
		Symbol:       res.Data[address].Symbol,
		TokenAddress: res.Data[address].TokenAddress,
		Price:        float64(res.Data[address].Price),
	}

	md, err := c.TokenSupply(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("error obtaining token metadata: %v\n", err)
	}
	// Convert supply/mc as float64 to string
	tokenDetails.Supply = FormatFloat(md.Supply)
	tokenDetails.Marketcap = FormatFloat(md.Supply * tokenDetails.Price)

	return tokenDetails, nil
}

func (c *Controller) TokenSupply(ctx context.Context, address model.TokenAddress) (*model.TokenMetaData, error) {
	var tokenMetadata = &model.TokenMetaData{}
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	pubKey := solana.MustPublicKeyFromBase58(string(address))
	supply, err := client.GetTokenSupply(ctx, pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("error obtaining token supply: %v\n", err)
	}

	res, err := strconv.ParseFloat(supply.Value.Amount, 64) // (res / 1e9)
	if err != nil {
		return nil, fmt.Errorf("error converting type: %v\n", err)
	}
	tokenMetadata.Supply = (res / 1e9)

	return tokenMetadata, nil
}

// FormatFloat converts a given float64 into a string
func FormatFloat(f float64) string {
	ff := strconv.FormatFloat(f, 'f', -1, 64)
	return ff
}
