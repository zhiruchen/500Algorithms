package types

type BinTreeNode struct {
	Data  int32
	Left  *BinTreeNode
	Right *BinTreeNode
}

type BinaryTree struct {
	Root *BinTreeNode
}

func NewBinaryTree(root *BinTreeNode) *BinaryTree {
	return &BinaryTree{Root: root}
}
