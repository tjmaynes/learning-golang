package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/tjmaynes/learning-golang/handler/http"
	"github.com/tjmaynes/learning-golang/pkg/cart"
)

// Initialize ..
func Initialize(cartService cart.Service) http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
	)

	cartHandler := handlers.NewCartHandler(cartService)
	router.Route("/", func(rt chi.Router) {
		rt.Mount("/cart", addCartRouter(cartHandler))
		rt.Get("/ping", handlers.GetPingHandler)
	})

	return router
}

func addCartRouter(cartHandler *handlers.CartHandler) http.Handler {
	router := chi.NewRouter()

	router.Get("/", cartHandler.GetCartItems)
	router.Get("/{id:[0-9]+}", cartHandler.GetCartItemByID)
	router.Post("/", cartHandler.AddCartItem)
	router.Put("/{id:[0-9]+}", cartHandler.UpdateCartItem)
	router.Delete("/{id:[0-9]+}", cartHandler.RemoveCartItem)

	return router
}
