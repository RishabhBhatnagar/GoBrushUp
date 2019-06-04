// 1) Array Implementation of Stack.

package main


type Stack struct{
    top int
    stack_size int
    arr []int
}


func (st *Stack) pop() int{
    if st.top < 1 {
        panic("StackUnderflow error: Can't pop from an empty stack")
    } else{
        st_top_value := st.arr[st.top-1]   // since array index started from 0
        st.top -= 1;
        return st_top_value
    }
    return -1  // error condition.
}


func (st *Stack) push(element int){
    if st.top >= st.stack_size {
        panic("StackUnderFlow error: Cannot push on a full stack")
    } else {
        st.arr = append(st.arr, element)
        st.top += 1
    }
}


func main(){
    st := Stack{stack_size: 10}
    st.push(23)
    st.pop()               // returns 23.
}
