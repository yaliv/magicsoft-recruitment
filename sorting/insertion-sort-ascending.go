// Subprogram 2: ascending insertion sort.
package main

// Perform the ascending insertion sort.
func ascSort(A []int) {
	// Visualize original array.
	visualize(A)

	// Insertion sort.
	for i := 1; i < len(A); i++ {
		// If the left element is bigger than the right element, switch position.
		for j := i; j > 0 && A[j-1] > A[j]; j-- {
			A[j], A[j-1] = A[j-1], A[j]

			// Visualize each step.
			visualize(A)
		}
	}
}
