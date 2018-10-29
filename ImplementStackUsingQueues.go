package leetcode

type MyStack struct {
	data []int
}


/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.data = append(this.data, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	l := len(this.data)
	if l == 0 {
		return 0
	}
	x := this.data[l-1]
	if l < 2 {
		this.data = []int{}
	} else {
		this.data = this.data[:l-1]
	}

	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	l := len(this.data)

	if l < 1 {
		return 0
	}
	return this.data[l-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
