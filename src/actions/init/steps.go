package actions

import (
	"log"
	"os"
	"path/filepath"
	// "os/exec"

	"github.com/go-git/go-git/v5"
	"gopkg.in/yaml.v2"
)

func createDirAndConfig(answers Answers) {
	porterrc := PorterRc{}
	porterrc.Resource.Type = answers.ResourceType
	porterrc.Resource.Name = answers.ResourceName
	porterrc.Owner.Type = answers.OwnerType
	porterrc.Owner.Github = answers.OwnerGithub
	porterrc.Owner.ReleasePublic = answers.ReleasePublic
	porterrc.Owner.ReleaseNamespace = answers.ReleaseNamespace

	d, err := yaml.Marshal(&porterrc)
	if err != nil {
		log.Fatal("error: %v", err)
	}

	err = os.Mkdir("./"+answers.ResourceName, 0755)
	if err != nil {
		log.Fatal("error: %v", err)
	}

	f, err := os.Create("./" + answers.ResourceName + "/.porterrc.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(`## created by opbi/porter\n` + string(d))
	if err != nil {
		log.Fatal(err)
	}
}

func initGit(answers Answers) {
	_, err := git.PlainInit("./"+answers.ResourceName, false)
	if err != nil {
		log.Fatal(err)
	}
}

func cdToDir(answers Answers) {
	wd, _ := os.Getwd()
	err := os.Chdir(filepath.Join(wd, answers.ResourceName))
	if err != nil {
		log.Fatal(err)
	}
}
