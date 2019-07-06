package cart

import (
	"context"
	"testing"

	"github.com/icrowley/fake"
)

func Test_Cart_Service_GetAllItems_WhenItemsExist_ShouldReturnAllItems(t *testing.T) {
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
	if err != nil {
		t.Fatalf("Should not have failed!")
	}

	if len(results) != len(items) {
		t.Errorf("Expected an array of cart items of size %d. Got %d", len(items), len(results))
	}

	callsToSend := len(mockRepository.GetItemsCalls())
	if callsToSend != 1 {
		t.Errorf("Send was called %d times", callsToSend)
	}

	if limitCalled != limit {
		t.Errorf("Unexpected recipient: %d", limitCalled)
	}
}

func Test_Cart_Service_GetItemByID_WhenItemExists_ShouldReturnItem(t *testing.T) {
	item := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}
	const id = 10
	var idCalled int64

	mockRepository := &RepositoryMock{
		GetItemByIDFunc: func(ctx context.Context, id int64) (Item, error) {
			idCalled = id
			return item, nil
		},
	}

	ctx := context.Background()
	sut := NewService(mockRepository)

	result, err := sut.GetItemByID(ctx, id)
	if err != nil {
		t.Fatalf("Should not have failed!")
	}

	if result != item {
		t.Errorf("Expected cart items %+v. Got %+v", item, result)
	}

	callsToSend := len(mockRepository.GetItemByIDCalls())
	if callsToSend != 1 {
		t.Errorf("Send was called %d times", callsToSend)
	}

	if idCalled != id {
		t.Errorf("Unexpected recipient: %d", id)
	}
}

func Test_Cart_Service_GetItemByID_WhenItemDoesNotExist_ShouldReturnError(t *testing.T) {
	testError := createError()

	mockRepository := &RepositoryMock{
		GetItemByIDFunc: func(ctx context.Context, id int64) (Item, error) {
			return Item{}, testError
		},
	}

	ctx := context.Background()
	sut := NewService(mockRepository)

	_, err := sut.GetItemByID(ctx, 0)
	if err != testError {
		t.Errorf("Expected error message %s. Got %s", testError, err)
	}

	callsToSend := len(mockRepository.GetItemByIDCalls())
	if callsToSend != 1 {
		t.Errorf("Send was called %d times", callsToSend)
	}
}

func Test_Cart_Service_AddItem_WhenGivenValidItem_ShouldReturnItem(t *testing.T) {
	var itemCalled *Item
	newItem := Item{
		ID:           int64(1),
		Name:         fake.ProductName(),
		Price:        Decimal(99),
		Manufacturer: fake.Brand(),
	}

	mockRepository := &RepositoryMock{
		AddItemFunc: func(ctx context.Context, item *Item) (Item, error) {
			itemCalled = &newItem
			return newItem, nil
		},
	}

	ctx := context.Background()
	sut := NewService(mockRepository)

	result, err := sut.AddCartItem(ctx, newItem.Name, newItem.Price, newItem.Manufacturer)
	if err != nil {
		t.Fatalf("Should not have failed!")
	}

	if result != *itemCalled {
		t.Errorf("Expected cart item: %+v. Got %+v", itemCalled, result)
	}

	callsToSend := len(mockRepository.AddItemCalls())
	if callsToSend != 1 {
		t.Errorf("Send was called %d times", callsToSend)
	}
}

func Test_Cart_Service_AddItem_WhenGivenInvalidItem_ShouldReturnError(t *testing.T) {
	newItem := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}

	mockRepository := &RepositoryMock{
		AddItemFunc: func(ctx context.Context, item *Item) (Item, error) {
			return Item{}, nil
		},
	}

	ctx := context.Background()
	sut := NewService(mockRepository)

	expectedErrorMessage := "price: must be no less than 99."

	_, err := sut.AddCartItem(ctx, newItem.Name, newItem.Price, newItem.Manufacturer)
	if err.Error() != expectedErrorMessage {
		t.Errorf("Error unexpected error message %s was given", err)
	}

	callsToSend := len(mockRepository.AddItemCalls())
	if callsToSend != 0 {
		t.Errorf("Send was called %d times", callsToSend)
	}
}
