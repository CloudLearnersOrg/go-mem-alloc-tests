# Memory Allocation Tests

A toolkit for testing and visualizing Go memory allocation patterns with a focus on struct field ordering and memory efficiency.

## Overview

This project provides a framework for measuring memory usage across different Go data structures and memory management techniques. It allows for:

- Testing different struct field orderings to minimize padding
- Measuring memory consumption of various data structures
- Visualizing memory usage differences
- Extensible architecture for adding new tests

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/mem-alloc-tests.git
cd mem-alloc-tests
```

## Usage

### Running Tests

Run all tests:

```bash
make run
```

List available tests:

```bash
make list
```

Run a specific test:

```bash
make test TEST=struct-small
```

### Visualizing Results

Generate visualizations of test results:

```bash
make visualize FORMAT=html
```

Available visualization formats:
- `terminal` - ASCII visualization in the terminal
- `html` - HTML report with charts

### Generate All Reports

Generate reports in all available formats:

```bash
make report
```

## GitHub Pages Integration

Test results are automatically exported to the `docs/` directory in a format suitable for GitHub Pages:

- Individual test results are saved as separate HTML files
- Historical results are preserved by date
- An index page links to all available test results

To view the results:
1. Enable GitHub Pages for your repository (Settings -> Pages)
2. Set the source to the `main` branch and `/docs` folder
3. Access your results at `https://yourusername.github.io/mem-alloc-tests/`

You can also view results locally by opening `docs/index.html` in your browser.

## Example Tests

### Struct Field Ordering Test

Tests memory usage with optimized vs. unoptimized struct field ordering:

```bash
make test TEST=struct-small
```

### Large Struct Field Ordering Test

Tests memory usage with large structs similar to those found in high-throughput services:

```bash
make test TEST=struct-big
```

## Interpreting Results

The test results show:

1. **Memory usage** - Total memory used by each test case
2. **Memory difference** - How much memory is saved by optimization
3. **Per-object size** - Memory used per object
4. **Memory saving percentage** - Percentage of memory saved

## Why Memory Optimization Matters

For services handling millions of objects, small memory savings per object can add up to significant reductions in overall memory usage. This can lead to:

- Reduced cloud infrastructure costs
- Better cache utilization
- Fewer garbage collection pauses
- Lower latency
- Higher throughput

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
