package main

func partition(a []int) int {
	p := 0
	i, j := 1, len(a)-1
	for i <= j {
		for i <= j && a[i] < a[p] {
			i++
		}
		for i <= j && a[j] >= a[p] {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	a[p], a[j] = a[j], a[p]

	return j
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	p := partition(a)
	quickSort(a[:p])
	quickSort(a[p+1:])
	return a
}

/*
Best Case: O(n log n)
Occurs when the pivot divides the array into two nearly equal halves at each step
Each level of recursion processes all n elements: O(n)
Number of levels in the recursion tree: O(log n)
Total: O(n) × O(log n) = O(n log n)

Average Case: O(n log n)

Worst Case: O(n²)
Occurs when the pivot is always the smallest or largest element
This creates maximally unbalanced partitions (one side has n-1 elements, other has 0)
Results in n levels of recursion, each processing up to n elements
Total: O(n²)

When Does Worst Case Happen?
Already sorted array (ascending or descending) with first/last element as pivot
All elements are the same
Reverse sorted array with poor pivot selection

Space Complexity
Best/Average Case: O(log n) - due to recursion stack depth
Worst Case: O(n) - when recursion depth becomes n

Improving Performance
Random pivot selection: Reduces probability of worst case
Median-of-three: Choose pivot as median of first, middle, and last elements
*/
