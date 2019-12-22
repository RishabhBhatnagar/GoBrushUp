package data_structures

import (
	"errors"
	)

type queue struct{
	front int
	rear  int
	arr   []int
}

func Queue(queue_size int) *queue{
	// default constructor like function to init queue and satellite fields.
	var newQueue queue
	newQueue.rear = -1
	newQueue.front = -1
	newQueue.arr = make([]int, queue_size, queue_size)
	return &newQueue
}

func (q *queue) Enqueue(ele int) error{
	// if ele is the first element,
	if q.IsEmpty() {
		q.front++
	} else {
		if q.Size() == len(q.arr){
			// queue is full
			return errors.New("Queue Overflow")
		} else {
			// queue is not full but rear points to end.
			q.realign()
		}
	}
	q.rear++
	q.arr[q.rear] = ele
	return nil
}

func (q *queue) Dequeue() (ele int, err error){
	if q.Size() != 0 {
		ele = q.arr[q.front]
		q.arr[q.front] = 0
		q.front++
	} else {
		err = errors.New("Cannot remove from an empty queue.")
	}
	return
}

func (q *queue) IsEmpty() bool {
	return q.front == -1
}

func (q *queue) Size() int {
	if q.IsEmpty() {
		return 0   // no element is inserted
	}
	return q.rear - q.front + 1
}

func (q *queue) realign() {
	// realign front and internal arr start
	displacement := q.front
	for run, n := 0, q.Size(); run < n; run++ {
		q.arr[run] = q.arr[run + q.front]
	}
	q.front = 0
	q.rear -= displacement
}
