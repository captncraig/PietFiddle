package main

type stack struct {
	arr  []int
	size int
}

func (self *stack) push(i int) {
	self.arr[self.size] = i
	self.size++
}

func (self *stack) pop() int {
	self.size--
	return self.arr[self.size]
}
