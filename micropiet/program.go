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
	return &Program{}
}

func (p *Program) Run(vm *machine.PietMachine) {
	cur := p.startCommand
	for cur != nil {
		cur.action(vm)
		cur = cur.next
	}
}

func (p *Program) AddCommand(cmd *Command) {
	if p.startCommand == nil {
		p.current = cmd
		p.startCommand = cmd
		return
	}
	p.current.next = cmd
	p.current = cmd
}

//Performs some action on a machine.
//If it returns a string, execution jumps to that label.
//empty string, execution continues to next command
type Action func(*machine.PietMachine) string

type Command struct {
	next   *Command
	action Action
}
