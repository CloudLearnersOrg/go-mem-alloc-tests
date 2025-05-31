.PHONY: run list visualize clean test all help

# Default target
all: test visualize

# Run all memory tests
run:
	go run main.go

# List available tests
list:
	go run main.go -list

# Run specific test(s)
test:
	go run main.go -test=$(TEST)

# Run tests and visualize results
visualize: run
	go run main.go -test=$(TEST) -viz -format=$(FORMAT)

# Generate visualizations in all formats
report:
	go run main.go -viz -format=terminal
	go run main.go -viz -format=html
	go run main.go -viz -format=png

# Clean generated files
clean:
	rm -f memory_test_results.*

# Help
help:
	@echo "Memory Allocation Tests"
	@echo ""
	@echo "Usage:"
	@echo "  make run              - Run all memory tests"
	@echo "  make list             - List available tests"
	@echo "  make test TEST=name   - Run specific test(s)"
	@echo "  make visualize        - Visualize test results"
	@echo "  make visualize TEST=name FORMAT=html - Run specific test with HTML output"
	@echo "  make report           - Generate all formats of reports"
	@echo "  make clean            - Remove generated files"
	@echo ""
	@echo "Examples:"
	@echo "  make test TEST=struct-small        - Run small struct test"
	@echo "  make test TEST=struct-big          - Run big struct test"
	@echo "  make visualize FORMAT=png TEST=all - Visualize all tests in PNG format"