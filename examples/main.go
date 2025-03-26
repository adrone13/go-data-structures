package main

import (
	"fmt"
	"go-data-structures/bst"
)

func main() {
	tree := bst.New()

	//values := []int{1, 2, 3, 4, 5, 6, 7}
	//               0  1   2  3  4   5   6
	values := []int{10, 5, 12, 1, 6, 11, 13, 0}
	//values := []int{10, 5, 12, 1, 6, 0}
	// index for left node = parent index * 2 + 1
	// index for right node = parent index * 2 + 2

	for _, v := range values {
		tree.Insert(v)
	}
	tree.DfsRecursive()
	fmt.Println()
	tree.Dfs()

	//for _, v := range values {
	//	if tree.Search(v) {
	//		log.Printf("Value %d found", v)
	//		log.Println()
	//	} else {
	//		log.Printf("Value %d not found", v)
	//		log.Println()
	//	}
	//}

	//tree.Insert(10)
	//tree.Insert(10)
	//tree.Insert(5)
	//tree.Insert(4)
	//tree.Insert(11)

}
