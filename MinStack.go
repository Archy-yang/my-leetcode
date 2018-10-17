package leetcode

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 || this.min > x {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	sLen := len(this.stack)
	if sLen > 1 {
		this.stack = this.stack[:len(this.stack)-1]

		this.min = this.stack[0]
		for _, i := range this.stack {
			if this.min > i {
				this.min = i
			}
		}
	} else {
		tmp := &MinStack{}
		this.stack = tmp.stack
		this.min = tmp.min
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
