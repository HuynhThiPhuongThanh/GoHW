// Last updated: 6/28/2025, 3:33:50 AM
type MyQueue struct {
	stackIn  []int 
	stackOut []int 
}


func Constructor() MyQueue {
	return MyQueue{
		stackIn:  []int{},
		stackOut: []int{},
	}
}


func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}


func (q *MyQueue) Pop() int {
	q.transfer() 
	if len(q.stackOut) == 0 {
		return -1 
	}
	x := q.stackOut[len(q.stackOut)-1] 
	q.stackOut = q.stackOut[:len(q.stackOut)-1] 
	return x
}


func (q *MyQueue) Peek() int {
	q.transfer() 
	if len(q.stackOut) == 0 {
		return -1 
	}
	return q.stackOut[len(q.stackOut)-1]
}


func (q *MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

func (q *MyQueue) transfer() {
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			top := q.stackIn[len(q.stackIn)-1]
			q.stackIn = q.stackIn[:len(q.stackIn)-1] 
			q.stackOut = append(q.stackOut, top)
		}
	}
}
