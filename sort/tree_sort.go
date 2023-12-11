package sort

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert inserts a key into the binary search tree.
func Insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil}
	}

	if key < root.Key {
		root.Left = Insert(root.Left, key)
	} else if key > root.Key {
		root.Right = Insert(root.Right, key)
	}

	return root
}

// InOrderTraversal performs an in-order traversal of the binary search tree.
func InOrderTraversal(root *Node, result *[]int) {
	if root != nil {
		InOrderTraversal(root.Left, result)
		*result = append(*result, root.Key)
		InOrderTraversal(root.Right, result)
	}
}

// TreeSort performs tree sort on an array of integers.
func TreeSort(arr []int) []int {
	var root *Node

	for _, value := range arr {
		root = Insert(root, value)
	}

	var result []int
	InOrderTraversal(root, &result)

	return result
}