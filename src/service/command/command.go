package command

import (
	"fmt"
)

type Command struct {
	Name         string
	Usage        string
	Action       any
	DefaultValue any
}

func (c *Command) String() string {
	return fmt.Sprintf("%s", c.DefaultValue)
}

func (c *Command) Set(value string) error {
	return nil
}
func NewCommand(name string, action, defaultValue any, usage string) *Command {
	comm := &Command{
		Name:         name,
		Usage:        usage,
		Action:       action,
		DefaultValue: defaultValue,
	}
	return comm
}
