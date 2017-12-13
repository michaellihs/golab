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

package cmd

import (
	"testing"
	"bytes"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"net/http"
	"fmt"
	"io"
	"strings"
	"io/ioutil"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmd Suite")
}

// Helper funcs for testing...

// see https://github.com/spf13/cobra/blob/master/command_test.go for basic implementation of this method
func executeCommand(root *Command, args ...string) (stdout string, output string, err error) {
	stdout, output, err = executeCommandC(root, args...)
	return strings.TrimRight(stdout, "\n"), output, err
}

// see https://github.com/spf13/cobra/blob/master/command_test.go for basic implementation of this method
func executeCommandC(root *Command, args ...string) (stdout string, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf) // this only redirects stderr, not stdout!
	root.SetArgs(args)

	// for capturing stdout, see https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	_, err = root.ExecuteC()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	stdout = <-outC

	return stdout, buf.String(), err
}

func resetCommandLineFlagSet() {
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
}

func testMethod(r *http.Request, want string) {
	if got := r.Method; got != want {
		Fail(fmt.Sprintf("Request method: %s, want %s", got, want))
	}
}

func readFixture(fixture string) string {
	path := fmt.Sprintf("fixtures/%s.json", fixture)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Fail(fmt.Sprintf("Fixture %s could not be found in %s", fixture, path))
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Fail(err.Error())
	}
	return string(content)
}
