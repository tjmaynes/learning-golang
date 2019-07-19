package cart

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/icrowley/fake"
)

func Test_Cart_Repository_GetItems_ShouldReturnItems(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	const limit = 5

	columns := []string{"id", "name", "price", "manufacturer"}
	item1 := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}
	item2 := Item{ID: 2, Name: fake.ProductName(), Price: 4, Manufacturer: fake.Brand()}
	item3 := Item{ID: 3, Name: fake.ProductName(), Price: 5, Manufacturer: fake.Brand()}
	item4 := Item{ID: 4, Name: fake.ProductName(), Price: 11, Manufacturer: fake.Brand()}
	item5 := Item{ID: 5, Name: fake.ProductName(), Price: 100, Manufacturer: fake.Brand()}

	mock.ExpectQuery("SELECT id, name, price, manufacturer FROM cart LIMIT \\?").
		WithArgs(limit).
		WillReturnRows(
			sqlmock.NewRows(columns).
				FromCSVString(convertObjectToCSV(item1)).
				FromCSVString(convertObjectToCSV(item2)).
				FromCSVString(convertObjectToCSV(item3)).
				FromCSVString(convertObjectToCSV(item4)).
				FromCSVString(convertObjectToCSV(item5)),
		).
		RowsWillBeClosed()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.GetItems(ctx, limit)
	if err != nil {
		t.Fatalf("Error '%s' was not expected when fetching cart items", err)
	}

	if len(result) != limit {
		t.Fatalf("Unexpected number of items were given, '%d'. Expected '%d'.", len(result), limit)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_GetItems_WhenErrorOccurs_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	const limit = 5
	error := createError()

	mock.ExpectQuery("SELECT id, name, price, manufacturer FROM cart LIMIT \\?").
		WithArgs(limit).
		WillReturnError(error)

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.GetItems(ctx, limit)
	if result != nil {
		t.Fatalf("Result '%s' was not expected when simulating a failed fetching cart item", err)
	}

	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating a failed fetching cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_GetItemByID_WhenItemExists_ShouldReturnItem(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	const limit = 5
	columns := []string{"id", "name", "price", "manufacturer"}
	item1 := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}

	mock.ExpectQuery("SELECT id, name, price, manufacturer FROM cart WHERE id = \\?").
		WithArgs(item1.ID).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString(convertObjectToCSV(item1))).
		RowsWillBeClosed()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.GetItemByID(ctx, item1.ID)
	if err != nil {
		t.Fatalf("Error '%s' was not expected when fetching cart item", err)
	}

	if result != item1 {
		t.Fatalf("Unexpected item was given, '%+v'. Expected '%+v'.", result, item1)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_GetItemByID_WhenItemDoesNotExist_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1ID := int64(1)
	error := createError()

	mock.ExpectQuery("SELECT id, name, price, manufacturer FROM cart WHERE id = \\?").
		WithArgs(item1ID).
		WillReturnError(error)

	sut := NewRepository(dbConn)
	ctx := context.Background()

	_, err = sut.GetItemByID(ctx, item1ID)
	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating a failed fetching cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_GetItemByID_WhenErrorOccurs_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1ID := int64(1)
	error := createError()

	mock.ExpectQuery("SELECT id, name, price, manufacturer FROM cart WHERE id = \\?").
		WithArgs(item1ID).
		WillReturnError(error)

	sut := NewRepository(dbConn)
	ctx := context.Background()

	_, err = sut.GetItemByID(ctx, item1ID)

	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating a failed fetching cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_AddItem_ShouldReturnInsertedItem(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item := Item{ID: 0, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO cart \\(name, price, manufacturer\\) VALUES \\(\\?, \\?, \\?\\)").
		WithArgs(item.Name, item.Price, item.Manufacturer).
		WillReturnResult(sqlmock.NewResult(item.ID, 0))
	mock.ExpectCommit()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.AddItem(ctx, &item)
	if err != nil {
		t.Fatalf("Error '%s' was not expected when adding an item to cart", err)
	}

	if result != item {
		t.Fatalf("Unexpected item was given, '%+v'. Expected '%+v'.", result, item)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_AddItem_WhenErrorOccurs_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}
	error := createError()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO cart \\(name, price, manufacturer\\) VALUES \\(\\?, \\?, \\?\\)").
		WithArgs(item.Name, item.Price, item.Manufacturer).
		WillReturnError(error)
	mock.ExpectRollback()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	_, err = sut.AddItem(ctx, &item)
	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating failure while adding cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_UpdateItem_ShouldUpdateSpecificItem(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1 := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE cart SET name = \\?, price = \\?, manufacturer = \\? WHERE id = \\?").
		WithArgs(item1.Name, item1.Price, item1.Manufacturer, item1.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.UpdateItem(ctx, &item1)
	if err != nil {
		t.Fatalf("Result '%s' was not expected when simulating failure while updating cart item", err)
	}

	if result != item1 {
		t.Fatalf("Unexpected item was given, '%+v'. Expected '%+v'.", result, item1)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_UpdateItem_WhenErrorOccurs_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1 := Item{ID: 1, Name: fake.ProductName(), Price: 23, Manufacturer: fake.Brand()}
	error := createError()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE cart SET name = \\?, price = \\?, manufacturer = \\? WHERE id = \\?").
		WithArgs(item1.Name, item1.Price, item1.Manufacturer, item1.ID).
		WillReturnError(error)
	mock.ExpectRollback()

	sut := NewRepository(dbConn)
	ctx := context.Background()

	_, err = sut.UpdateItem(ctx, &item1)
	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating failure while updating cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_RemoveItem_ShouldReturnID(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1ID := int64(1)

	mock.ExpectPrepare("DELETE FROM cart WHERE id = \\?").
		ExpectExec().
		WithArgs(item1ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.RemoveItem(ctx, item1ID)
	if err != nil {
		t.Fatalf("Result '%s' was not expected when simulating failure while removing cart item", err)
	}

	if result != item1ID {
		t.Fatalf("Unexpected id was given, '%d'. Expected '%d'.", result, item1ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Cart_Repository_RemoveItem_WhenErrorOccurs_ShouldReturnError(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	item1ID := int64(1)
	error := createError()

	mock.ExpectPrepare("DELETE FROM cart WHERE id = \\?").
		ExpectExec().
		WithArgs(item1ID).
		WillReturnError(error)

	sut := NewRepository(dbConn)
	ctx := context.Background()

	result, err := sut.RemoveItem(ctx, item1ID)
	if result != item1ID {
		t.Fatalf("Unexpected id was given, '%d'. Expected '%d'.", result, item1ID)
	}

	if error != err {
		t.Fatalf("Expected failure '%s', but received '%s' when simulating failure while removing cart item", error, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func convertObjectToCSV(item Item) string {
	return fmt.Sprintf("%d,%s,%d,%s", item.ID, item.Name, item.Price, item.Manufacturer)
}

func createError() error {
	return fmt.Errorf("some error")
}
