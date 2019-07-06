// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cart

import (
	"context"
	"sync"
)

var (
	lockRepositoryMockAddItem     sync.RWMutex
	lockRepositoryMockGetItemByID sync.RWMutex
	lockRepositoryMockGetItems    sync.RWMutex
	lockRepositoryMockRemoveItem  sync.RWMutex
	lockRepositoryMockUpdateItem  sync.RWMutex
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked Repository
//         mockedRepository := &RepositoryMock{
//             AddItemFunc: func(ctx context.Context, item *Item) (Item, error) {
// 	               panic("mock out the AddItem method")
//             },
//             GetItemByIDFunc: func(ctx context.Context, id int64) (Item, error) {
// 	               panic("mock out the GetItemByID method")
//             },
//             GetItemsFunc: func(ctx context.Context, limit int64) ([]Item, error) {
// 	               panic("mock out the GetItems method")
//             },
//             RemoveItemFunc: func(ctx context.Context, id int64) (int64, error) {
// 	               panic("mock out the RemoveItem method")
//             },
//             UpdateItemFunc: func(ctx context.Context, item *Item) (Item, error) {
// 	               panic("mock out the UpdateItem method")
//             },
//         }
//
//         // use mockedRepository in code that requires Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// AddItemFunc mocks the AddItem method.
	AddItemFunc func(ctx context.Context, item *Item) (Item, error)

	// GetItemByIDFunc mocks the GetItemByID method.
	GetItemByIDFunc func(ctx context.Context, id int64) (Item, error)

	// GetItemsFunc mocks the GetItems method.
	GetItemsFunc func(ctx context.Context, limit int64) ([]Item, error)

	// RemoveItemFunc mocks the RemoveItem method.
	RemoveItemFunc func(ctx context.Context, id int64) (int64, error)

	// UpdateItemFunc mocks the UpdateItem method.
	UpdateItemFunc func(ctx context.Context, item *Item) (Item, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddItem holds details about calls to the AddItem method.
		AddItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Item is the item argument value.
			Item *Item
		}
		// GetItemByID holds details about calls to the GetItemByID method.
		GetItemByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// GetItems holds details about calls to the GetItems method.
		GetItems []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Limit is the limit argument value.
			Limit int64
		}
		// RemoveItem holds details about calls to the RemoveItem method.
		RemoveItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// UpdateItem holds details about calls to the UpdateItem method.
		UpdateItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Item is the item argument value.
			Item *Item
		}
	}
}

// AddItem calls AddItemFunc.
func (mock *RepositoryMock) AddItem(ctx context.Context, item *Item) (Item, error) {
	if mock.AddItemFunc == nil {
		panic("RepositoryMock.AddItemFunc: method is nil but Repository.AddItem was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Item *Item
	}{
		Ctx:  ctx,
		Item: item,
	}
	lockRepositoryMockAddItem.Lock()
	mock.calls.AddItem = append(mock.calls.AddItem, callInfo)
	lockRepositoryMockAddItem.Unlock()
	return mock.AddItemFunc(ctx, item)
}

// AddItemCalls gets all the calls that were made to AddItem.
// Check the length with:
//     len(mockedRepository.AddItemCalls())
func (mock *RepositoryMock) AddItemCalls() []struct {
	Ctx  context.Context
	Item *Item
} {
	var calls []struct {
		Ctx  context.Context
		Item *Item
	}
	lockRepositoryMockAddItem.RLock()
	calls = mock.calls.AddItem
	lockRepositoryMockAddItem.RUnlock()
	return calls
}

// GetItemByID calls GetItemByIDFunc.
func (mock *RepositoryMock) GetItemByID(ctx context.Context, id int64) (Item, error) {
	if mock.GetItemByIDFunc == nil {
		panic("RepositoryMock.GetItemByIDFunc: method is nil but Repository.GetItemByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockGetItemByID.Lock()
	mock.calls.GetItemByID = append(mock.calls.GetItemByID, callInfo)
	lockRepositoryMockGetItemByID.Unlock()
	return mock.GetItemByIDFunc(ctx, id)
}

// GetItemByIDCalls gets all the calls that were made to GetItemByID.
// Check the length with:
//     len(mockedRepository.GetItemByIDCalls())
func (mock *RepositoryMock) GetItemByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockGetItemByID.RLock()
	calls = mock.calls.GetItemByID
	lockRepositoryMockGetItemByID.RUnlock()
	return calls
}

// GetItems calls GetItemsFunc.
func (mock *RepositoryMock) GetItems(ctx context.Context, limit int64) ([]Item, error) {
	if mock.GetItemsFunc == nil {
		panic("RepositoryMock.GetItemsFunc: method is nil but Repository.GetItems was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Limit int64
	}{
		Ctx:   ctx,
		Limit: limit,
	}
	lockRepositoryMockGetItems.Lock()
	mock.calls.GetItems = append(mock.calls.GetItems, callInfo)
	lockRepositoryMockGetItems.Unlock()
	return mock.GetItemsFunc(ctx, limit)
}

// GetItemsCalls gets all the calls that were made to GetItems.
// Check the length with:
//     len(mockedRepository.GetItemsCalls())
func (mock *RepositoryMock) GetItemsCalls() []struct {
	Ctx   context.Context
	Limit int64
} {
	var calls []struct {
		Ctx   context.Context
		Limit int64
	}
	lockRepositoryMockGetItems.RLock()
	calls = mock.calls.GetItems
	lockRepositoryMockGetItems.RUnlock()
	return calls
}

// RemoveItem calls RemoveItemFunc.
func (mock *RepositoryMock) RemoveItem(ctx context.Context, id int64) (int64, error) {
	if mock.RemoveItemFunc == nil {
		panic("RepositoryMock.RemoveItemFunc: method is nil but Repository.RemoveItem was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockRemoveItem.Lock()
	mock.calls.RemoveItem = append(mock.calls.RemoveItem, callInfo)
	lockRepositoryMockRemoveItem.Unlock()
	return mock.RemoveItemFunc(ctx, id)
}

// RemoveItemCalls gets all the calls that were made to RemoveItem.
// Check the length with:
//     len(mockedRepository.RemoveItemCalls())
func (mock *RepositoryMock) RemoveItemCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockRemoveItem.RLock()
	calls = mock.calls.RemoveItem
	lockRepositoryMockRemoveItem.RUnlock()
	return calls
}

// UpdateItem calls UpdateItemFunc.
func (mock *RepositoryMock) UpdateItem(ctx context.Context, item *Item) (Item, error) {
	if mock.UpdateItemFunc == nil {
		panic("RepositoryMock.UpdateItemFunc: method is nil but Repository.UpdateItem was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Item *Item
	}{
		Ctx:  ctx,
		Item: item,
	}
	lockRepositoryMockUpdateItem.Lock()
	mock.calls.UpdateItem = append(mock.calls.UpdateItem, callInfo)
	lockRepositoryMockUpdateItem.Unlock()
	return mock.UpdateItemFunc(ctx, item)
}

// UpdateItemCalls gets all the calls that were made to UpdateItem.
// Check the length with:
//     len(mockedRepository.UpdateItemCalls())
func (mock *RepositoryMock) UpdateItemCalls() []struct {
	Ctx  context.Context
	Item *Item
} {
	var calls []struct {
		Ctx  context.Context
		Item *Item
	}
	lockRepositoryMockUpdateItem.RLock()
	calls = mock.calls.UpdateItem
	lockRepositoryMockUpdateItem.RUnlock()
	return calls
}