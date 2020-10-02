package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {

	// test command
	cmd := exec.Command("go", "run", ".", "--no-color")

	// expected output
	lines := []string{
		`├── .editorconfig`,
		`├── actions`,
		`├── go.mod`,
		`├── go.sum`,
		`├── main.go`,
		`└── main_test.go`,
	}

	// get output
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		outputString := fmt.Sprintf("%s", output)

		for _, line := range lines {
			if !strings.Contains(outputString, line) {
				fmt.Println("line: ", line)
				t.Fail()
			}
		}
	}
}
