package memory

import (
	"runtime"
	"runtime/debug"
)

// PrepareMemoryTest disables GC and runs a collection to start with a clean state
func PrepareMemoryTest() uint64 {
	// Disable GC during the test to get more accurate memory measurements
	debug.SetGCPercent(-1)

	// Force a GC to start with a clean state
	runtime.GC()

	// Return initial memory usage
	return Usage()
}

// CleanupAfterTest re-enables GC and runs a collection
func CleanupAfterTest() {
	// Re-enable GC
	debug.SetGCPercent(100)

	// Force GC to clean up
	runtime.GC()
}
