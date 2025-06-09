package command

import "strings"

type Command struct {
	Name string   // command name
	Args []string // command args
}

// ParseCommand parses a command from a list of arguments.
func ParseCommand(args []string) *Command {
	if len(args) == 0 {
		return nil
	}

	cmd := &Command{
		Name: args[0],
		Args: make([]string, 0),
	}

	// handle quoted arguments
	inQuotes := false
	var current strings.Builder

	for _, arg := range args[1:] {
		// if the argument starts with a quote, it is a quoted argument
		if strings.HasPrefix(arg, "\"") {
			inQuotes = true
			current.WriteString(arg[1:])
			continue
		}

		// if the argument ends with a quote, it is a quoted argument
		if strings.HasSuffix(arg, "\"") {
			inQuotes = false
			current.WriteString(arg[:len(arg)-1])
			cmd.Args = append(cmd.Args, current.String())
			current.Reset()
			continue
		}

		// if the argument is not quoted, it is a normal argument
		if inQuotes {
			current.WriteString(arg)
		} else {
			cmd.Args = append(cmd.Args, arg)
		}
	}

	return cmd
}
