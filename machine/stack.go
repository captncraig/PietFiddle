package machine

//TODO: Auto expand
type stack struct {
	arr  []int64
	size int
}

func (self *stack) push(i int64) {
	self.arr[self.size] = i
	self.size++
}

func (self *stack) pop() int64 {
	self.size--
	return self.arr[self.size]
}
