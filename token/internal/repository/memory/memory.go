package memory

import (
	"context"
	"sync"

	"github.com/jakobsym/sol_service/token/internal/repository"
	model "github.com/jakobsym/sol_service/token/pkg"
)

type Repository struct {
	sync.RWMutex
	data map[model.TokenAddress]*model.TokenDetails
}

func New() *Repository {
	return &Repository{data: map[model.TokenAddress]*model.TokenDetails{}}
}

// get a token via address
func (r *Repository) Get(_ context.Context, address model.TokenAddress) (*model.TokenDetails, error) {
	r.RLock()
	defer r.RUnlock()
	token, err := r.data[address]
	if !err {
		return nil, repository.ErrNotFound
	}
	return token, nil
}

// add a token via address
func (r *Repository) Put(_ context.Context, address model.TokenAddress, token *model.TokenDetails) error {
	r.RLock()
	defer r.RUnlock()
	r.data[address] = token
	return nil
}
