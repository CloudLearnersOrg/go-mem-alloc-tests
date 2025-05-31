package main

import (
	"flag"
	"fmt"
	"mem-tests/pkg/memory"
	"mem-tests/pkg/visualizer"
	structs "mem-tests/tests/struct"
	"os"
	"strings"
)

func main() {
	// Define available tests
	tests := map[string]memory.MemoryTest{
		"struct-small": &structs.StructOrderTest{},
		"struct-big":   &structs.StructBigTest{},
		"struct-multi": &structs.MultiTypeStructTest{},
		// Add more tests here as you create them
	}

	// Parse command line flags
	var listTests bool
	var testName string
	var visualize bool
	var outputFormat string

	flag.BoolVar(&listTests, "list", false, "List available tests")
	flag.StringVar(&testName, "test", "", "Name of test to run (comma separated for multiple)")
	flag.BoolVar(&visualize, "viz", false, "Visualize test results")
	flag.StringVar(&outputFormat, "format", "stdout", "Output format: stdout, html, png")
	flag.Parse()

	// List available tests if requested
	if listTests {
		fmt.Println("Available tests:")
		for name, test := range tests {
			fmt.Printf("  %s - %s\n", name, test.Name())
		}
		return
	}

	var results []memory.TestResult

	// Run all tests if none specified
	if testName == "" {
		fmt.Println("Running all tests...")
		results = runAllTests(tests)
	} else {
		// Run specified tests
		testNames := strings.Split(testName, ",")
		for _, name := range testNames {
			name = strings.TrimSpace(name)
			test, exists := tests[name]
			if !exists {
				fmt.Printf("Test '%s' not found. Use -list to see available tests.\n", name)
				os.Exit(1)
			}

			fmt.Printf("\n=== Running test: %s ===\n", test.Name())
			initialMem := memory.PrepareMemoryTest()
			fmt.Printf("Initial memory usage: %d bytes\n", initialMem)

			result := test.Run()
			results = append(results, result)

			fmt.Println("\n=== Results ===")
			printTestResult(result)

			memory.CleanupAfterTest()
		}
	}

	// Visualize results if requested
	if visualize && len(results) > 0 {
		v := visualizer.New(outputFormat)
		if err := v.Visualize(results); err != nil {
			fmt.Printf("Error visualizing results: %v\n", err)
		}
	}
}

func runAllTests(tests map[string]memory.MemoryTest) []memory.TestResult {
	var results []memory.TestResult

	for _, test := range tests {
		fmt.Printf("\n\n=== Running test: %s ===\n", test.Name())
		initialMem := memory.PrepareMemoryTest()
		fmt.Printf("Initial memory usage: %d bytes\n", initialMem)

		result := test.Run()
		results = append(results, result)

		fmt.Println("\n=== Results ===")
		printTestResult(result)

		memory.CleanupAfterTest()
	}

	return results
}

func printTestResult(result memory.TestResult) {
	fmt.Printf("Test: %s\n", result.Name)
	fmt.Printf("Memory difference: %d bytes\n", result.MemoryUsed)

	if result.PerObjectSize > 0 {
		fmt.Printf("Memory per object: %.2f bytes\n", result.PerObjectSize)
	}

	// Check for multi-type test results
	if typeResults, ok := result.OtherStats["TypeResults"].(map[string]map[string]interface{}); ok {
		fmt.Println("\nDetailed results by type:")

		for typeName, typeData := range typeResults {
			fmt.Printf("\n  %s:\n", typeName)

			if count, ok := typeData["ObjectCount"].(int); ok {
				fmt.Printf("    Object count: %d\n", count)
			}

			if optMem, ok := typeData["OptimizedMemory"].(uint64); ok {
				fmt.Printf("    Optimized memory: %d bytes\n", optMem)
			}

			if unoptMem, ok := typeData["UnoptimizedMemory"].(uint64); ok {
				fmt.Printf("    Unoptimized memory: %d bytes\n", unoptMem)
			}

			if saved, ok := typeData["MemorySaved"].(uint64); ok {
				fmt.Printf("    Memory saved: %d bytes\n", saved)
			}

			if pct, ok := typeData["SavingPercent"].(float64); ok {
				fmt.Printf("    Saving: %.2f%%\n", pct)
			}
		}

		if totalSaving, ok := result.OtherStats["TotalSaving"].(uint64); ok {
			fmt.Printf("\nTotal memory saving: %d bytes\n", totalSaving)
		}

		return
	}

	// Print standard stats
	for key, value := range result.OtherStats {
		switch v := value.(type) {
		case float64:
			fmt.Printf("%s: %.2f\n", key, v)
		default:
			fmt.Printf("%s: %v\n", key, v)
		}
	}
}
