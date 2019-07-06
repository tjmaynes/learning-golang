package cart

import (
	"context"
	"database/sql"
)

// Repository ..
type Repository interface {
	GetItems(ctx context.Context, limit int64) ([]Item, error)
	GetItemByID(ctx context.Context, id int64) (Item, error)
	AddItem(ctx context.Context, item *Item) (Item, error)
	UpdateItem(ctx context.Context, item *Item) (Item, error)
	RemoveItem(ctx context.Context, id int64) (int64, error)
}

// NewRepository ..
func NewRepository(DBConn *sql.DB) Repository {
	return &repository{DBConn: DBConn}
}

// repository ..
type repository struct {
	DBConn *sql.DB
}

// FetchQuery ..
func (r *repository) FetchQuery(ctx context.Context, query string, args ...interface{}) ([]Item, error) {
	rows, err := r.DBConn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(&data.ID, &data.Name, &data.Price, &data.Manufacturer)
		if err != nil {
			return nil, err
		}
		payload = append(payload, *data)
	}

	return payload, nil
}

// GetItems ..
func (r *repository) GetItems(ctx context.Context, limit int64) ([]Item, error) {
	rows, err := r.DBConn.QueryContext(ctx, "SELECT id, name, price, manufacturer FROM cart LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(&data.ID, &data.Name, &data.Price, &data.Manufacturer)
		if err != nil {
			return nil, err
		}
		payload = append(payload, *data)
	}

	return payload, nil
}

// GetItemByID ..
func (r *repository) GetItemByID(ctx context.Context, id int64) (Item, error) {
	var item Item
	row := r.DBConn.QueryRowContext(ctx, "SELECT id, name, price, manufacturer FROM cart WHERE id = ?", id)
	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.Manufacturer)
	if err != nil && err != sql.ErrNoRows {
		return Item{}, err
	}

	return item, nil
}

// AddItem ..
func (r *repository) AddItem(ctx context.Context, item *Item) (Item, error) {
	newItem := Item{}

	stmt, err := r.DBConn.PrepareContext(ctx, "INSERT INTO cart (name, price, manufacturer) VALUES (?, ?, ?)")
	if err != nil {
		return newItem, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, item.Name, item.Price, item.Manufacturer)
	if err != nil {
		return newItem, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return newItem, err
	}

	return Item{
		ID:           insertedID,
		Name:         item.Name,
		Price:        item.Price,
		Manufacturer: item.Manufacturer,
	}, nil
}

// UpdateItem ..
func (r *repository) UpdateItem(ctx context.Context, item *Item) (Item, error) {
	updatedItem := Item{}

	stmt, err := r.DBConn.PrepareContext(ctx, "UPDATE cart SET name = ?, price = ?, manufacturer = ? WHERE id = ?")
	if err != nil {
		return updatedItem, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, item.ID, item.Name, item.Price, item.Manufacturer)
	if err != nil {
		return updatedItem, err
	}

	return *item, nil
}

// RemoveItem ..
func (r *repository) RemoveItem(ctx context.Context, id int64) (int64, error) {
	stmt, err := r.DBConn.PrepareContext(ctx, "DELETE FROM cart WHERE id = ?")
	if err != nil {
		return id, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return id, err
	}

	return id, nil
}
