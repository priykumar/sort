package main

func insertionSort(a []int) []int {
	n := len(a)

	for i := 1; i < n; i++ {
		ele := a[i]
		j := i - 1

		// Shift elements greater than ele to the right
		for j >= 0 && a[j] > ele {
			a[j+1] = a[j]
			j--
		}

		// Place ele at the correct position
		a[j+1] = ele
	}

	return a
}

/*
Time Complexities
Best Case: O(n)  -->  Occurs when the array is already sorted
Average Case: O(n²)
Worst Case: O(n²)  -->  Occurs when the array is reverse sorted

Space Complexity
O(1)
*/
