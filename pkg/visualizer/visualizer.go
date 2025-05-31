package visualizer

import (
	"bytes"
	"fmt"
	"mem-tests/pkg/memory"
	"os"
	"strings"
)

// Visualizer is responsible for creating visual representations of test results
type Visualizer interface {
	// Visualize takes test results and renders them in the visualizer's format
	Visualize(results []memory.TestResult) error
}

// Factory function to create a visualizer based on the specified format
func New(format string) Visualizer {
	switch strings.ToLower(format) {
	case "stdout", "terminal":
		return &TerminalVisualizer{}
	case "html", "":
		// Use HTML as the default if no format is specified
		return &HTMLVisualizer{}
	default:
		fmt.Printf("Unknown format %q, defaulting to HTML\n", format)
		return &HTMLVisualizer{}
	}
}

// TerminalVisualizer displays results in the terminal with ASCII bars
type TerminalVisualizer struct{}

// Visualize implements the Visualizer interface for terminal output
func (t *TerminalVisualizer) Visualize(results []memory.TestResult) error {
	if len(results) == 0 {
		return fmt.Errorf("no results to visualize")
	}

	// Find the max value for scaling
	var maxVal uint64
	for _, r := range results {
		if v, ok := r.OtherStats["OptimizedMemory"].(uint64); ok && v > maxVal {
			maxVal = v
		}
		if v, ok := r.OtherStats["UnoptimizedMemory"].(uint64); ok && v > maxVal {
			maxVal = v
		}
	}

	// Terminal width (adjust if needed)
	termWidth := 80
	maxBarWidth := termWidth - 40 // Leave space for labels

	fmt.Println("\n=== Memory Usage Visualization ===")
	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Printf("%-20s %-15s %s\n", "Test", "Memory (bytes)", "Bar")
	fmt.Println(strings.Repeat("-", termWidth))

	for _, r := range results {
		// Check for multiple type results
		if typeResults, ok := r.OtherStats["TypeResults"].(map[string]map[string]interface{}); ok {
			fmt.Printf("=== %s ===\n", r.Name)

			// Find max value for scaling within this test
			var maxTypeVal uint64
			for _, typeData := range typeResults {
				if v, ok := typeData["UnoptimizedMemory"].(uint64); ok && v > maxTypeVal {
					maxTypeVal = v
				}
			}

			// Display each type result
			for typeName, typeData := range typeResults {
				fmt.Printf("--- %s ---\n", typeName)

				// Display optimized struct memory
				if optMem, ok := typeData["OptimizedMemory"].(uint64); ok {
					barLength := int(float64(optMem) / float64(maxTypeVal) * float64(maxBarWidth))
					bar := strings.Repeat("█", barLength)
					fmt.Printf("%-20s %-15d %s\n", "Optimized", optMem, bar)
				}

				// Display unoptimized struct memory
				if unoptMem, ok := typeData["UnoptimizedMemory"].(uint64); ok {
					barLength := int(float64(unoptMem) / float64(maxTypeVal) * float64(maxBarWidth))
					bar := strings.Repeat("█", barLength)
					fmt.Printf("%-20s %-15d %s\n", "Unoptimized", unoptMem, bar)
				}

				// Add a memory saving percentage
				if savingPct, ok := typeData["SavingPercent"].(float64); ok {
					fmt.Printf("Memory Saving: %.2f%%\n", savingPct)
				}

				fmt.Println(strings.Repeat("-", termWidth))
			}

			if totalSaving, ok := r.OtherStats["TotalSaving"].(uint64); ok {
				fmt.Printf("Total Memory Saving: %d bytes\n", totalSaving)
			}

			continue
		}

		// Standard visualization for simple tests
		if optMem, ok := r.OtherStats["OptimizedMemory"].(uint64); ok {
			barLength := int(float64(optMem) / float64(maxVal) * float64(maxBarWidth))
			bar := strings.Repeat("█", barLength)
			fmt.Printf("%-20s %-15d %s\n", r.Name+" (Opt)", optMem, bar)
		}

		if unoptMem, ok := r.OtherStats["UnoptimizedMemory"].(uint64); ok {
			barLength := int(float64(unoptMem) / float64(maxVal) * float64(maxBarWidth))
			bar := strings.Repeat("█", barLength)
			fmt.Printf("%-20s %-15d %s\n", r.Name+" (Unopt)", unoptMem, bar)
		}

		if savingPct, ok := r.OtherStats["MemorySavingPercent"].(float64); ok {
			fmt.Printf("%-20s %.2f%%\n", "Memory Saving:", savingPct)
		}

		fmt.Println(strings.Repeat("-", termWidth))
	}

	return nil
}

// HTMLVisualizer generates HTML visualizations with JavaScript charts
type HTMLVisualizer struct{}

// Visualize implements the Visualizer interface for HTML output
func (h *HTMLVisualizer) Visualize(results []memory.TestResult) error {
	fmt.Println("Generating HTML visualization...")

	var html bytes.Buffer

	// HTML header and CSS
	html.WriteString(`<!DOCTYPE html>
<html>
<head>
    <title>Memory Allocation Test Results</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        .chart-container { height: 400px; margin-bottom: 40px; }
        table { border-collapse: collapse; width: 100%; margin: 20px 0; }
        th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
        th { background-color: #f2f2f2; }
        tr:nth-child(even) { background-color: #f9f9f9; }
        h1, h2 { color: #333; }
        .memory-cell { white-space: nowrap; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Memory Allocation Test Results</h1>
`)

	// Generate charts and tables for each test
	for i, r := range results {
		// Add chart container
		html.WriteString(fmt.Sprintf(`
        <h2>%s</h2>
        <div class="chart-container">
            <canvas id="chart%d"></canvas>
        </div>
`, r.Name, i))

		// Generate table for results
		html.WriteString(`
        <table>
            <tr>
                <th>Type</th>
                <th>Optimized Memory</th>
                <th>Unoptimized Memory</th>
                <th>Memory Saving</th>
                <th>Saving Percentage</th>
            </tr>
`)

		// Check for multiple type results
		if typeResults, ok := r.OtherStats["TypeResults"].(map[string]map[string]interface{}); ok {
			for typeName, typeData := range typeResults {
				optMem, _ := typeData["OptimizedMemory"].(uint64)
				unoptMem, _ := typeData["UnoptimizedMemory"].(uint64)
				saved, _ := typeData["MemorySaved"].(uint64)
				savingPct, _ := typeData["SavingPercent"].(float64)

				html.WriteString(fmt.Sprintf(`
            <tr>
                <td>%s</td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td>%.2f%%</td>
            </tr>
`, typeName,
					memory.FormatBytes(optMem), optMem,
					memory.FormatBytes(unoptMem), unoptMem,
					memory.FormatBytes(saved), saved,
					savingPct))
			}

			// Add total row
			if totalSaving, ok := r.OtherStats["TotalSaving"].(uint64); ok {
				html.WriteString(fmt.Sprintf(`
            <tr style="font-weight: bold;">
                <td>Total</td>
                <td>-</td>
                <td>-</td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td>-</td>
            </tr>
`, memory.FormatBytes(totalSaving), totalSaving))
			}
		} else {
			// Standard results
			optMem, _ := r.OtherStats["OptimizedMemory"].(uint64)
			unoptMem, _ := r.OtherStats["UnoptimizedMemory"].(uint64)
			savingPct, _ := r.OtherStats["MemorySavingPercent"].(float64)

			html.WriteString(fmt.Sprintf(`
            <tr>
                <td>Standard</td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td class="memory-cell">%s <span class="bytes-value">(%d bytes)</span></td>
                <td>%.2f%%</td>
            </tr>
`, memory.FormatBytes(optMem), optMem,
				memory.FormatBytes(unoptMem), unoptMem,
				memory.FormatBytes(r.MemoryUsed), r.MemoryUsed,
				savingPct))
		}

		html.WriteString(`
        </table>
`)
	}

	// Add JavaScript for charts and formatting
	html.WriteString(`
        <script>
            // Helper function to format bytes to human-readable format
            function formatBytes(bytes, decimals = 2) {
                if (bytes === 0) return '0 Bytes';
                
                const k = 1024;
                const dm = decimals < 0 ? 0 : decimals;
                const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
                
                const i = Math.floor(Math.log(bytes) / Math.log(k));
                
                return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
            }
            
            // Toggle byte values visibility
            document.addEventListener('DOMContentLoaded', function() {
                const bytesValues = document.querySelectorAll('.bytes-value');
                let bytesVisible = false;
                
                // Add button to toggle byte values
                const toggleBtn = document.createElement('button');
                toggleBtn.innerText = 'Show Raw Bytes';
                toggleBtn.style = 'margin: 20px 0; padding: 8px 16px;';
                toggleBtn.onclick = function() {
                    bytesVisible = !bytesVisible;
                    bytesValues.forEach(el => {
                        el.style.display = bytesVisible ? 'inline' : 'none';
                    });
                    this.innerText = bytesVisible ? 'Hide Raw Bytes' : 'Show Raw Bytes';
                };
                
                // Initially hide byte values
                bytesValues.forEach(el => {
                    el.style.display = 'none';
                });
                
                // Add button to page
                document.querySelector('.container').insertBefore(toggleBtn, document.querySelector('h2'));
`)

	// Generate chart creation code for each test
	for i, r := range results {
		html.WriteString(fmt.Sprintf(`
                // Chart for %s
                new Chart(document.getElementById('chart%d'), {
                    type: 'bar',
                    data: {
`, r.Name, i))

		// Handle multiple type results
		if typeResults, ok := r.OtherStats["TypeResults"].(map[string]map[string]interface{}); ok {
			// Prepare labels and data
			var labels []string
			var optimizedData []uint64
			var unoptimizedData []uint64

			for typeName, typeData := range typeResults {
				labels = append(labels, typeName)

				if optMem, ok := typeData["OptimizedMemory"].(uint64); ok {
					optimizedData = append(optimizedData, optMem)
				} else {
					optimizedData = append(optimizedData, 0)
				}

				if unoptMem, ok := typeData["UnoptimizedMemory"].(uint64); ok {
					unoptimizedData = append(unoptimizedData, unoptMem)
				} else {
					unoptimizedData = append(unoptimizedData, 0)
				}
			}

			// Generate labels JSON
			html.WriteString("                        labels: [")
			for j, label := range labels {
				if j > 0 {
					html.WriteString(", ")
				}
				html.WriteString(fmt.Sprintf("'%s'", label))
			}
			html.WriteString("],\n")

			// Generate datasets
			html.WriteString(`                        datasets: [
                            {
                                label: 'Optimized',
                                backgroundColor: 'rgba(54, 162, 235, 0.8)',
                                data: [`)

			for j, data := range optimizedData {
				if j > 0 {
					html.WriteString(", ")
				}
				html.WriteString(fmt.Sprintf("%d", data))
			}

			html.WriteString(`]
                            },
                            {
                                label: 'Unoptimized',
                                backgroundColor: 'rgba(255, 99, 132, 0.8)',
                                data: [`)

			for j, data := range unoptimizedData {
				if j > 0 {
					html.WriteString(", ")
				}
				html.WriteString(fmt.Sprintf("%d", data))
			}

			html.WriteString(`]
                            }
                        ]
`)
		} else {
			// Standard results
			optMem, _ := r.OtherStats["OptimizedMemory"].(uint64)
			unoptMem, _ := r.OtherStats["UnoptimizedMemory"].(uint64)

			html.WriteString(fmt.Sprintf(`                        labels: ['%s'],
                        datasets: [
                            {
                                label: 'Optimized',
                                backgroundColor: 'rgba(54, 162, 235, 0.8)',
                                data: [%d]
                            },
                            {
                                label: 'Unoptimized',
                                backgroundColor: 'rgba(255, 99, 132, 0.8)',
                                data: [%d]
                            }
                        ]
`, r.Name, optMem, unoptMem))
		}

		// Chart options with custom tooltips for human-readable sizes
		html.WriteString(`                    },
                    options: {
                        responsive: true,
                        maintainAspectRatio: false,
                        scales: {
                            y: {
                                beginAtZero: true,
                                title: {
                                    display: true,
                                    text: 'Memory Usage'
                                },
                                ticks: {
                                    callback: function(value) {
                                        return formatBytes(value);
                                    }
                                }
                            }
                        },
                        plugins: {
                            tooltip: {
                                callbacks: {
                                    label: function(context) {
                                        var label = context.dataset.label || '';
                                        if (label) {
                                            label += ': ';
                                        }
                                        label += formatBytes(context.raw);
                                        return label;
                                    }
                                }
                            }
                        }
                    }
                });
`)
	}

	// Close JavaScript and HTML
	html.WriteString(`
            });
        </script>
    </div>
</body>
</html>
`)

	// Write to file
	if err := os.WriteFile("memory_test_results.html", html.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write HTML file: %w", err)
	}

	fmt.Println("HTML visualization saved to memory_test_results.html")
	return nil
}
