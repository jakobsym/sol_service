package memory

import (
	"context"
	"sync"

	"github.com/jakobsym/sol_service/token/internal/repository"
	model "github.com/jakobsym/sol_service/token/pkg"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.TokenDetails
}

func New() *Repository {
	return &Repository{data: map[string]*model.TokenDetails{}}
}

// get a token via address
func (r *Repository) Get(_ context.Context, address string) (*model.TokenDetails, error) {
	r.RLock()
	defer r.RUnlock()
	token, err := r.data[address]
	if !err {
		return nil, repository.ErrNotFound
	}
	return token, nil
}

// add a token via address
func (r *Repository) Put(_ context.Context, address string, token *model.TokenDetails) error {
	r.RLock()
	defer r.RUnlock()
	r.data[address] = token
	return nil
}
