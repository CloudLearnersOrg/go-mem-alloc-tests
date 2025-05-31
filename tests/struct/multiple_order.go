package structs

import (
	"fmt"
	"math/rand"
	model "mem-tests/model/struct"
	"mem-tests/pkg/memory"
	"time"
	"unsafe"
)

// MultiTypeStructTest tests memory efficiency across different struct types
type MultiTypeStructTest struct{}

// Name returns the name of this test
func (t *MultiTypeStructTest) Name() string {
	return "Multiple Struct Types Test"
}

// Run executes the test and returns results
func (t *MultiTypeStructTest) Run() memory.TestResult {
	result := memory.TestResult{
		Name:       t.Name(),
		OtherStats: make(map[string]any),
	}

	// Define the test cases with custom object counts
	testCases := []struct {
		name        string
		objectCount int
		optimFn     func(count int) uint64
		unoptimFn   func(count int) uint64
	}{
		{
			name:        "API Request",
			objectCount: 1000000, // High volume of API requests
			optimFn:     testAPIOptimized,
			unoptimFn:   testAPIUnoptimized,
		},
		{
			name:        "Config",
			objectCount: 10000, // Fewer config objects
			optimFn:     testConfigOptimized,
			unoptimFn:   testConfigUnoptimized,
		},
		{
			name:        "GraphQL",
			objectCount: 500000, // Medium volume of GraphQL objects
			optimFn:     testGraphQLOptimized,
			unoptimFn:   testGraphQLUnoptimized,
		},
		{
			name:        "Database Entity",
			objectCount: 250000, // Medium volume of DB entities
			optimFn:     testDBEntityOptimized,
			unoptimFn:   testDBEntityUnoptimized,
		},
	}

	// Store detailed results for each test case
	typeResults := make(map[string]map[string]interface{})
	var totalSaving uint64

	// Run all test cases
	for _, tc := range testCases {
		fmt.Printf("\n=== Testing %s Structs ===\n", tc.name)

		// Run optimized version
		fmt.Printf("\n--- Optimized %s Struct ---\n", tc.name)
		optimizedMem := tc.optimFn(tc.objectCount)

		// Force GC to clean up
		memory.CleanupAfterTest()

		// Run unoptimized version
		fmt.Printf("\n--- Unoptimized %s Struct ---\n", tc.name)
		unoptimizedMem := tc.unoptimFn(tc.objectCount)

		// Calculate savings
		memorySaved := unoptimizedMem - optimizedMem
		totalSaving += memorySaved

		var savingPct float64
		if unoptimizedMem > 0 {
			savingPct = float64(memorySaved) / float64(unoptimizedMem) * 100
		}

		// Print results
		fmt.Printf("\n--- %s Results ---\n", tc.name)
		fmt.Printf("Optimized memory: %d bytes\n", optimizedMem)
		fmt.Printf("Unoptimized memory: %d bytes\n", unoptimizedMem)
		fmt.Printf("Memory saved: %d bytes (%.2f%%)\n", memorySaved, savingPct)
		fmt.Printf("Memory saved per object: %.2f bytes\n", float64(memorySaved)/float64(tc.objectCount))

		// Store detailed results
		typeResults[tc.name] = map[string]interface{}{
			"ObjectCount":       tc.objectCount,
			"OptimizedMemory":   optimizedMem,
			"UnoptimizedMemory": unoptimizedMem,
			"MemorySaved":       memorySaved,
			"SavingPercent":     savingPct,
			"PerObjectSaving":   float64(memorySaved) / float64(tc.objectCount),
		}

		// Force GC to clean up
		memory.CleanupAfterTest()
	}

	// Store aggregated results
	result.MemoryUsed = totalSaving
	result.OtherStats["TypeResults"] = typeResults
	result.OtherStats["TotalSaving"] = totalSaving

	return result
}

// API struct test functions
func testAPIOptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.APIOptimizedStruct, count)

	// Initialize with some data
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		structs[i] = model.APIOptimizedStruct{
			RequestID:     uint64(rnd.Int63()),
			UserID:        uint64(rnd.Intn(1000000)),
			Timestamp:     time.Now().UnixNano(),
			SessionID:     uint64(rnd.Int63()),
			StatusCode:    200,
			Latency:       float32(rnd.Float64() * 100),
			APIVersion:    1,
			Method:        byte('G'), // GET
			Authenticated: true,
			Cached:        false,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample API struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

func testAPIUnoptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.APIUnoptimizedStruct, count)

	// Initialize with same data
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		structs[i] = model.APIUnoptimizedStruct{
			Method:        byte('G'), // GET
			Authenticated: true,
			UserID:        uint64(rnd.Intn(1000000)),
			Cached:        false,
			StatusCode:    200,
			APIVersion:    1,
			RequestID:     uint64(rnd.Int63()),
			Timestamp:     time.Now().UnixNano(),
			SessionID:     uint64(rnd.Int63()),
			Latency:       float32(rnd.Float64() * 100),
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample API struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

// Config struct test functions
func testConfigOptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.ConfigOptimizedStruct, count)

	// Initialize with some data
	envs := []string{"dev", "staging", "production"}
	for i := 0; i < count; i++ {
		structs[i] = model.ConfigOptimizedStruct{
			Name:           "app-config",
			Description:    "Main application configuration",
			Environment:    envs[i%len(envs)],
			UpdatedAt:      time.Now().Unix(),
			CreatedAt:      time.Now().Add(-24 * time.Hour).Unix(),
			MaxConnections: 100,
			Timeout:        30,
			Port:           8080,
			Debug:          false,
			Enabled:        true,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample Config struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

func testConfigUnoptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.ConfigUnoptimizedStruct, count)

	// Initialize with same data
	envs := []string{"dev", "staging", "production"}
	for i := 0; i < count; i++ {
		structs[i] = model.ConfigUnoptimizedStruct{
			Debug:          false,
			Enabled:        true,
			Port:           8080,
			Name:           "app-config",
			Timeout:        30,
			Environment:    envs[i%len(envs)],
			MaxConnections: 100,
			CreatedAt:      time.Now().Add(-24 * time.Hour).Unix(),
			Description:    "Main application configuration",
			UpdatedAt:      time.Now().Unix(),
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample Config struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

// GraphQL struct test functions
func testGraphQLOptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.GraphQLOptimizedStruct, count)

	// Initialize with some data
	operations := []string{"query", "mutation", "subscription"}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		structs[i] = model.GraphQLOptimizedStruct{
			QueryID:         fmt.Sprintf("query-%d", i),
			Operation:       operations[i%len(operations)],
			ClientID:        fmt.Sprintf("client-%d", i%1000),
			Timestamp:       time.Now().UnixNano(),
			Duration:        int64(rnd.Intn(1000)),
			Depth:           int32(rnd.Intn(10) + 1),
			ComplexityScore: float32(rnd.Float64() * 100),
			FragmentCount:   uint16(rnd.Intn(5)),
			IsMutation:      i%3 == 1,
			HasVariables:    i%2 == 0,
			Cached:          i%5 == 0,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample GraphQL struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

func testGraphQLUnoptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.GraphQLUnoptimizedStruct, count)

	// Initialize with same data
	operations := []string{"query", "mutation", "subscription"}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		structs[i] = model.GraphQLUnoptimizedStruct{
			IsMutation:      i%3 == 1,
			Cached:          i%5 == 0,
			Depth:           int32(rnd.Intn(10) + 1),
			Operation:       operations[i%len(operations)],
			HasVariables:    i%2 == 0,
			FragmentCount:   uint16(rnd.Intn(5)),
			Timestamp:       time.Now().UnixNano(),
			QueryID:         fmt.Sprintf("query-%d", i),
			ComplexityScore: float32(rnd.Float64() * 100),
			ClientID:        fmt.Sprintf("client-%d", i%1000),
			Duration:        int64(rnd.Intn(1000)),
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample GraphQL struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

// Database Entity struct test functions
func testDBEntityOptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.DBEntityOptimizedStruct, count)

	// Initialize with some data
	now := time.Now()
	rnd := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < count; i++ {
		structs[i] = model.DBEntityOptimizedStruct{
			ID:          fmt.Sprintf("user-%d", i),
			Name:        fmt.Sprintf("User %d", i),
			Email:       fmt.Sprintf("user%d@example.com", i),
			CreatedAt:   now.Add(-time.Duration(rnd.Intn(10000)) * time.Hour),
			UpdatedAt:   now,
			LastLoginAt: now.Add(-time.Duration(rnd.Intn(100)) * time.Hour),
			LoginCount:  int32(rnd.Intn(100)),
			Status:      int32(rnd.Intn(3)),
			AccessLevel: uint16(rnd.Intn(5)),
			IsActive:    true,
			IsAdmin:     i%50 == 0, // 2% are admins
			HasMFA:      i%3 == 0,  // 33% have MFA
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample DB Entity struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}

func testDBEntityUnoptimized(count int) uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.DBEntityUnoptimizedStruct, count)

	// Initialize with same data
	now := time.Now()
	rnd := rand.New(rand.NewSource(now.UnixNano()))

	for i := 0; i < count; i++ {
		structs[i] = model.DBEntityUnoptimizedStruct{
			IsActive:    true,
			IsAdmin:     i%50 == 0, // 2% are admins
			AccessLevel: uint16(rnd.Intn(5)),
			Email:       fmt.Sprintf("user%d@example.com", i),
			Status:      int32(rnd.Intn(3)),
			LastLoginAt: now.Add(-time.Duration(rnd.Intn(100)) * time.Hour),
			HasMFA:      i%3 == 0, // 33% have MFA
			ID:          fmt.Sprintf("user-%d", i),
			Name:        fmt.Sprintf("User %d", i),
			LoginCount:  int32(rnd.Intn(100)),
			CreatedAt:   now.Add(-time.Duration(rnd.Intn(10000)) * time.Hour),
			UpdatedAt:   now,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem

	// Prevent optimizer from removing our structs
	fmt.Printf("Sample DB Entity struct size: %d bytes\n", unsafe.Sizeof(structs[0]))
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n", memUsed, float64(memUsed)/(1024*1024))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(count))

	return memUsed
}
