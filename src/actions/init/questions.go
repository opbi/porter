package actions

import (
	"github.com/opbi/porter/src/constants"
	"gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name: "resourceType",
		Prompt: &survey.Select{
			Message: "Which type of resource do you want to create?",
			Options: constants.RESOURCE_TYPES,
		},
	},
	{
		Name: "resourceName",
		Prompt: &survey.Input{
			Message: "Enter the name of the resource:",
		},
		Validate: survey.Required,
	},
	{
		Name: "ownerType",
		Prompt: &survey.Select{
			Message: "What is the type of the owner of this?",
			Options: constants.OWNER_TYPES,
		},
	},
	{
		Name: "ownerGithub",
		Prompt: &survey.Input{
			Message: "Enter the Github ID of the owner:",
		},
	},
	{
		Name: "releasePublic",
		Prompt: &survey.Confirm{
			Message: "Is the resource open to the public?",
			Help:    "use the default option of not public if you are not sure",
			Default: false,
		},
	},
	{
		Name: "releaseNamespace",
		Prompt: &survey.Input{
			Message: "What's the namespace where the package would be distributed?",
		},
	},
}
