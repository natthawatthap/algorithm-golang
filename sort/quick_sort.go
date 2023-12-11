package sort 

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Choose a pivot element.
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Partition the array around the pivot.
	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			// Swap elements at left and right.
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-arrays.
	QuickSort(arr[:right+1])
	QuickSort(arr[left:])
}