package cart

import (
	"context"
)

// Service ..
type Service interface {
	GetAllItems(ctx context.Context, limit int64) ([]Item, error)
	GetItemByID(ctx context.Context, id int64) (Item, error)
	AddCartItem(
		ctx context.Context,
		name string,
		price Decimal,
		manufacturer string,
	) (Item, error)
	UpdateCartItem(
		ctx context.Context,
		id int64,
		name string,
		price Decimal,
		manufacturer string,
	) (Item, ServiceError)
	RemoveCartItem(ctx context.Context, id int64) (int64, ServiceError)
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

// GetAllItems ..
func (s *service) GetAllItems(ctx context.Context, limit int64) ([]Item, error) {
	return s.Repository.GetItems(ctx, limit)
}

// GetItemByID ..
func (s *service) GetItemByID(ctx context.Context, id int64) (Item, error) {
	return s.Repository.GetItemByID(ctx, id)
}

// AddCartItem ..
func (s *service) AddCartItem(ctx context.Context, name string, price Decimal, manufacturer string) (Item, error) {
	item := Item{Name: name, Price: price, Manufacturer: manufacturer}

	err := item.Validate()
	if err != nil {
		return Item{}, err
	}

	return s.Repository.AddItem(ctx, &item)
}

// UpdateCartItem ..
func (s *service) UpdateCartItem(ctx context.Context, id int64, name string, price Decimal, manufacturer string) (Item, ServiceError) {
	item := Item{ID: id, Name: name, Price: price, Manufacturer: manufacturer}

	err := item.Validate()
	if err != nil {
		return Item{}, CreateServiceError(err.Error(), InvalidItem)
	}

	result, err := s.Repository.UpdateItem(ctx, &item)
	if err != nil {
		return Item{}, CreateServiceError(err.Error(), UnknownException)
	}

	return result, nil
}

// RemoveCartItem ..
func (s *service) RemoveCartItem(ctx context.Context, id int64) (int64, ServiceError) {
	result, err := s.Repository.RemoveItem(ctx, id)
	if err != nil {
		return id, CreateServiceError(err.Error(), UnknownException)
	}

	return result, nil
}
