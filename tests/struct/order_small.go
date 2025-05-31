package structs

import (
	"fmt"
	model "mem-tests/model/struct"
	"mem-tests/pkg/memory"
	"unsafe"
)

// StructOrderTest tests memory efficiency with different struct field orders
type StructOrderTest struct{}

// Name returns the name of this test
func (t *StructOrderTest) Name() string {
	return "Struct Field Order Test"
}

// Run executes the test and returns results
func (t *StructOrderTest) Run() memory.TestResult {
	result := memory.TestResult{
		Name:       t.Name(),
		OtherStats: make(map[string]any),
	}

	// Analyze struct layouts
	analyzeSmallStructLayout(&result)

	// Test with optimized structs
	fmt.Println("\n=== Testing OptimizedStruct (largest to smallest) ===")
	optimizedMem := testSmallOptimizedStructs()

	// Force GC to clean up
	memory.CleanupAfterTest()

	// Test with unoptimized structs
	fmt.Println("\n=== Testing UnoptimizedStruct (smallest to largest) ===")
	unoptimizedMem := testSmallUnoptimizedStructs()

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
	result.OtherStats["OptimizedStructSize"] = unsafe.Sizeof(model.OptimizedStruct{})
	result.OtherStats["UnoptimizedStructSize"] = unsafe.Sizeof(model.UnoptimizedStruct{})

	return result
}

func analyzeSmallStructLayout(result *memory.TestResult) {
	fmt.Println("\n=== Small Struct Layout Analysis ===")

	// Analyze OptimizedStruct
	optimized := model.OptimizedStruct{}
	optSize := unsafe.Sizeof(optimized)
	fmt.Printf("OptimizedStruct size: %d bytes (%s)\n", optSize, memory.FormatBytes(uint64(optSize)))

	fmt.Println("Field offsets (bytes from start of struct):")
	fmt.Printf("  Int64Field:  %d\n", unsafe.Offsetof(optimized.Int64Field))
	fmt.Printf("  Int64FieldB: %d\n", unsafe.Offsetof(optimized.Int64FieldB))
	fmt.Printf("  Int32Field:  %d\n", unsafe.Offsetof(optimized.Int32Field))
	fmt.Printf("  Int32FieldB: %d\n", unsafe.Offsetof(optimized.Int32FieldB))
	fmt.Printf("  Int16Field:  %d\n", unsafe.Offsetof(optimized.Int16Field))
	fmt.Printf("  Int16FieldB: %d\n", unsafe.Offsetof(optimized.Int16FieldB))
	fmt.Printf("  Int8Field:   %d\n", unsafe.Offsetof(optimized.Int8Field))
	fmt.Printf("  BoolField:   %d\n", unsafe.Offsetof(optimized.BoolField))

	// Analyze UnoptimizedStruct
	unoptimized := model.UnoptimizedStruct{}
	unoptSize := unsafe.Sizeof(unoptimized)
	fmt.Printf("\nUnoptimizedStruct size: %d bytes (%s)\n", unoptSize, memory.FormatBytes(uint64(unoptSize)))

	fmt.Println("Field offsets (bytes from start of struct):")
	fmt.Printf("  BoolField:   %d\n", unsafe.Offsetof(unoptimized.BoolField))
	fmt.Printf("  Int8Field:   %d\n", unsafe.Offsetof(unoptimized.Int8Field))
	fmt.Printf("  Int16Field:  %d\n", unsafe.Offsetof(unoptimized.Int16Field))
	fmt.Printf("  Int32Field:  %d\n", unsafe.Offsetof(unoptimized.Int32Field))
	fmt.Printf("  BoolFieldB:  %d\n", unsafe.Offsetof(unoptimized.BoolFieldB))
	fmt.Printf("  Int64Field:  %d\n", unsafe.Offsetof(unoptimized.Int64Field))
	fmt.Printf("  Int16FieldB: %d\n", unsafe.Offsetof(unoptimized.Int16FieldB))
	fmt.Printf("  Int32FieldB: %d\n", unsafe.Offsetof(unoptimized.Int32FieldB))
	fmt.Printf("  Int64FieldB: %d\n", unsafe.Offsetof(unoptimized.Int64FieldB))

	// Calculate theoretical memory difference
	sizeDiff := unoptSize - optSize
	if sizeDiff > 0 {
		fmt.Printf("\nTheoretical memory waste per struct: %d bytes\n", sizeDiff)
		totalWaste := sizeDiff * uintptr(numObjects)
		fmt.Printf("Theoretical total memory waste for %d objects: %d bytes (%s)\n",
			numObjects, totalWaste, memory.FormatBytes(uint64(totalWaste)))
	}
}

func testSmallOptimizedStructs() uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.OptimizedStruct, numObjects)

	// Initialize each struct with the same data
	for i := 0; i < numObjects; i++ {
		structs[i] = model.OptimizedStruct{
			Int64Field:  123456789,
			Int64FieldB: 987654321,
			Int32Field:  123456,
			Int32FieldB: 654321,
			Int16Field:  1234,
			Int16FieldB: 4321,
			Int8Field:   123,
			BoolField:   true,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem
	fmt.Printf("Memory used: %d bytes (%s)\n", memUsed, memory.FormatBytes(memUsed))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(numObjects))

	// Prevent optimizer from removing our structs before measurements
	fmt.Printf("Sample value: %v\n", structs[0].Int64Field)

	return memUsed
}

func testSmallUnoptimizedStructs() uint64 {
	startMem := memory.Usage()

	// Create a slice to hold all the structs
	structs := make([]model.UnoptimizedStruct, numObjects)

	// Initialize each struct with the same data
	for i := 0; i < numObjects; i++ {
		structs[i] = model.UnoptimizedStruct{
			BoolField:   true,
			Int8Field:   123,
			Int16Field:  1234,
			Int32Field:  123456,
			BoolFieldB:  false,
			Int64Field:  123456789,
			Int16FieldB: 4321,
			Int32FieldB: 654321,
			Int64FieldB: 987654321,
		}
	}

	endMem := memory.Usage()
	memUsed := endMem - startMem
	fmt.Printf("Memory used: %d bytes (%s)\n", memUsed, memory.FormatBytes(memUsed))
	fmt.Printf("Memory per struct: %.2f bytes\n", float64(memUsed)/float64(numObjects))

	// Prevent optimizer from removing our structs before measurements
	fmt.Printf("Sample value: %v\n", structs[0].Int64Field)

	return memUsed
}
