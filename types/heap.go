package types

// Heap is tree based datastructure
// https://www.wikiwand.com/en/Heap_(data_structure)
type Heap struct {
	nodes []int32
	isMin bool
}

func NewHeap(arr []int32) *Heap {
	h := &Heap{isMin: false}
	h.nodes = []int32{-1}

	for i := 0; i < len(arr); i++ {
		h.nodes = append(h.nodes, arr[i])
		h.siftUp()
	}
	return h
}

// NewMinHeap create a min heap
func NewMinHeap(array []int32) *Heap {
	h := &Heap{isMin: true}
	h.nodes = []int32{-1}

	for i := 0; i < len(array); i++ {
		h.nodes = append(h.nodes, array[i])
		h.siftUp()
	}

	return h
}

func (h *Heap) Add(val int32) {
	h.nodes = append(h.nodes, val)
	h.siftUp()
}

func (h *Heap) Elements() []int32 {
	return h.nodes[1:]
}

func (h *Heap) siftUp() {
	index := len(h.nodes) - 1
Loop:
	for index > 1 {
		parentIndex := index / 2

		if !h.isMin {
			if h.nodes[parentIndex] < h.nodes[index] {
				h.nodes[parentIndex], h.nodes[index] = h.nodes[index], h.nodes[parentIndex]
				index = parentIndex
			} else {
				break Loop
			}
		} else {
			if h.nodes[parentIndex] > h.nodes[index] {
				h.nodes[parentIndex], h.nodes[index] = h.nodes[index], h.nodes[parentIndex]
				index = parentIndex
			} else {
				break Loop
			}
		}

	}
}

// Pop pop root node
func (h *Heap) Pop() int32 {
	if len(h.nodes) <= 1 {
		return -1
	}

	top := h.nodes[1]
	h.nodes[1] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[0 : len(h.nodes)-1]
	h.siftDown()
	return top
}

func (h *Heap) siftDown() {
	start := 1

	for start < len(h.nodes) {
		var val int32
		var index int
		if h.isMin {
			val, index = h.getMinChild(start)
		} else {
			val, index = h.getMaxChild(start)
		}

		if val == -1 || index == -1 {
			break
		}

		if !h.isMin {
			if h.nodes[start] < val {
				h.nodes[start], h.nodes[index] = h.nodes[index], h.nodes[start]
				start = index
			} else {
				return
			}
		} else {
			if h.nodes[start] > val {
				h.nodes[start], h.nodes[index] = h.nodes[index], h.nodes[start]
				start = index
			} else {
				return
			}
		}
	}
}

func (h *Heap) getMaxChild(start int) (val int32, index int) {
	left, right := start*2, start*2+1
	if left >= len(h.nodes) {
		return -1, -1
	}

	leftChildVal := h.nodes[left]

	if right >= len(h.nodes) {
		val, index = leftChildVal, left
		return
	}

	rightChildVal := h.nodes[right]
	if leftChildVal > rightChildVal {
		val, index = leftChildVal, left
	} else {
		val, index = rightChildVal, right
	}

	return
}

func (h *Heap) getMinChild(start int) (val int32, index int) {
	left, right := start*2, start*2+1
	if left >= len(h.nodes) {
		return -1, -1
	}

	leftChildVal := h.nodes[left]

	if right >= len(h.nodes) {
		val, index = leftChildVal, left
		return
	}

	rightChildVal := h.nodes[right]
	if leftChildVal > rightChildVal {
		val, index = rightChildVal, right
	} else {
		val, index = leftChildVal, left
	}

	return
}
