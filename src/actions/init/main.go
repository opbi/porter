package actions

import (
	"fmt"

	"github.com/thatisuday/commando"
	"gopkg.in/AlecAivazis/survey.v1"
)

func async(step func(Answers)) func(Answers) <-chan interface{} {
	return func(answers Answers) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			step(answers)
		}()
		return c
	}
}

func Init(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

	answers := Answers{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	createDirAndConfig(answers)
	initGit(answers)
	cdToDir(answers)

	fmt.Println("created new [" + answers.ResourceType + "]: " + answers.ResourceName + "\ncheck .porterrc.yml and run 'porter setup' to finish")
}
