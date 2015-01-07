package machine

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

func (p *PietMachine) GetStack() []int64 {
	return p.stack.arr[0:p.stack.size]
}
