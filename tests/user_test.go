// Copyright © 2017 Michael Lihs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tests

import (
	"testing"
	"os/exec"
	"fmt"
	"os"
	"encoding/json"
	"log"

	"github.com/xanzy/go-gitlab"
	"github.com/magiconair/properties/assert"
	"strconv"
)

var golabBinary = "golab"

// see http://lucapette.me/writing-integration-tests-for-a-go-cli-application for basic ideas of testing a CLI with golang
func TestMain(m *testing.M) {
	chdirToParent()
	os.Exit(m.Run())
}

func TestUserLs(t *testing.T) {
	args := []string{"user", "ls"}
	cmd := exec.Command(golabBinary, args...)
	out, err := cmd.CombinedOutput()
	fatalOnErr(err, out, t)
}

func TestUserCreateGetDelete(t *testing.T) {
	// create a user
	args :=  []string{"user", "create", "-u", "root10", "-p", "12341234", "-e", "test10@test.de",  "-n", "root10", "--skip_confirmation"}
	cmd := exec.Command(golabBinary, args...)
	out, err := cmd.CombinedOutput()
	fatalOnErr(err, out, t)

	// get it
	args = []string{"user", "get", "-u", "root10"}
	cmd = exec.Command(golabBinary, args...)
	stdout, err := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var user = gitlab.User{}
	fatalOnErr(err, out, t)
	if err := json.NewDecoder(stdout).Decode(&user); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, user.Username, "root10")
	assert.Equal(t, user.Email, "test10@test.de")

	// delete it again
	args =  []string{"user", "delete", "-i", strconv.Itoa(user.ID)}
	cmd = exec.Command(golabBinary, args...)
	out, err = cmd.CombinedOutput()
	fatalOnErr(err, out, t)
}

func fatalOnErr(err error, out []byte, t *testing.T) {
	if err != nil {
		fmt.Println(string(out))
		t.Fatal(err)
	}
}

func chdirToParent() {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
}