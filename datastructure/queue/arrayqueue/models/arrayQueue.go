package models

type ArrayQueue struct {
	Arr   []int
	Front int
	Rear  int
}

func NewArrayQueue(arrLen, front, rear int) *ArrayQueue {
	return &ArrayQueue{
		Arr:   make([]int, arrLen),
		Front: front,
		Rear:  rear,
	}
}
