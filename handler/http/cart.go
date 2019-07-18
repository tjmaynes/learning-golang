package handler

import (
	"encoding/json"
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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

// GetCartItemByID ..
func (c *CartHandler) GetCartItemByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, errorCode := getID(r.URL.Path)
	if errorCode >= 400 {
		http.Error(w, http.StatusText(errorCode), errorCode)
	}

	data, err := c.Service.GetItemByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string]cart.Item{"data": data})
}

// AddCartItem ..
func (c *CartHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	itemName := r.Form.Get("name")
	itemManufacturer := r.Form.Get("manufacturer")
	itemPrice, errorCode := getItemPrice(r.Form.Get("price"))
	if errorCode >= 400 {
		http.Error(w, http.StatusText(errorCode), errorCode)
		return
	}

	data, err := c.Service.AddCartItem(r.Context(), itemName, cart.Decimal(itemPrice), itemManufacturer)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusCreated, map[string]cart.Item{"data": data})
}

// UpdateCartItem ..
func (c *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, errorCode := getID(r.URL.Path)
	if errorCode >= 400 {
		http.Error(w, http.StatusText(errorCode), errorCode)
		return
	}

	decoder := json.NewDecoder(r.Body)
	type itemRequest struct {
		Name         string `json:"name"`
		Price        string `json:"price"`
		Manufacturer string `json:"manufacturer"`
	}
	var item itemRequest
	err := decoder.Decode(&item)
	if err != nil {
		panic(err)
	}

	itemPrice, errorCode := getItemPrice(item.Price)
	if errorCode >= 400 {
		http.Error(w, "Price value is invalid.", errorCode)
		return
	}

	result, err := c.Service.GetItemByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if id != result.ID {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	data, serviceError := c.Service.UpdateCartItem(r.Context(), id, item.Name, cart.Decimal(itemPrice), item.Manufacturer)
	if serviceError != nil {
		switch serviceError.StatusCode() {
		case cart.InvalidItem:
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		default:
			http.Error(w, http.StatusText(500), 500)
		}
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string]cart.Item{"data": data})
}

// RemoveCartItem ..
func (c *CartHandler) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, errorCode := getID(r.URL.Path)
	if errorCode >= 400 {
		http.Error(w, http.StatusText(errorCode), errorCode)
	}

	_, serviceError := c.Service.RemoveCartItem(r.Context(), id)
	if serviceError != nil {
		http.Error(w, serviceError.Message(), 500)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, http.StatusText(200))
}

func getID(urlPath string) (int64, int) {
	params := strings.Split(urlPath, "/")
	if len(params) < 2 {
		return 0, http.StatusBadRequest
	}

	id, err := strconv.ParseInt(params[2], 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}

	return id, 0
}

func getItemPrice(rawPrice string) (int64, int) {
	result, err := strconv.ParseInt(rawPrice, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return result, 0
}
