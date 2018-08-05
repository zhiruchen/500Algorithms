package types

// Heap is tree based datastructure
// https://www.wikiwand.com/en/Heap_(data_structure)
type Heap struct {
	nodes []int32
}

func NewHeap(arr []int32) *Heap {
	h := &Heap{nodes: make([]int32, len(arr)+1)}
	for i := 1; i <= len(arr); i++ {
		h.nodes[i] = arr[i-1]
	}
	return h
}

func (h *Heap) Add(val int32) {
	h.nodes = append(h.nodes, val)
	h.siftUp()
}

func (h *Heap) siftUp() {
	index := len(h.nodes) - 1
	for index > 1 {
		parentIndex := index / 2
		if h.nodes[parentIndex] < h.nodes[index] {
			h.nodes[parentIndex], h.nodes[index] = h.nodes[index], h.nodes[parentIndex]
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	start := 1

	var maxVal int32
	var maxIndex int
	for start < len(h.nodes) {
		left, right := start*2, start*2+1
		if left >= len(h.nodes) {
			break
		}

		leftChildVal := h.nodes[left]

		if right >= len(h.nodes) {
			maxVal, maxIndex = leftChildVal, left
		} else {
			rightChildVal := h.nodes[right]
			if leftChildVal > rightChildVal {
				maxVal, maxIndex = leftChildVal, left
			} else {
				maxVal, maxIndex = rightChildVal, right
			}
		}

		if h.nodes[start] < maxVal {
			h.nodes[start], h.nodes[maxIndex] = h.nodes[maxIndex], h.nodes[start]
			start = maxIndex
		}
	}
}
