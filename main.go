package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

// binary search tree
type bst struct {
	root   *node
	length int
}

func main() {

	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}

	b := bst{
		root:   n,
		length: 3,
	}
	//show bst
	fmt.Println(b)

	//adding values
	b.add(5)
	b.add(4)
	b.add(6)

	//show bst
	fmt.Println(b)

	//searching values
	fmt.Println(b.search(6))
	fmt.Println(b.search(3))

	//delete values
	b.remove(1)
	fmt.Println(b)
	b.remove(5)
	fmt.Println(b)
}

// ---STRING---
func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	stringBuilder := strings.Builder{}
	b.inOrderTraversal(&stringBuilder)
	return stringBuilder.String()
}

// ---TRAVERSAL---
func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	//Inorder-Tree-Walk
	//Nodes from the left subtree get visited first, followed by the root node and right subtree.
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.right)
}

// ---INSERTION---
func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.length++
}

func (b *bst) addByNode(root *node, value int) *node {
	//New nodes are inserted as leaf nodes in the BST

	//check if it is the bottom of the tree and create node
	if root == nil {
		return &node{value: value}
	}
	//check the adding value with the node value
	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	//return *node for another iteration
	return root
}

// ---SEARCH---
func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	//check if bottom of the tree was reached
	if root == nil {
		return nil, false
	}
	//check input value with node value
	if root.value == value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

// ---REMOVE---
func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
	b.length--
}

func (b *bst) removeByNode(root *node, value int) *node {
	//check if bottom of the tree was reached
	if root == nil {
		return root
	}

	//1° - search for the value to be deleted
	if root.value < value {
		root.right = b.removeByNode(root.right, value)
	} else if root.value > value {
		root.left = b.removeByNode(root.left, value)
	} else {
		//2° - once found, if the left or right node is empty, return the opposite node
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		//3° - if both children nodes exists, you need to find the next highest value in the BST
		//find the max value of the left subtree
		maxValueNode := root.left
		for maxValueNode.right != nil {
			maxValueNode = maxValueNode.right
		}
		//replace root value with the max value of the left subtree
		root.value = maxValueNode.value
		//delete the node with the input value
		root.left = b.removeByNode(root.left, maxValueNode.value)
	}
	return root
}
