package bst

import "fmt"

func PrintPostorder(node *node) {
	if node != nil {
		PrintPostorder(node.left)
		PrintPostorder(node.right)
		fmt.Println(node.value)
	}
}
