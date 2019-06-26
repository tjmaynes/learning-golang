package cart

import (
	"context"
	"testing"

	"github.com/icrowley/fake"
)

func Test_Cart_Service_GetAllItems_ShouldReturnPaginatedResponse(t *testing.T) {
	items := []Item{
		Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()},
		Item{ID: 2, Name: fake.ProductName(), Price: 4, Manufacturer: fake.Brand()},
		Item{ID: 3, Name: fake.ProductName(), Price: 5, Manufacturer: fake.Brand()},
		Item{ID: 4, Name: fake.ProductName(), Price: 11, Manufacturer: fake.Brand()},
		Item{ID: 5, Name: fake.ProductName(), Price: 100, Manufacturer: fake.Brand()},
	}

	const limit = 10
	var limitCalled int64

	mockRepository := &RepositoryMock{
		GetItemsFunc: func(ctx context.Context, limit int64) ([]Item, error) {
			limitCalled = limit
			return items, nil
		},
	}

	ctx := context.Background()
	sut := NewService(mockRepository)

	results, err := sut.GetAllItems(ctx, limit)
	if len(results) != len(items) {
		t.Errorf("Expected an array of cart items of size %d. Got %d", len(items), len(results))
	}

	callsToSend := len(mockRepository.GetItemsCalls())
	if err != nil {
		t.Fatalf("Should not have failed!")
	}

	if callsToSend != 1 {
		t.Errorf("Send was called %d times", callsToSend)
	}

	if limitCalled != limit {
		t.Errorf("Unexpected recipient: %d", limitCalled)
	}
}
