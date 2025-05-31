package handlers

import (
	"net/http"

	"homeserver/types"

	"github.com/gin-gonic/gin"
)

// GetServices handles the services endpoint
func GetServices(c *gin.Context) {
	// Mock service data - in a real implementation, you'd query actual system services
	services := []types.Service{
		{Name: "homeserver-api", Status: "running", PID: 1234},
		{Name: "nginx", Status: "running", PID: 567},
		{Name: "docker", Status: "running", PID: 890},
		{Name: "ssh", Status: "running", PID: 22},
	}

	c.JSON(http.StatusOK, services)
}
