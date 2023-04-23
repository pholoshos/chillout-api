// cmd/my-api/main.go

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/myuser/my-api/internal/handlers"
	"github.com/myuser/my-api/internal/middleware"
	"github.com/myuser/my-api/internal/server"
	"github.com/myuser/my-api/pkg/utils"
)

func main() {
	// Create a new logger
	logger := utils.NewLogger(os.Stdout)

	// Load configuration from environment variables
	config := server.LoadConfigFromEnv()

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Create a new server
	s := server.New(
		config,
		router,
		logger,
	)

	// Register middleware
	router.Use(middleware.Logger(logger))

	// Register handlers
	handlers.RegisterHandlers(router, s)

	// Start the server
	srv := &http.Server{
		Handler:      router,
		Addr:         config.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Starting server on ", config.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("Server failed: ", err)
	}
}
