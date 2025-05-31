package structs

import (
	"fmt"
	"mem-tests/pkg/memory"
)

// StructAnalyzer provides utility functions for analyzing struct memory usage
type StructAnalyzer struct{}

// CalculateMemorySavings computes memory savings statistics between optimized and unoptimized versions
func (a *StructAnalyzer) CalculateMemorySavings(result *memory.TestResult, optimizedMem, unoptimizedMem uint64, optimizedSize, unoptimizedSize uintptr) {
	result.MemoryUsed = unoptimizedMem - optimizedMem
	result.OtherStats["OptimizedMemory"] = optimizedMem
	result.OtherStats["UnoptimizedMemory"] = unoptimizedMem
	result.PerObjectSize = float64(result.MemoryUsed) / float64(numObjects)

	if unoptimizedMem > optimizedMem {
		result.OtherStats["MemorySavingPercent"] = float64(unoptimizedMem-optimizedMem) / float64(unoptimizedMem) * 100
	} else {
		result.OtherStats["MemorySavingPercent"] = 0.0
	}

	result.OtherStats["OptimizedStructSize"] = optimizedSize
	result.OtherStats["UnoptimizedStructSize"] = unoptimizedSize
}

// AnalyzeStructLayout provides a generic function to analyze any struct layout
func (a *StructAnalyzer) AnalyzeStructLayout(testName string, optimizedSize, unoptimizedSize uintptr, objectCount int) {
	fmt.Printf("\n=== %s Layout Analysis ===\n", testName)
	fmt.Printf("Optimized struct size: %d bytes\n", optimizedSize)
	fmt.Printf("Unoptimized struct size: %d bytes\n", unoptimizedSize)

	// Calculate theoretical memory difference for all objects
	sizeDiff := unoptimizedSize - optimizedSize
	if sizeDiff > 0 {
		fmt.Printf("\nTheoretical memory waste per struct: %d bytes\n", sizeDiff)
		fmt.Printf("Theoretical total memory waste for %d objects: %d bytes (%.2f MB)\n",
			objectCount,
			sizeDiff*uintptr(objectCount),
			float64(sizeDiff*uintptr(objectCount))/(1024*1024))
	} else {
		fmt.Printf("\nNo theoretical memory waste detected from struct layout\n")
	}
}

// PrintFieldOffsets is a helper to print field offsets for detailed analysis
func (a *StructAnalyzer) PrintFieldOffsets(fieldName string, offset uintptr) {
	fmt.Printf("  %s: %d\n", fieldName, offset)
}

// PrintMemoryStats prints standard memory statistics after a test
func (a *StructAnalyzer) PrintMemoryStats(memUsed uint64, objectCount int, isMB bool) {
	fmt.Printf("Memory used: %d bytes (%s)\n", memUsed, memory.FormatBytes(memUsed))
	perObjectBytes := float64(memUsed) / float64(objectCount)
	fmt.Printf("Memory per struct: %.2f bytes (%s)\n",
		perObjectBytes,
		memory.FormatBytes(uint64(perObjectBytes)))
}
