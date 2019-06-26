package handler

import (
	"net/http"
	// "strconv"

	jsonHandler "github.com/tjmaynes/learning-golang/handler/json"
	cart "github.com/tjmaynes/learning-golang/pkg/cart"
)

// NewCartHandler ..
func NewCartHandler(service cart.Service) *CartHandler {
	return &CartHandler{Service: service}
}

// CartHandler ..
type CartHandler struct {
	Service cart.Service
}

// GetCartItems ..
func (c *CartHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	data, err := c.Service.GetAllItems(r.Context(), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string][]cart.Item{"data": data})
}

// // AddCartItem ..
// func (c *CartHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	post := post.Post{}
// 	json.NewDecoder(r.Body).Decode(&post)
// 	newCartItem, err := c.Repo.AddCartItem(r.Context(), &post)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	json, _ := json.Marshal(newCartItem)
// 	jsonHandler.CreateResponse(w, http.StatusCreated, map[string][]byte{"data": json})
// }

// // GetCartItemByID ..
// func (c *CartHandler) GetCartItemByID(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	ids, ok := r.URL.Query()["id"]
// 	if !ok || len(ids) < 1 {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	newID, err := strconv.ParseInt(ids[0], 10, 64)
// 	if err != nil {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	post, err := c.Repo.GetCartItemByID(r.Context(), newID)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	json, _ := json.Marshal(post)
// 	jsonHandler.CreateResponse(w, http.StatusCreated, map[string][]byte{"data": json})
// }

// // UpdateCartItem ..
// func (c *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "PUT" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	post := post.Post{}
// 	json.NewDecoder(r.Body).Decode(&post)

// 	updatedPost, err := c.Repo.UpdatePost(r.Context(), &post)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	json, _ := json.Marshal(updatedPost)
// 	jsonHandler.CreateResponse(w, http.StatusCreated, map[string][]byte{"data": json})
// }

// // DeleteCartItem ..
// func (c *CartHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "DELETE" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	ids, ok := r.URL.Query()["id"]
// 	if !ok || len(ids) < 1 {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	newID, err := strconv.ParseInt(ids[0], 10, 64)
// 	if err != nil {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	_, err = c.Repo.DeleteCartItem(r.Context(), newID)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	message := fmt.Sprintf("Deleted Post ID: %v", newID)
// 	jsonHandler.CreateResponse(w, http.StatusCreated, map[string]string{"message": message})
// }
