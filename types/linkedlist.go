package types

type LinkedListNode struct {
	Data int32
	Next *LinkedListNode
}

type LinkedList struct {
	Head  *LinkedListNode
	Tail  *LinkedListNode
	Count int32
}
