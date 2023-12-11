package search 

func LinearSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i // Target found, return its index
        }
    }
    return -1 // Target not found
}