package types

import (
	"fmt"
)

// BinarySearchTree https://www.wikiwand.com/en/Binary_search_tree
/*
	Binary Search Tree, is a node-based binary tree data structure which has the following properties:
	The left subtree of a node contains only nodes with keys lesser than the node’s key.
	The right subtree of a node contains only nodes with keys greater than the node’s key.
	The left and right subtree each must also be a binary search tree.
	There must be no duplicate nodes.
*/
type BinarySearchTree struct {
	Root *BinTreeNode
}

func (bst *BinarySearchTree) Search(key int32) *BinTreeNode {
	return search(bst.Root, key)
}

func search(node *BinTreeNode, key int32) *BinTreeNode {
	if node == nil {
		return nil
	}

	if node.Data == key {
		return node
	}

	if key < node.Data {
		return search(node.Left, key)
	} else {
		return search(node.Right, key)
	}
}

func (bst *BinarySearchTree) InOrder() []int32 {
	var data []int32
	inorderTraversal(bst.Root, &data)
	return data
}

func inorderTraversal(node *BinTreeNode, data *[]int32) {
	if node != nil {
		inorderTraversal(node.Left, data)
		fmt.Println(node.Data)
		*data = append(*data, node.Data)
		inorderTraversal(node.Right, data)
	}
}

func (bst *BinarySearchTree) Insert(key int32) {
	bst.Root = insert(bst.Root, key)
}

func insert(node *BinTreeNode, key int32) *BinTreeNode {
	if node == nil {
		return &BinTreeNode{Data: key}
	}

	if key < node.Data {
		node.Left = insert(node.Left, key)
	} else {
		node.Right = insert(node.Right, key)
	}

	return node
}

func (bst *BinarySearchTree) Delete(key int32) {

}
