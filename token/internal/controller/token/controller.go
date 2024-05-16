package token

import (
	"context"
	"errors"

	"github.com/jakobsym/sol_service/token/internal/repository"
	model "github.com/jakobsym/sol_service/token/pkg"
)

var ErrNotFound = errors.New("not found")

type tokenRepo interface {
	Get(ctx context.Context, address string) (*model.Token, error)
}

type Controller struct {
	repo tokenRepo
}

func New(repo tokenRepo) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Get(ctx context.Context, address string) (*model.Token, error) {
	res, err := c.repo.Get(ctx, address)
	// check if error matches error from repository call
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
