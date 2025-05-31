package types

// SystemInfo represents the overall system information
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

// MemoryInfo represents memory usage information
type MemoryInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

// CPUInfo represents CPU usage information
type CPUInfo struct {
	UsagePercent []float64 `json:"usage_percent"`
	Cores        int       `json:"cores"`
}

// DiskInfo represents disk usage information
type DiskInfo struct {
	Device      string  `json:"device"`
	Mountpoint  string  `json:"mountpoint"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

// NetworkInfo represents network interface information
type NetworkInfo struct {
	Name      string `json:"name"`
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}

// Service represents a system service
type Service struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	PID    int32  `json:"pid"`
}

// FileInfo represents file information
type FileInfo struct {
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	IsDirectory  bool   `json:"is_directory"`
	ModifiedTime string `json:"modified_time"`
	Permissions  string `json:"permissions"`
}

// NetworkStatus represents network connectivity status
type NetworkStatus struct {
	IsConnected  bool   `json:"is_connected"`
	PublicIP     string `json:"public_ip,omitempty"`
	LocalIP      string `json:"local_ip"`
	Gateway      string `json:"gateway,omitempty"`
	DNS          string `json:"dns,omitempty"`
	Latency      string `json:"latency,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	Uptime    string `json:"uptime"`
}
