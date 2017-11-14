package tests

import (
	"testing"
	"os/exec"
	"path"
	"fmt"
	"os"
)

// see http://lucapette.me/writing-integration-tests-for-a-go-cli-application for basic ideas of testing a CLI with golang

func TestUserLs(t *testing.T) {

	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

    binaryName := "golab"
    args := []string{"user", "ls"}
	cmd := exec.Command(path.Join(".", binaryName), args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		t.Fatal(err)
	}

}