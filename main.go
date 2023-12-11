package main

import (
	"algorithm-golang/search"
	"algorithm-golang/sort" // Update with the correct module path
	"fmt"
)

func main() {
	arr := []int{64, 25, 12, 22, 11}

	fmt.Println("Unsorted array:", arr)

	//sort.SelectionSort(arr)
	//sort.QuickSort(arr)
	//sort.BubbleSort(arr)
	//sort.InsertionSort(arr)
	sort.TreeSort(arr)
	fmt.Println("Sorted array:", arr)

	target := 12
	index := search.BinarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found in the array.\n", target)
	} else {
		fmt.Printf("Element %d not found in the array.\n", target)
	}

}
