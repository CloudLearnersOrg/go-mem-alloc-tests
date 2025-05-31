package memory

// TestResult contains the results of a memory test
type TestResult struct {
	// Name of the test
	Name string

	// Total memory difference between optimized and unoptimized versions
	MemoryUsed uint64

	// Memory difference per object (if applicable)
	PerObjectSize float64

	// Additional test-specific statistics
	OtherStats map[string]any
}

// MemoryTest defines the interface for all memory tests
type MemoryTest interface {
	// Name returns a human-readable name for the test
	Name() string

	// Run executes the test and returns the results
	Run() TestResult
}
