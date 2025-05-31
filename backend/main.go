package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"homeserver/handlers"
)

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/api/health", handlers.GetHealth)

	// System information endpoint
	r.GET("/api/system", handlers.GetSystemInfo)

	// Services endpoint (mock data for now)
	r.GET("/api/services", handlers.GetServices)

	// Files endpoint (basic file listing)
	r.GET("/api/files", handlers.GetFiles)

	// Network status endpoint
	r.GET("/api/network", handlers.GetNetworkStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
