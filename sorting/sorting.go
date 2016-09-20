// Sorting main program.
package main

import (
	"fmt"
	"time"
)

func main() {
	A := []int{1, 4, 5, 6, 8, 2}

	ascA := make([]int, len(A))
	copy(ascA, A)
	descA := make([]int, len(A))
	copy(descA, A)

	// Original array.
	visualize(A)
	fmt.Println("\nThis is original array.")

	time.Sleep(3 * time.Second)

	// Ascending insertion sort.
	ascSort(ascA)
	fmt.Println("\nThis is sorted array in ascending order.")

	time.Sleep(3 * time.Second)

	// Descending insertion sort.
	descSort(descA)
	fmt.Println("\nThis is sorted array in descending order.")
}
