package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetHealth handles the health check endpoint
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"service":   "homeserver-api",
	})
}
