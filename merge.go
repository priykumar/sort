package main

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	l, r := 0, len(a)-1
	m := l + (r-l)/2
	b := mergeSort(a[:m+1])
	c := mergeSort(a[m+1:])
	return mergeArrays(b, c)
}

func mergeArrays(b, c []int) []int {
	m, n := len(b), len(c)
	i, j := 0, 0
	res := []int{}

	for i < m && j < n {
		if b[i] <= c[j] {
			res = append(res, b[i])
			i++
		} else {
			res = append(res, c[j])
			j++
		}
	}

	if i < m {
		res = append(res, b[i:]...)
	}
	if j < n {
		res = append(res, c[j:]...)
	}

	return res
}

/*
Time Complexities
Best Case: O(n log n)
Average Case: O(n log n)
Worst Case: O(n log n)

Always splits the array exactly in half: m := l + (r-l)/2. Creates a balanced binary tree of recursive calls. Tree depth is always log n levels, 
hence all cases are the same

Space Complexity
O(n) 
*/


