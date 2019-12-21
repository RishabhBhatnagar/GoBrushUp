package data_structures

import "testing"

func TestStack(t *testing.T) {
	testStack := stack{stackSize:10};
	for _, ele := range []int{1, 4, 2, 3, 7} {
		testStack.push(ele)
	}
	tables := []struct{
		result int
		expected int
		fieldName string
	}{
		{testStack.top, 4, "stackTop"},
		{testStack.stackSize, 10, "stackSize"},
		{testStack.peek(), 7, "Peek"},
		{testStack.pop(), 7, "popFunction"},
		{testStack.peek(), 3, "Peek"},
	}
	for _, ele := range tables {
		if ele.expected != ele.result {
			t.Errorf("%s not consistent. Expected: %d, Found: %d", ele.fieldName, ele.expected, ele.result)
		}
	}
}

