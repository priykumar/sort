package main

func bubbleSort(a []int) []int {
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

/*
Time Complexities
Best Case: O(n²)
Average Case: O(n²)
Worst Case: O(n²)

Optimised Best Case O(n)  -->  When we use a swapped flag and if no swap happen in first itertaion then just break

Space Complexity
O(1)
*/
