package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	driver "github.com/tjmaynes/learning-golang/driver"
	handler "github.com/tjmaynes/learning-golang/handler"
)

// App ..
type App struct {
	DbConn sql.DB
	Router http.Handler
}

// NewApp ..
func NewApp(dbSource string, dbType string) *App {
	dbConn, err := driver.ConnectDB(dbSource, dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	h := handler.Initialize(dbConn)

	return &App{
		DbConn: *dbConn,
		Router: h,
	}
}

// Run ..
func (a *App) Run(serverPort string) {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", serverPort),
		Handler:        a.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(fmt.Sprintf("Running server on port %s...", serverPort))

	idleConnsClosed := make(chan struct{})
	go setupGracefulShutdown(server, &a.DbConn, idleConnsClosed)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Server closed: %v", err)
	}

	<-idleConnsClosed
}

func setupGracefulShutdown(server *http.Server, db *sql.DB, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}

	defer db.Close()

	close(idleConnsClosed)
}
