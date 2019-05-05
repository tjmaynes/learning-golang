package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tjmaynes/learning-golang/db"
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

func setupGracefulShutdown(server *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}

	close(idleConnsClosed)
}

// Run ..
func Run(dbConn *db.DB, serverPort string) {
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

	server := http.Server{
		Addr:           fmt.Sprintf(":%s", serverPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(fmt.Sprintf("Running server on port %s...", serverPort))

	idleConnsClosed := make(chan struct{})
	go setupGracefulShutdown(&server, idleConnsClosed)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Server closed: %v", err)
	}

	<-idleConnsClosed
}
