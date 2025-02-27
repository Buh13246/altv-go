package alt

type Invoker struct {
	commands map[string]Cmd
}

type Cmd interface {
	Name() string
	Execute(p IPlayer)
}

func (i Invoker) RegisterCmd(c Cmd) {
	_, exists := i.commands[c.Name()]

	if exists {
		LogError("Command already registered.")
		return
	}

	i.commands[c.Name()] = c
}

func (i Invoker) ExecuteCmd(p IPlayer, name string) {
	cmd, exists := i.commands[name]

	if !exists {
		return
	}

	cmd.Execute(p)
}
