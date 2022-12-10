package day10

type commandName int

const (
	noop commandName = iota
	addX
)

type addXCommand struct {
	value int
}

func (axc addXCommand) getCommandName() commandName {
	return addX
}

type noopCommand struct {
}

func (nopc noopCommand) getCommandName() commandName {
	return noop
}

type command interface {
	getCommandName() commandName
}
