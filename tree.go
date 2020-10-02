package main

import (
	"github.com/thatisuday/commando"
)

func RegisterExecutable() {
	// configure commando
	commando.
		SetExecutableName("tree").
		SetVersion("1.0.0").
		SetDescription("This tool lists the contents of a directory in tree-like format.\nIt can also display information about files and folders like size, permission and ownership.")
}

func RegisterCommand() {
	// configure the root command
	commando.
		Register(nil).
		AddArgument("dir", "local directory path", "./").                     // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, 1).      // default `1`
		AddFlag("size", "display size of the each file", commando.Bool, nil). // default `false`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir := args["dir"].Value
			level, _ := flags["level"].GetInt()
			size, _ := flags["size"].GetBool()

			list(false, dir, level, size)
		})
}

func RegisterSubcommand() {
	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").                  // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir := args["dir"].Value
			level, _ := flags["level"].GetInt()

			list(true, dir, level, false)
		})
}

func main() {
	RegisterExecutable()
	RegisterCommand()
	RegisterSubcommand()
	// parse command-line arguments
	commando.Parse(nil)

}
