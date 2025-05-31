package main

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemInfo struct {
	Hostname     string        `json:"hostname"`
	Platform     string        `json:"platform"`
	OS           string        `json:"os"`
	Architecture string        `json:"architecture"`
	CPUCount     int           `json:"cpu_count"`
	Memory       MemoryInfo    `json:"memory"`
	CPU          CPUInfo       `json:"cpu"`
	Disks        []DiskInfo    `json:"disks"`
	Network      []NetworkInfo `json:"network"`
	Uptime       uint64        `json:"uptime"`
	LoadAverage  []float64     `json:"load_average"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type CPUInfo struct {
	UsagePercent []float64 `json:"usage_percent"`
	Cores        int       `json:"cores"`
}

type DiskInfo struct {
	Device      string  `json:"device"`
	Mountpoint  string  `json:"mountpoint"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type NetworkInfo struct {
	Name      string `json:"name"`
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}

type Service struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	PID    int32  `json:"pid"`
}

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
			"service":   "homeserver-api",
		})
	})

	// System information endpoint
	r.GET("/api/system", getSystemInfo)

	// Services endpoint (mock data for now)
	r.GET("/api/services", getServices)

	// Files endpoint (basic file listing)
	r.GET("/api/files", getFiles)

	// Network status endpoint
	r.GET("/api/network", getNetworkStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func getSystemInfo(c *gin.Context) {
	hostInfo, _ := host.Info()
	memInfo, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(time.Second, true)
	diskUsage, _ := disk.Usage("/")
	netIO, _ := net.IOCounters(true)

	var networkInfo []NetworkInfo
	for _, io := range netIO {
		networkInfo = append(networkInfo, NetworkInfo{
			Name:      io.Name,
			BytesSent: io.BytesSent,
			BytesRecv: io.BytesRecv,
		})
	}

	var disks []DiskInfo
	if diskUsage != nil {
		disks = append(disks, DiskInfo{
			Device:      "/",
			Mountpoint:  "/",
			Fstype:      diskUsage.Fstype,
			Total:       diskUsage.Total,
			Free:        diskUsage.Free,
			Used:        diskUsage.Used,
			UsedPercent: diskUsage.UsedPercent,
		})
	}

	systemInfo := SystemInfo{
		Hostname:     hostInfo.Hostname,
		Platform:     hostInfo.Platform,
		OS:           hostInfo.OS,
		Architecture: runtime.GOARCH,
		CPUCount:     runtime.NumCPU(),
		Memory: MemoryInfo{
			Total:       memInfo.Total,
			Available:   memInfo.Available,
			Used:        memInfo.Used,
			UsedPercent: memInfo.UsedPercent,
		},
		CPU: CPUInfo{
			UsagePercent: cpuPercent,
			Cores:        runtime.NumCPU(),
		},
		Disks:       disks,
		Network:     networkInfo,
		Uptime:      hostInfo.Uptime,
		LoadAverage: []float64{}, // Load average not available on all platforms
	}

	c.JSON(http.StatusOK, systemInfo)
}

func getServices(c *gin.Context) {
	// Mock service data - in a real implementation, you'd query actual system services
	services := []Service{
		{Name: "homeserver-api", Status: "running", PID: 1234},
		{Name: "nginx", Status: "running", PID: 567},
		{Name: "docker", Status: "running", PID: 890},
		{Name: "ssh", Status: "running", PID: 22},
	}

	c.JSON(http.StatusOK, services)
}

func getFiles(c *gin.Context) {
	// Basic file listing - you might want to implement proper file browsing
	path := c.DefaultQuery("path", "/")

	files, err := os.ReadDir(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fileList []map[string]interface{}
	for _, file := range files {
		info, _ := file.Info()
		fileList = append(fileList, map[string]interface{}{
			"name":     file.Name(),
			"is_dir":   file.IsDir(),
			"size":     info.Size(),
			"modified": info.ModTime(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"path":  path,
		"files": fileList,
	})
}

func getNetworkStatus(c *gin.Context) {
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
