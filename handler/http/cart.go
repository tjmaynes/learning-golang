package handler

import (
	"net/http"
	"strconv"
	"strings"

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

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	data, err := c.Service.GetAllItems(r.Context(), limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string][]cart.Item{"data": data})
}

// AddCartItem ..
func (c *CartHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	r.ParseForm()

	itemName := r.Form.Get("name")
	itemManufacturer := r.Form.Get("manufacturer")
	itemPrice, err := strconv.ParseInt(r.Form.Get("price"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid price value.", 400)
	}

	data, err := c.Service.AddCartItem(r.Context(), itemName, cart.Decimal(itemPrice), itemManufacturer)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusCreated, map[string]cart.Item{"data": data})
}

// GetCartItemByID ..
func (c *CartHandler) GetCartItemByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	params := strings.Split(r.URL.Path, "/")
	if len(params) < 2 {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	id, err := strconv.ParseInt(params[2], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	data, err := c.Service.GetItemByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string]cart.Item{"data": data})
}

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
