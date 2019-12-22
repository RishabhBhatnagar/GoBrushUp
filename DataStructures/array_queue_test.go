package data_structures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	// Testing the initial parameters of the queue.
	testQueue := Queue(10)
	if testQueue.rear != -1 || testQueue.front != -1 {
		t.Errorf("Queue initialisation not proper.")
	}
}

func TestQueue_Operations(t *testing.T) {
	testQueue := Queue(10)
	_, errCode := testQueue.Dequeue()
	if errCode == nil {
		t.Errorf("Dequeue from an empty queue didn't raise any error.")
	}

	for _, ele := range []int{1, 2, 3, 4, 5} {
		_ = testQueue.Enqueue(ele)
	}

	if testQueue.Size() != 5 {
		t.Errorf("Queue Size not correct")
	}

	if ele, _ := testQueue.Dequeue(); ele != 1 {
		t.Errorf("Dequeue not proper")
	}
	// inserting 5 more element such that queue is full to brim.
	for _, ele := range []int{1, 2, 3, 4, 5} {
		_ = testQueue.Enqueue(ele)
	}

	// emptying the entire queue such that, front points to rear at the end of the array.
	for testQueue.Size() != 0 {
		testQueue.Dequeue()
	}

	// now, front is ahead of rear by one distance at the last of the array.
	// queue is empty but element cannot be assigned at q.rear+1. Realignment must work.
	testQueue.Enqueue(1)
	// after realignment and insertion of a single element, front and rear both should point to 0.
	if testQueue.front != 0 || testQueue.rear != 0 {
		t.Errorf("Realign function didn't work well.")
	}
}
