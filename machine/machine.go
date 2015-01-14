package machine

import ()

type PietMachine struct {
	top  *stackNode
	size int
}

type stackNode struct {
	value int64
	next  *stackNode
}

func NewMachine() *PietMachine {
	return &PietMachine{
		top:  nil,
		size: 0,
	}
}

func (p *PietMachine) GetStack() []int64 {
	if p.size == 0 {
		return []int64{}
	}
	x := make([]int64, p.size)
	n := p.top
	i := 0
	for n != nil {
		x[i] = n.value
		i++
		n = n.next
	}
	return x
}

func (p *PietMachine) Push(val int64) {
	p.size++
	p.top = &stackNode{value: val, next: p.top}
}

func (p *PietMachine) Pop() {
	if p.size > 0 {
		p.size--
		p.top = p.top.next
	}
}

func (p *PietMachine) pop() int64 {
	if p.size > 0 {
		p.size--
		top := p.top
		p.top = p.top.next
		return top.value
	}
	panic("unchecked pop")
}

func (p *PietMachine) Roll() {
	if p.size > 1 {
		count := int(p.pop())
		depth := int(p.pop())
		if count >= 0 {
			for i := 0; i < count; i++ {
				bury := p.top
				p.top = bury.next
				current := bury
				currentIdx := 1
				for currentIdx < depth {
					current = current.next
					currentIdx++
				}
				bury.next = current.next
				current.next = bury
			}
		} else {
			for i := 0; i < -count; i++ {
				current := p.top
				var prev *stackNode = nil
				currentIdx := 1
				for currentIdx < depth {
					prev = current
					current = current.next
					currentIdx++
				}
				if prev != nil {
					prev.next = current.next
				}
				current.next = p.top
				p.top = current
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
