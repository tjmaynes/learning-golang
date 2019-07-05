package cart

import (
	"context"
)

// Service ..
type Service interface {
	GetAllItems(ctx context.Context, limit int64) ([]Item, error)
	GetItemByID(ctx context.Context, id int64) (Item, error)
	AddCartItem(ctx context.Context, name string, price Decimal, manufacturer string) (Item, error)
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

func (s *service) GetItemByID(ctx context.Context, id int64) (Item, error) {
	return s.Repository.GetItemByID(ctx, id)
}

func (s *service) AddCartItem(ctx context.Context, name string, price Decimal, manufacturer string) (Item, error) {
	item := Item{Name: name, Price: price, Manufacturer: manufacturer}

	err := item.Validate()
	if err != nil {
		return Item{}, err
	}

	return s.Repository.AddItem(ctx, &item)
}
