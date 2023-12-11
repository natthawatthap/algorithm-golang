package sort 

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Assume the current index is the minimum.
		minIndex := i

		// Find the index of the minimum element in the remaining unsorted part of the array.
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the element at the current index.
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}