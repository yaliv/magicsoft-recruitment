// Subprogram 1: array visualizer.
package main

import "fmt"

// Create a vertical barcharts visualization.
func visualize(A []int) {
	fmt.Println()

	// Get the highest number.
	var max int

	for _, v := range A {
		if v > max {
			max = v
		}
	}

	// Print the vertical barcharts.
	for row := max; row > 0; row-- {
		for _, col := range A {
			if col >= row {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// Print the values.
	for _, v := range A {
		fmt.Print(v)
	}

	fmt.Println()
}
