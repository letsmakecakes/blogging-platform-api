package main

import (
	"bloggingplatformapi/internal/config"
	"bloggingplatformapi/internal/routes"
	"bloggingplatformapi/internal/utils"
	"bloggingplatformapi/pkg/db"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize Logrus
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	log.Info("Loaded config file")

	// Initialize database
	database, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}(database)

	log.Info("Database initialized")

	// Set Gin to release mode if not in development
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Add Logrus logging middleware
	router.Use(utils.GinLogrus(logger), gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router, database)

	// Create custom HTTP server with timeouts
	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	// Start server
	log.Printf("Server running on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Could not start server: %v", err)
	}
}
