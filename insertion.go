package main

func insertionSort(a []int) []int {
	n := len(a)

	for i := 1; i < n; i++ {
		ele := a[i]
		j := i - 1
		for j >= 0 && a[j] > ele {
			a[j], a[j+1] = a[j+1], a[j]
		}
		a[j+1] = ele
	}

	return a
}
