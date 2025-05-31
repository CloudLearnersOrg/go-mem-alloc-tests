package memory

import (
	"fmt"
	"runtime"
)

// Usage returns the current memory allocation in bytes
func Usage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// UsageMB returns the current memory allocation in megabytes as a float
func UsageMB() float64 {
	return float64(Usage()) / (1024 * 1024)
}

// FormatBytes converts bytes to a human-readable string with appropriate unit (B, KB, MB, GB)
func FormatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// PrintUsage prints current memory usage in a human-readable format
func PrintUsage(label string) {
	bytes := Usage()
	fmt.Printf("%s: %s (%d bytes)\n", label, FormatBytes(bytes), bytes)
}
