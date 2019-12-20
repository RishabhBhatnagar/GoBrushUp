// 1) Array Implementation of Stack.

package main

type stack struct {
	top       int
	stackSize int
	arr       []int
}

func (st *stack) pop() int {
	if st.top < 1 {
		panic("StackUnderflow error: Can't pop from an empty stack")
	} else {
		stTopValue := st.arr[st.top-1] // since array index started from 0
		st.top--
		return stTopValue
	}
}

func (st *stack) push(element int) {
	if st.top >= st.stackSize {
		panic("StackUnderFlow error: Cannot push on a full stack")
	} else {
		st.arr = append(st.arr, element)
		st.top++
	}
}

func main() {
	st := stack{stackSize: 10}
	st.push(23)
	st.pop() // returns 23.
}
