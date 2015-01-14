package micropiet

import (
	"github.com/captncraig/pietfiddle/machine"
)

type Program struct {
	startCommand *Command
	labels       map[string]*Command

	current       *Command
	pendingLabels []string
}

func NewProgram() *Program {
	return &Program{
		startCommand:  nil,
		labels:        map[string]*Command{},
		current:       nil,
		pendingLabels: []string{},
	}
}

func (p *Program) Run(vm *machine.PietMachine) {
	cur := p.startCommand
	for cur != nil {
		jmp := cur.action(vm)
		if jmp == "" {
			cur = cur.next
		} else {
			cur = p.labels[jmp]
		}
	}
}

func (p *Program) Mark(lbl string) {
	lbl = lbl[0 : len(lbl)-1]
	p.pendingLabels = append(p.pendingLabels, lbl)
}

func (p *Program) AddCommand(cmd *Command) {
	if p.startCommand == nil {
		p.current = cmd
		p.startCommand = cmd
		return
	}
	p.current.next = cmd
	p.current = cmd
	for _, lbl := range p.pendingLabels {
		//TODO: catch duplicate!
		p.labels[lbl] = cmd
	}
	p.pendingLabels = []string{}
}

//Performs some action on a machine.
//If it returns a string, execution jumps to that label.
//empty string, execution continues to next command
type Action func(*machine.PietMachine) string

type Command struct {
	next   *Command
	action Action
}
