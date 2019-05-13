package post

import (
	"context"
	"errors"

	driver "github.com/tjmaynes/learning-golang/driver"
)

// Repository ..
type Repository interface {
	GetPosts(ctx context.Context, limit int64) ([]*Post, error)
	GetByPostID(ctx context.Context, id int64) (*Post, error)
	AddPost(ctx context.Context, post *Post) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) (*Post, error)
	DeletePost(ctx context.Context, id int64) (bool, error)
}

// NewPostRepository ..
func NewPostRepository(DBConn *driver.DB) Repository {
	return &Repo{DBConn: DBConn}
}

// Repo ..
type Repo struct {
	DBConn *driver.DB
}

// FetchQuery ..
func (repo *Repo) FetchQuery(ctx context.Context, query string, args ...interface{}) ([]*Post, error) {
	rows, err := repo.DBConn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*Post, 0)
	for rows.Next() {
		data := new(Post)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}

	return payload, nil
}

// GetPosts ..
func (repo *Repo) GetPosts(ctx context.Context, limit int64) ([]*Post, error) {
	return repo.FetchQuery(ctx, "SELECT id, title, content FROM post LIMIT $1", limit)
}

// GetByPostID ..
func (repo *Repo) GetByPostID(ctx context.Context, id int64) (*Post, error) {
	rows, err := repo.FetchQuery(ctx, "SELECT id, title, content FROM post WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	post := &Post{}
	if len(rows) > 0 {
		post = rows[0]
	} else {
		return nil, errors.New("requested item not found")
	}

	return post, nil
}

// AddPost ..
func (repo *Repo) AddPost(ctx context.Context, post *Post) (*Post, error) {
	var id int64
	err := repo.DBConn.QueryRowContext(ctx, "INSERT INTO post (title, content) VALUES ($1, $2) RETURNING id", post.Title, post.Content).Scan(&id)
	if err != nil {
		panic(err)
	}

	newPost := &Post{ID: id, Title: post.Title, Content: post.Content}
	return newPost, nil
}

// UpdatePost ..
func (repo *Repo) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	stmt, err := repo.DBConn.PrepareContext(ctx, "UPDATE post SET title = $2, content = $3 WHERE id = $1")
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(ctx, post.ID, post.Title, post.Content)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return post, nil
}

// DeletePost ..
func (repo *Repo) DeletePost(ctx context.Context, id int64) (bool, error) {
	stmt, err := repo.DBConn.PrepareContext(ctx, "DELETE FROM post WHERE id = $1")
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	return true, nil
}
