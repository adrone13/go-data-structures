package bst

import (
	"fmt"
	"log"
	"strings"

	"go-data-structures/stack"
)

// Trees to implement:
// - BST
// - AVL
// - Red-Black Tree

var pad = "--"

type node struct {
	val   int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
}

//			        10
//				  /   \
//				5      12
//			  /  \    /  \
//			 1    6 11    13
//         /
//        0

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(value int) {
	if bst.root == nil {
		bst.root = &node{val: value}
		return
	}

	current := bst.root
	for {
		if current.val == value {
			log.Printf("Node already %d exists", value)
			break
		}
		if current.val > value && current.left == nil {
			log.Printf("Found a place for %d left from %d", value, current.val)
			current.left = &node{val: value}
			break
		}
		if current.val < value && current.right == nil {
			log.Printf("Found a place for %d right from %d", value, current.val)
			current.right = &node{val: value}
			break
		}

		if current.val > value && current.left != nil {
			log.Printf("Looking left from %d", current.left.val)
			current = current.left
			continue
		}

		if current.val < value && current.right != nil {
			log.Printf("Looking right from %d", current.right.val)
			current = current.right
		}
	}
}

func (bst *BinarySearchTree) Delete(value int) {
	current := bst.root
	for {
		if current.val == value {

		}
	}
}

func (bst *BinarySearchTree) Search(value int) bool {
	log.Printf("Searching value %d", value)

	if bst.root == nil {
		return false
	}

	current := bst.root
	for {
		log.Printf("Current node %d", current.val)

		if current.val == value {
			return true
		}
		if value < current.val && current.left == nil {
			return false
		}
		if value > current.val && current.right == nil {
			return false
		}

		if value < current.val {
			current = current.left
		} else {
			current = current.right
		}
	}
}

func (bst *BinarySearchTree) Dfs() {
	if bst.root == nil {
		fmt.Println("Empty tree")
		return
	}

	type stackVal struct {
		n     *node
		pos   string
		level int
	}

	stack := stack.New[stackVal]()
	stack.Push(stackVal{bst.root, "Root", 0})
	for stack.Len() > 0 {
		current := stack.Pop()
		fmt.Printf("%s%s: %d\n", strings.Repeat(pad, current.level), current.pos, current.n.val)

		if current.n.right != nil {
			stack.Push(stackVal{current.n.right, "Right", current.level + 1})
		}
		if current.n.left != nil {
			stack.Push(stackVal{current.n.left, "Left", current.level + 1})
		}
	}
}

func (bst *BinarySearchTree) DfsRecursive() {
	fmt.Println("Root:", bst.root.val)
	if bst.root.left != nil {
		fmt.Printf("%sLeft: %d\n", pad, bst.root.left.val)
		bst.dfsRecursive(bst.root.left, pad+pad)
	}
	if bst.root.right != nil {
		fmt.Printf("%sRight: %d\n", pad, bst.root.right.val)
		bst.dfsRecursive(bst.root.right, pad+pad)
	}
}

func (bst *BinarySearchTree) dfsRecursive(n *node, padding string) {
	if n == nil {
		return
	}
	if n.left != nil {
		fmt.Printf("%sLeft: %d\n", padding, n.left.val)
		bst.dfsRecursive(n.left, padding+pad)
	}
	if n.right != nil {
		fmt.Printf("%sRight: %d\n", padding, n.right.val)
		bst.dfsRecursive(n.right, padding+pad)
	}
}

func (bst *BinarySearchTree) Bfs() {

}
