package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"homeserver/types"

	"github.com/gin-gonic/gin"
)

// GetServices handles the services endpoint
func GetServices(c *gin.Context) {
	services := getActualSystemServices()
	c.JSON(http.StatusOK, services)
}

// getActualSystemServices queries the actual system for service information
func getActualSystemServices() []types.Service {
	services := []types.Service{}

	// List of common services to check on macOS
	servicesToCheck := []string{
		"ssh",
		"nginx",
		"docker",
		"postgresql",
		"mysql",
		"httpd",
		"apache2",
		"redis-server",
		"mongodb",
	}

	for _, serviceName := range servicesToCheck {
		service := checkServiceStatus(serviceName)
		if service.Status != "not found" {
			services = append(services, service)
		}
	}

	// Also check for running processes by name
	additionalProcesses := []string{
		"homeserver-api",
		"node",
		"go",
	}

	for _, processName := range additionalProcesses {
		processes := checkProcessByName(processName)
		services = append(services, processes...)
	}

	return services
}

// checkServiceStatus checks if a service is running using launchctl on macOS
func checkServiceStatus(serviceName string) types.Service {
	service := types.Service{
		Name:   serviceName,
		Status: "not found",
		PID:    0,
	}

	// Try launchctl first (for system services)
	cmd := exec.Command("launchctl", "list")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(strings.ToLower(line), strings.ToLower(serviceName)) {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					service.Status = "running"
					if pidStr := fields[0]; pidStr != "-" {
						if pid, err := strconv.ParseInt(pidStr, 10, 32); err == nil {
							service.PID = int32(pid)
						}
					}
					return service
				}
			}
		}
	}

	// If not found in launchctl, try checking running processes
	process := checkProcessByName(serviceName)
	if len(process) > 0 {
		return process[0]
	}

	return service
}

// checkProcessByName checks for running processes by name using ps
func checkProcessByName(processName string) []types.Service {
	var services []types.Service

	// Use ps to find processes
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return services
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(processName)) && !strings.Contains(line, "ps aux") {
			fields := strings.Fields(line)
			if len(fields) >= 11 {
				service := types.Service{
					Name:   processName,
					Status: "running",
				}

				// Parse PID (second field in ps aux output)
				if pid, err := strconv.ParseInt(fields[1], 10, 32); err == nil {
					service.PID = int32(pid)
				}

				// Make service name more descriptive based on command
				if len(fields) >= 11 {
					cmdParts := fields[10:]
					if len(cmdParts) > 0 {
						service.Name = fmt.Sprintf("%s (%s)", processName, strings.Join(cmdParts[:min(3, len(cmdParts))], " "))
					}
				}

				services = append(services, service)
			}
		}
	}

	return services
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
