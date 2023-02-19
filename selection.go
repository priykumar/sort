package main

func selectionSort(a []int) []int {
	n := len(a)

	for i := 0; i < n; i++ {
		min, minIndex := a[i], i
		for j := i; j < n; j++ {
			if a[j] < min {
				min = a[j]
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}

	return a
}
