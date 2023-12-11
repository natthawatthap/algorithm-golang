package search

func HashSearch(arr []int, target int) bool {
	hashTable := make(map[int]bool)

	for _, value := range arr {
		hashTable[value] = true
	}

	return hashTable[target]
}
