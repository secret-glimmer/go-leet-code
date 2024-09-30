package main

type CustomStack struct {
	stack []int
	top   int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{stack: make([]int, maxSize), top: -1}
}

func (this *CustomStack) Push(x int) {
	if this.top < len(this.stack)-1 {
		this.top++
		this.stack[this.top] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.top != -1 {
		val := this.stack[this.top]
		this.top--
		return val
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(k, this.top+1); i++ {
		this.stack[i] += val
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
