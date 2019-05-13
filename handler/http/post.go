package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "github.com/tjmaynes/learning-golang/driver"
	"github.com/tjmaynes/learning-golang/post"
)

// NewPostHandler ,,
func NewPostHandler(db *driver.DB) *PostHandler {
	return &PostHandler{Repo: post.NewPostRepository(db)}
}

// PostHandler ..
type PostHandler struct {
	Repo post.Repository
}

// GetPosts ..
func (p *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	posts, err := p.Repo.GetPosts(r.Context(), 10)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json, _ := json.Marshal(posts)
	RespondWithJSON(w, http.StatusCreated, map[string][]byte{"data": json})
}

// AddPost ..
func (p *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	post := post.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	newPost, err := p.Repo.AddPost(r.Context(), &post)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json, _ := json.Marshal(newPost)
	RespondWithJSON(w, http.StatusCreated, map[string][]byte{"data": json})
}

// GetPostByID ..
func (p *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	newID, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	post, err := p.Repo.GetByPostID(r.Context(), newID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json, _ := json.Marshal(post)
	RespondWithJSON(w, http.StatusCreated, map[string][]byte{"data": json})
}

// UpdatePost ..
func (p *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	post := post.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	updatedPost, err := p.Repo.UpdatePost(r.Context(), &post)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json, _ := json.Marshal(updatedPost)
	RespondWithJSON(w, http.StatusCreated, map[string][]byte{"data": json})
}

// DeletePost ..
func (p *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	newID, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	_, err = p.Repo.DeletePost(r.Context(), newID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	message := fmt.Sprintf("Deleted Post ID: %v", newID)
	RespondWithJSON(w, http.StatusCreated, map[string]string{"message": message})
}
