package main

type MyHeap struct{
	capacity int
	size int
	items [] int
	reverse bool
}

func (mx *MyHeap ) GetLeftChildIndex(parentIndex int)int{
	return 2 * parentIndex + 1
}

func (mx *MyHeap ) GetRightChildIndex(parentIndex int)int{
	return 2 * parentIndex + 2
}

func (mx *MyHeap ) GetParentIndex(childIndex int)int {
	return (childIndex - 1) / 2
}

func (mx *MyHeap ) HasLeftChildIndex(index int)bool{
	return mx.GetLeftChildIndex(index) < mx.size
}

func (mx *MyHeap ) HasRightChildIndex(index int)bool{
	return mx.GetRightChildIndex(index) < mx.size
}

func (mx *MyHeap ) HasParent(index int)bool{
	return mx.GetParentIndex(index) >= 0
}

func (mx *MyHeap ) LeftChild(index int)int{
	return mx.items[mx.GetLeftChildIndex(index)]
}

func (mx *MyHeap ) RightChild(index int)int{
	return mx.items[mx.GetRightChildIndex(index)]
}

func (mx *MyHeap )GetParent(index int)int{
	return mx.items[mx.GetParentIndex(index)]
}

func (mx *MyHeap ) Swap(a,b int)*MyHeap {
	temp := mx.items[a]
	mx.items[a] = mx.items[b]
	mx.items[b] = temp
	return mx
}

func (mx *MyHeap ) Peek()int {
	if mx.size == 0 {
		panic("Heap is empty")
	}
	return mx.items[0]
}
func (mx *MyHeap ) Poll()int {
	if mx.size == 0 {
		panic("Heap is empty")
	}
	item := mx.items[0]
	mx.items[0] = mx.items[mx.size-1]
	mx.size--
	mx.HeapifyDown()
	return item
}

func (mx *MyHeap ) Add(item int)*MyHeap {
	mx.items = append(mx.items,item)
	mx.size++
	mx.HeapifyUp()
	return mx
}
func (mx *MyHeap ) HeapifyUp()*MyHeap {
	index := mx.size - 1
	if mx.reverse {
		for mx.HasParent(index) && mx.GetParent(index) < mx.items[index] {
			mx.Swap(mx.GetParentIndex(index), index)
			index = mx.GetParentIndex(index)
		}
	} else {
		for mx.HasParent(index) && mx.GetParent(index) > mx.items[index] {
			mx.Swap(mx.GetParentIndex(index), index)
			index = mx.GetParentIndex(index)
		}
	}
	return mx
}

func (mx *MyHeap ) HeapifyDown()*MyHeap {
	index := 0
	for mx.HasLeftChildIndex(index) {
		tempIndex := mx.GetLeftChildIndex(index)
		if mx.reverse {
			if mx.HasRightChildIndex(index) && mx.GetRightChildIndex(index) > mx.GetLeftChildIndex(index) {
				tempIndex = mx.GetRightChildIndex(index)
				mx.Swap(index, tempIndex)
			}
			if mx.items[index] > mx.items[tempIndex ] {
				break
			} else {
				mx.Swap(index, tempIndex)
			}
		} else {
			if mx.HasRightChildIndex(index) && mx.GetRightChildIndex(index) < mx.GetLeftChildIndex(index) {
				tempIndex = mx.GetRightChildIndex(index)
				mx.Swap(index, tempIndex)
			}
			if mx.items[index] < mx.items[tempIndex ] {
				break
			} else {
				mx.Swap(index, tempIndex)
			}
		}
	}
	return mx
}

