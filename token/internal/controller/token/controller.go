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
	model "github.com/jakobsym/sol_service/token/pkg"
)

var ErrNotFound = errors.New("not found")
var Url = "https://price.jup.ag/v4/price"

type tokenRepo interface {
	Put(ctx context.Context, address model.TokenAddress, token *model.TokenDetails) error
	Get(ctx context.Context, address model.TokenAddress) (*model.TokenDetails, error)
}

type Controller struct {
	repo tokenRepo
}

func New(repo tokenRepo) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Put(ctx context.Context, address model.TokenAddress, token *model.TokenDetails) error {
	return c.repo.Put(ctx, address, token)
}

func (c *Controller) Get(ctx context.Context, address model.TokenAddress) (*model.TokenDetails, error) {
	res, err := c.Getjup(ctx, address)
	if err != nil {
		return nil, ErrNotFound
	}
	return res, err
	/*
		res, err := c.repo.Get(ctx, address) // check if token in database
		if err != nil && errors.Is(err, repository.ErrNotFound) {
			res, err = c.Getjup(ctx, address)
			if err != nil {
				return nil, ErrNotFound
			}
			return res, fmt.Errorf("error getting token from mem./api req: %v\n", err)
		}
		return res, err
	*/
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
	values.Add("ids", string(address))
	req.URL.RawQuery = values.Encode()

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	defer response.Body.Close()

	// TODO: this doesnt work? 400 response and wasnt being invoked??
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

	s, err := c.TokenSupply(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("error obtaining token metadata: %v\n", err)
	}
	// Convert supply/mc as float64 to string
	tokenDetails.Supply = FormatFloat(s)
	tokenDetails.Marketcap = FormatFloat(s * tokenDetails.Price)

	return tokenDetails, nil
}

// TokenSupply() returns supply as a float of a given token
// in order to commpute marketcap of given token
func (c *Controller) TokenSupply(ctx context.Context, address model.TokenAddress) (float64, error) {

	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	pubKey := solana.MustPublicKeyFromBase58(string(address))
	supply, err := client.GetTokenSupply(ctx, pubKey, rpc.CommitmentFinalized)

	if err != nil {
		return 0, fmt.Errorf("error obtaining token supply: %v\n", err)
	}

	res, err := strconv.ParseFloat(supply.Value.UiAmountString, 64) // (res / 1e9)
	if err != nil {
		return 0, fmt.Errorf("error converting type: %v\n", err)
	}

	return res, nil
}

// FormatFloat converts a given float64 into a string
func FormatFloat(f float64) string {
	ff := strconv.FormatFloat(f, 'f', -1, 64)
	return ff
}
