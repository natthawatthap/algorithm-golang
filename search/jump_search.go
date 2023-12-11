package search

import "math"

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))

	prev := 0
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1 // Target not found
		}
	}

	for i := prev; i < min(step, n); i++ {
		if arr[i] == target {
			return i // Target found, return its index
		}
	}

	return -1 // Target not found
}