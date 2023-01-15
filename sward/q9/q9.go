package q9

type CQueue struct {
	num1 []int
	num2 []int
}

func Constructor() CQueue {
	return CQueue{num1: []int{}, num2: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.num1 = append(this.num1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.num2) == 0 {
		if len(this.num1) == 0 {
			return -1
		}
		this.in2out()
	}
	ret := this.num2[len(this.num2)-1]
	this.num2 = this.num2[:len(this.num2)-1]
	return ret
}

func (this *CQueue) in2out() {
	for len(this.num1) > 0 {
		this.num2 = append(this.num2, this.num1[len(this.num1)-1])
		this.num1 = this.num1[:len(this.num1)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
