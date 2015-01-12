package machine

type PietMachine struct {
	arr  []int64
	size int
}

func NewMachine() *PietMachine {
	return &PietMachine{
		arr:  make([]int64, 1000, 1000),
		size: 0,
	}
}

func (p *PietMachine) GetStack() []int64 {
	if p.size == 0 {
		return []int64{}
	}
	x := make([]int64, p.size)
	for i := 0; i < p.size; i++ {
		x[i] = p.arr[p.size-1-i]
	}
	return x
}

func (p *PietMachine) Push(val int64) {
	p.arr[p.size] = val
	p.size++
}

func (p *PietMachine) Pop() {
	if p.size > 0 {
		p.size--
	}
}

func (p *PietMachine) pop() int64 {
	if p.size > 0 {
		p.size--
		return p.arr[p.size]
	}
	panic("unchecked pop")
}

func (p *PietMachine) Roll() {
	if p.size > 1 {
		count := int(p.pop())
		depth := int(p.pop())
		if count >= 0 {
			for i := 0; i < count; i++ {
				top := p.size - 1
				buryVal := p.arr[top]
				for j := top; j > top-depth; j-- {
					p.arr[j] = p.arr[j-1]
				}
				p.arr[top-depth+1] = buryVal
			}
		} else {
			for i := 0; i < -count; i++ {
				top := p.size - 1
				fetchVal := p.arr[top-depth+1]
				for j := top - (depth - 2); j <= top; j++ {
					p.arr[j-1] = p.arr[j]
				}
				p.arr[top] = fetchVal
			}
		}
	}
}

func (p *PietMachine) Not() {
	if p.size > 0 {
		val := p.pop()
		if val == 0 {
			p.Push(1)
		} else {
			p.Push(0)
		}
	}
}

func (p *PietMachine) Dup() {
	if p.size > 0 {
		val := p.pop()
		p.Push(val)
		p.Push(val)
	}
}

func (p *PietMachine) Binary(op string) {
	if p.size < 2 {
		return
	}
	a := p.pop()
	b := p.pop()
	switch op {
	case "+":
		p.Push(b + a)
	case "-":
		p.Push(b - a)
	case "*":
		p.Push(b * a)
	case "/":
		if a == 0 {
			p.Push(99999999)
			//this is what npiet does. The spec recommends ignoring:
			//p.stack.push(b)
			//p.stack.push(a)
			//TODO: make these a configuration option?
		} else {
			p.Push(b / a)
		}
	case "%":
		p.Push(b % a)
	case ">":
		if b > a {
			p.Push(1)
		} else {
			p.Push(0)
		}
	default:
		panic("unknown op: " + op)
	}

}
