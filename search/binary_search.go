package search 

func BinarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1

    for low <= high {
        mid := (low + high) / 2

        if arr[mid] == target {
            return mid // Target found, return its index
        } else if arr[mid] < target {
            low = mid + 1 // Discard the left half
        } else {
            high = mid - 1 // Discard the right half
        }
    }

    return -1 // Target not found
}