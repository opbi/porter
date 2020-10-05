package main

import (
	"github.com/opbi/porter/src/actions/init"
	"github.com/thatisuday/commando"
)

func main() {

	commando.
		SetExecutableName("porter").
		SetVersion("1.0.1").
		SetDescription("the universal config manager to coordinate setup, toolings, ci and secret across repos for easy consistency")

	commando.
		Register("init").
		SetShortDescription("create the config specs of a new component").
		SetAction(actions.Init)

	// set input from ENV_VAR to nil
	// parse only the command-line arguments
	commando.Parse(nil)

}
