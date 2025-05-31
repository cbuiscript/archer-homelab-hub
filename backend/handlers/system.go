package handlers

import (
	"net/http"
	"runtime"
	"time"

	"homeserver/types"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// GetSystemInfo handles the system information endpoint
func GetSystemInfo(c *gin.Context) {
	hostInfo, _ := host.Info()
	memInfo, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(time.Second, true)
	diskUsage, _ := disk.Usage("/")
	netIO, _ := net.IOCounters(true)

	var networkInfo []types.NetworkInfo
	for _, io := range netIO {
		networkInfo = append(networkInfo, types.NetworkInfo{
			Name:      io.Name,
			BytesSent: io.BytesSent,
			BytesRecv: io.BytesRecv,
		})
	}

	var disks []types.DiskInfo
	if diskUsage != nil {
		disks = append(disks, types.DiskInfo{
			Device:      "/",
			Mountpoint:  "/",
			Fstype:      diskUsage.Fstype,
			Total:       diskUsage.Total,
			Free:        diskUsage.Free,
			Used:        diskUsage.Used,
			UsedPercent: diskUsage.UsedPercent,
		})
	}

	systemInfo := types.SystemInfo{
		Hostname:     hostInfo.Hostname,
		Platform:     hostInfo.Platform,
		OS:           hostInfo.OS,
		Architecture: runtime.GOARCH,
		CPUCount:     runtime.NumCPU(),
		Memory: types.MemoryInfo{
			Total:       memInfo.Total,
			Available:   memInfo.Available,
			Used:        memInfo.Used,
			UsedPercent: memInfo.UsedPercent,
		},
		CPU: types.CPUInfo{
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
