package structs

import (
	"fmt"
	"math/rand"
	model "mem-tests/model/struct"
	"mem-tests/pkg/memory"
	"time"
	"unsafe"
)

// StructBigTest tests memory efficiency with different struct field orders for large structs
type StructBigTest struct{}

// Name returns the name of this test
func (t *StructBigTest) Name() string {
	return "Large Struct Field Order Test"
}

// Run executes the test and returns results
func (t *StructBigTest) Run() memory.TestResult {
	result := memory.TestResult{
		Name:       t.Name(),
		OtherStats: make(map[string]any),
	}

	// Analyze struct layouts
	analyzeLargeStructLayout(&result)

	// Test with optimized structs
	fmt.Println("\n=== Testing LargeOptimizedStruct (largest to smallest) ===")
	optimizedMem := testLargeOptimizedStructs()

	// Force GC to clean up
	memory.CleanupAfterTest()

	// Test with unoptimized structs
	fmt.Println("\n=== Testing LargeUnoptimizedStruct (mixed order) ===")
	unoptimizedMem := testLargeUnoptimizedStructs()

	// Store results
	result.MemoryUsed = unoptimizedMem - optimizedMem
	result.OtherStats["OptimizedMemory"] = optimizedMem
	result.OtherStats["UnoptimizedMemory"] = unoptimizedMem
	result.PerObjectSize = float64(result.MemoryUsed) / float64(numObjects)

	if unoptimizedMem > optimizedMem {
		result.OtherStats["MemorySavingPercent"] = float64(unoptimizedMem-optimizedMem) / float64(unoptimizedMem) * 100
	} else {
		result.OtherStats["MemorySavingPercent"] = 0.0
	}

	// Print struct sizes
	result.OtherStats["OptimizedStructSize"] = unsafe.Sizeof(model.LargeOptimizedStruct{})
	result.OtherStats["UnoptimizedStructSize"] = unsafe.Sizeof(model.LargeUnoptimizedStruct{})

	return result
}

func analyzeLargeStructLayout(result *memory.TestResult) {
	fmt.Println("\n=== Large Struct Layout Analysis ===")

	// Analyze LargeOptimizedStruct
	optimized := model.LargeOptimizedStruct{}
	optSize := unsafe.Sizeof(optimized)
	fmt.Printf("LargeOptimizedStruct size: %d bytes (%s)\n", optSize, memory.FormatBytes(uint64(optSize)))

	// Analyze LargeUnoptimizedStruct
	unoptimized := model.LargeUnoptimizedStruct{}
	unoptSize := unsafe.Sizeof(unoptimized)
	fmt.Printf("LargeUnoptimizedStruct size: %d bytes (%s)\n", unoptSize, memory.FormatBytes(uint64(unoptSize)))

	// Calculate theoretical memory difference for all objects
	sizeDiff := unoptSize - optSize
	if sizeDiff > 0 {
		fmt.Printf("\nTheoretical memory waste per struct: %d bytes\n", sizeDiff)
		totalWaste := sizeDiff * uintptr(numObjects)
		fmt.Printf("Theoretical total memory waste for %d objects: %d bytes (%s)\n",
			numObjects, totalWaste, memory.FormatBytes(uint64(totalWaste)))
	} else {
		fmt.Printf("\nNo theoretical memory waste detected from struct layout\n")
	}
}

func testLargeOptimizedStructs() uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.LargeOptimizedStruct, numObjects)

	// Initialize each struct with realistic data
	now := time.Now()
	rnd := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < numObjects; i++ {
		structs[i] = model.LargeOptimizedStruct{
			CreatedAt:         now.Add(-time.Duration(rnd.Intn(3600)) * time.Second),
			UpdatedAt:         now,
			TransactionID:     uint64(rnd.Int63()),
			UserID:            uint64(rnd.Intn(1000000)),
			OrderID:           uint64(rnd.Int63()),
			RequestTimestamp:  now.Add(-time.Duration(rnd.Intn(500)) * time.Millisecond).UnixNano(),
			ResponseTimestamp: now.UnixNano(),
			StatusCode:        200,
			ResponseCode:      0,
			RequestCount:      int32(1 + rnd.Intn(5)),
			RetryAttempts:     int32(rnd.Intn(3)),
			ServiceTime:       float32(rnd.Float64() * 100),
			CPUTime:           float32(rnd.Float64() * 50),
			ErrorCode:         0,
			ProtocolVersion:   2,
			ServerRegion:      uint16(rnd.Intn(10)),
			IsSuccess:         true,
			IsRetry:           false,
			IsCached:          rnd.Float32() < 0.3, // 30% cache hit rate
			Priority:          uint8(rnd.Intn(5)),
			CompressionLevel:  uint8(rnd.Intn(10)),
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem
	fmt.Printf("Memory used: %d bytes (%s)\n", memUsed, memory.FormatBytes(memUsed))
	fmt.Printf("Memory per struct: %.2f bytes (%s)\n",
		float64(memUsed)/float64(numObjects),
		memory.FormatBytes(uint64(float64(memUsed)/float64(numObjects))))

	// Prevent optimizer from removing our structs before measurements
	fmt.Printf("Sample value: %v\n", structs[0].TransactionID)

	return memUsed
}

func testLargeUnoptimizedStructs() uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.LargeUnoptimizedStruct, numObjects)

	// Initialize each struct with the same realistic data as optimized version
	now := time.Now()
	rnd := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < numObjects; i++ {
		structs[i] = model.LargeUnoptimizedStruct{
			IsSuccess:         true,
			Priority:          uint8(rnd.Intn(5)),
			UserID:            uint64(rnd.Intn(1000000)),
			IsRetry:           false,
			StatusCode:        200,
			ServiceTime:       float32(rnd.Float64() * 100),
			ErrorCode:         0,
			CreatedAt:         now.Add(-time.Duration(rnd.Intn(3600)) * time.Second),
			TransactionID:     uint64(rnd.Int63()),
			IsCached:          rnd.Float32() < 0.3, // 30% cache hit rate
			ResponseCode:      0,
			ProtocolVersion:   2,
			ServerRegion:      uint16(rnd.Intn(10)),
			UpdatedAt:         now,
			OrderID:           uint64(rnd.Int63()),
			RequestTimestamp:  now.Add(-time.Duration(rnd.Intn(500)) * time.Millisecond).UnixNano(),
			ResponseTimestamp: now.UnixNano(),
			RequestCount:      int32(1 + rnd.Intn(5)),
			RetryAttempts:     int32(rnd.Intn(3)),
			CPUTime:           float32(rnd.Float64() * 50),
			CompressionLevel:  uint8(rnd.Intn(10)),
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem
	fmt.Printf("Memory used: %d bytes (%s)\n", memUsed, memory.FormatBytes(memUsed))
	fmt.Printf("Memory per struct: %.2f bytes (%s)\n",
		float64(memUsed)/float64(numObjects),
		memory.FormatBytes(uint64(float64(memUsed)/float64(numObjects))))

	// Prevent optimizer from removing our structs before measurements
	fmt.Printf("Sample value: %v\n", structs[0].TransactionID)

	return memUsed
}
