package search

func InterpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high && target >= arr[low] && target <= arr[high] {
		pos := low + ((high-low)*(target-arr[low]))/(arr[high]-arr[low])

		if arr[pos] == target {
			return pos // Target found, return its index
		}

		if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1 // Target not found
}