package clicommand

import (
	"fmt"
)

func (cmd *Command) Help(data *Data) {
	fmt.Printf("\n")
	fmt.Printf("%s\n", cmd.GetCommandNameChain())
	fmt.Printf("%s\n", cmd.desc)
	fmt.Printf("\n")

	cmd.HelpOptionsRecurseRev()

	if len(cmd.children) > 0 {
		fmt.Printf("Available subcommands:\n")
		for _, v := range cmd.children {
			fmt.Printf("  %-12s %s\n", v.name, v.desc)
		}
		fmt.Printf("\n")
	}

	if cmd.handler == nil {
		fmt.Printf("For help information run:\n")
		fmt.Printf("  '%s help'\n", cmd.GetCommandNameTop())
		fmt.Printf("  '%s <commands>* help'\n", cmd.GetCommandNameTop())
		fmt.Printf("  '%s [commands]* help [subcommand]*'\n", cmd.GetCommandNameTop())
		fmt.Printf("\n")
	}
}

func (cmd *Command) HelpOptionsRecurseRev() {
	if cmd.parent != nil {
		cmd.parent.HelpOptionsRecurseRev()
	}

	cmd.HelpOptions()
}

func (cmd *Command) HelpOptions() {
	if len(cmd.args) == 0 {
		return
	}

	fmt.Printf("%s options:\n", cmd.GetCommandNameChain())
	for _, arg := range cmd.args {
		var prefix string
		var suffix string

		if arg.param {
			prefix = "--"
			suffix = " <arg>"
		} else {
			prefix = "-"
		}

		fmt.Printf("  %2s%-20s %s\n", prefix, arg.name+suffix, arg.desc)
	}

	fmt.Printf("\n")
}
