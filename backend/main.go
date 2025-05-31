package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"homeserver/config"
	"homeserver/handlers"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode based on configuration
	gin.SetMode(cfg.GinMode)

	r := gin.Default()

	// Configure CORS using configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.CORSOrigins
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Health check endpoint
	r.GET("/api/health", handlers.GetHealth)

	// System information endpoint
	r.GET("/api/system", handlers.GetSystemInfo)

	// Services endpoint (mock data for now)
	r.GET("/api/services", handlers.GetServices)

	// Files endpoint (basic file listing) - only if enabled
	if cfg.EnableFileBrowser {
		r.GET("/api/files", handlers.GetFiles)
	}

	// Network status endpoint
	r.GET("/api/network", handlers.GetNetworkStatus)

	log.Printf("Starting server on port %s with CORS origins: %v", cfg.Port, cfg.CORSOrigins)
	r.Run(":" + cfg.Port)
}
