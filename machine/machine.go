package machine

import ()

type PietMachine struct {
	stack stack
}

func NewMachine() *PietMachine {
	return &PietMachine{
		stack: stack{
			arr:  make([]int64, 1000, 1000),
			size: 0,
		},
	}
}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PietMachine) GetStack() []int64 {
	x := make([]int64, p.stack.size)
	for i := 0; i < p.stack.size; i++ {
		x[i] = p.stack.arr[p.stack.size-1-i]
	}
	return x
}

func (p *PietMachine) Push(val int64) {
	p.stack.push(val)
}

func (p *PietMachine) Binary(op string) {
	if p.stack.size < 2 {
		return
	}
	a := p.stack.pop()
	b := p.stack.pop()
	switch op {
	case "+":
		p.stack.push(b + a)
	case "-":
		p.stack.push(b - a)
	default:
		panic("unknown op: " + op)
	}

}
