package binarytree

import (
	"fmt"

	list "github.com/emirpasic/gods/lists/arraylist"
	stack "github.com/emirpasic/gods/stacks/arraystack"

	"github.com/zhiruchen/500Algorithms/types"
)

// InorderTraversal left -> root -> right
func InorderTraversal(root *types.BinTreeNode) {
	s := stack.New()
	curr := root

	for !s.Empty() || curr != nil {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {
			v, _ := s.Pop()
			curr = v.(*types.BinTreeNode)
			fmt.Println(curr.Data)
			curr = curr.Right
		}
	}
}

// PreorderTraversal root -> left -> right
func PreorderTraversal(root *types.BinTreeNode) {
	s := stack.New()
	s.Push(root)

	for !s.Empty() {
		v, _ := s.Pop()
		curr := v.(*types.BinTreeNode)
		fmt.Println(curr.Data)

		if curr.Right != nil {
			s.Push(curr.Right)
		}

		if curr.Left != nil {
			s.Push(curr.Left)
		}
	}
}

// PostorderTraversal left -> right -> root
func PostorderTraversal(root *types.BinTreeNode) {
	s := stack.New()
	s.Push(root)
	out := stack.New()

	for !s.Empty() {
		v, _ := s.Pop()
		curr := v.(*types.BinTreeNode)
		out.Push(curr.Data)

		if curr.Left != nil {
			s.Push(curr.Left)
		}

		if curr.Right != nil {
			s.Push(curr.Right)
		}
	}

	for !out.Empty() {
		v, _ := out.Pop()
		curr := v.(*types.BinTreeNode)
		fmt.Println(curr.Data)
	}
}

// LevelOrderTraversal level order
func LevelOrderTraversal(root *types.BinTreeNode) {
	if root == nil {
		return
	}

	lst := list.New()
	lst.Add(root)

	for !lst.Empty() {
		v, _ := lst.Get(0)
		lst.Remove(0)
		curr := v.(*types.BinTreeNode)
		fmt.Println(curr.Data)

		if curr.Left != nil {
			lst.Add(curr.Left)
		}

		if curr.Right != nil {
			lst.Add(curr.Right)
		}
	}
}
