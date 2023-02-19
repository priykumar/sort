package main

import "fmt"

func getInput() [][]int {
	return [][]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		{1, 2, 3, 4, 5, 5, 6, 6, 7},
		{1, 2, 1, 1, 3, 2, 1, 3},
		{1, 6, 2, 3, 8, 9, 0},
		{5, 2, 3, 4, 5, 1},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1},
		{2, 1, 3},
		{2, 1},
		{1},
		{},
	}
}
func main() {

	input := getInput()
	fmt.Println("--------------------------------------------------")
	fmt.Println("                   Bubble Sort                    ")
	fmt.Println("TC: O(n*n)                                SC: O(1)")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - ")
	for _, e := range input {
		fmt.Println(bubbleSort(e))
	}
	fmt.Println("..................................................")

	input = getInput()
	fmt.Println("--------------------------------------------------")
	fmt.Println("                  Selection Sort                  ")
	fmt.Println("TC: O(n*n)                                SC: O(1)")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - ")
	for _, e := range input {
		fmt.Println(selectionSort(e))
	}
	fmt.Println("..................................................")

	input = getInput()
	fmt.Println("--------------------------------------------------")
	fmt.Println("                    Merge Sort                    ")
	fmt.Println("TC: O(n*logn)                             SC: O(n)")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - ")
	for _, e := range input {
		fmt.Println(mergeSort(e))
	}
	fmt.Println("..................................................")

	input = getInput()
	fmt.Println("--------------------------------------------------")
	fmt.Println("                  Insertion Sort                  ")
	fmt.Println("TC: O(n*n)                                SC: O(1)")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - ")
	for _, e := range input {
		fmt.Println(insertionSort(e))
	}
	fmt.Println("..................................................")

	input = getInput()
	fmt.Println("--------------------------------------------------")
	fmt.Println("                   Quick Sort                     ")
	fmt.Println("TC: O(n*logn)                             SC: O(1)")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - ")
	for _, e := range input {
		fmt.Println(quickSort(e))
	}
	fmt.Println("..................................................")
}
