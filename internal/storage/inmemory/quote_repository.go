package inmemory

import (
	"context"
	"sync"

	"github.com/willbicks/charisms/internal/model"
	"github.com/willbicks/charisms/internal/service"
	storage "github.com/willbicks/charisms/internal/storage/common"
)

type QuoteRepository struct {
	mu sync.RWMutex
	m  map[string]model.Quote
}

func NewQuoteRepository() service.QuoteRepository {
	return &QuoteRepository{
		m: make(map[string]model.Quote, 0),
	}
}

func (r *QuoteRepository) Create(ctx context.Context, q model.Quote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[q.ID] = q
	return nil
}

func (r *QuoteRepository) Update(ctx context.Context, q model.Quote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.m[q.ID]

	if !ok {
		return storage.ErrNotFound
	}

	r.m[q.ID] = q
	return nil
}

func (r *QuoteRepository) FindByID(ctx context.Context, id string) (model.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q, ok := r.m[id]
	if !ok {
		return model.Quote{}, storage.ErrNotFound
	}

	return q, nil
}

func (r *QuoteRepository) FindAll(ctx context.Context) ([]model.Quote, error) {
	v := make([]model.Quote, 0, len(r.m))

	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, value := range r.m {
		v = append(v, value)
	}

	return v, nil
}
