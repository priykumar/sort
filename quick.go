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
