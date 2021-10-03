package datastructures

type PriorityQueue struct {
	items []int
}

func (q *PriorityQueue) Size() int  {
	return len(q.items)
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{items: make([]int,0)}
}

func (q *PriorityQueue) reArrange()  {
	NewMaxHeap(q.items).BuildMaxHeap()
}

func (q *PriorityQueue) IsEmpty() bool{
	return q.Size() == 0
}

func (q *PriorityQueue) Peak() int {
	if q.IsEmpty(){
		panic("No item is in the queue.")
	}
	return q.items[0]
}

func (q *PriorityQueue) Dequeue() int {
	if q.IsEmpty(){
		panic("No item is in the queue.")
	}
	if q.IsEmpty(){
		panic("No item is in the queue.")
	}
	max := q.items[0]
	swapArrayElements(q.items, 0, q.Size()-1)
	q.items = q.items[:q.Size()-1]
	q.reArrange()
	return max
}

func (q *PriorityQueue) UpdateValue(index int, value int){
	if index < q.Size(){
		q.items[index] = value
		parentIndex := (index+1)/2-1
		for index >= 0 && q.items[parentIndex] < q.items[index]{
			swapArrayElements(q.items, index, parentIndex)
			index = parentIndex
		}
	}
}

func (q *PriorityQueue) Enqueue(item int) {
	if q.IsEmpty(){
		q.items = append(q.items, item)
	} else {
		q.items = append(q.items, item)
		q.reArrange()
	}
}

func (q *PriorityQueue) EnqueueItems(items []int) {
	for _, item := range items{
		q.Enqueue(item)
	}
}

func (q *PriorityQueue) Remove(index int) int{
	if q.IsEmpty(){
		panic("No item is in the queue.")
	}
	item := q.items[index]
	q.items[index] = q.items[q.Size()-1]
	q.items = q.items[:q.Size()-1]
	q.reArrange()
	return item
}


//Swap first and second index elements in items array
func swapArrayElements(items []int, first int, second int) {
	if first < len(items) && second < len(items) && first != second {
		temp := items[first]
		items[first] = items[second]
		items[second] = temp
	}
}