package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/net"
)

// GetNetworkStatus handles the network status endpoint
func GetNetworkStatus(c *gin.Context) {
	netStats, _ := net.IOCounters(false)
	connections, _ := net.Connections("all")

	var totalBytesSent, totalBytesRecv uint64
	if len(netStats) > 0 {
		totalBytesSent = netStats[0].BytesSent
		totalBytesRecv = netStats[0].BytesRecv
	}

	c.JSON(http.StatusOK, gin.H{
		"total_bytes_sent":     totalBytesSent,
		"total_bytes_received": totalBytesRecv,
		"active_connections":   len(connections),
		"timestamp":            time.Now().Unix(),
	})
}
