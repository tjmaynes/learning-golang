package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/tjmaynes/learning-golang/handler/http"
)

func addPostRouter(postHandler *handlers.PostHandler) http.Handler {
	router := chi.NewRouter()

	router.Get("/", postHandler.GetPosts)
	router.Get("/{id:[0-9]+}", postHandler.GetPostByID)
	router.Post("/", postHandler.AddPost)
	router.Put("/{id:[0-9]+}", postHandler.UpdatePost)
	router.Delete("/{id:[0-9]+}", postHandler.DeletePost)

	return router
}

// Initialize ..
func Initialize(dbConn *sql.DB) http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
	)

	postHandler := handlers.NewPostHandler(dbConn)
	router.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", addPostRouter(postHandler))
		rt.Get("/ping", handlers.GetPingHandler)
	})

	return router
}
