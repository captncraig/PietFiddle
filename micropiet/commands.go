package micropiet

import (
	"github.com/captncraig/pietfiddle/machine"
	"strconv"
)

func NewPushCommand(num string) *Command {
	val, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		panic(err)
	}
	return &Command{
		action: func(vm *machine.PietMachine) string {
			vm.Push(val)
			return ""
		}}
}
func NewArithmeticCommand(op string) *Command {
	return &Command{
		action: func(vm *machine.PietMachine) string {
			vm.Binary(op)
			return ""
		}}
}
func NewPopCommand() *Command {
	return &Command{
		action: func(vm *machine.PietMachine) string {
			vm.Pop()
			return ""
		}}
}
func NewNotCommand() *Command {
	return &Command{
		action: func(vm *machine.PietMachine) string {
			vm.Not()
			return ""
		}}
}

func NewDupCommand() *Command {
	return &Command{
		action: func(vm *machine.PietMachine) string {
			vm.Dup()
			return ""
		}}
}
