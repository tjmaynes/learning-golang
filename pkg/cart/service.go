package cart

import (
	"context"
)

// Service ..
type Service interface {
	GetAllItems(ctx context.Context, limit int64) ([]Item, error)
	// GetByCartItemID(ctx context.Context, id int64) (*CartItem, error)
	// AddCartItem(ctx context.Context, name string, price int64, manufacturer string) (*CartItem, error)
	// UpdateCartItem(ctx context.Context, item *CartItem) (*CartItem, error)
	// RemoveCartItem(ctx context.Context, id int64) (int64, error)
}

// NewService ..
func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}

type service struct {
	Repository Repository
}

func (s *service) GetAllItems(ctx context.Context, limit int64) ([]Item, error) {
	return s.Repository.GetItems(ctx, limit)
}
