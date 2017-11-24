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
	"os"
	"errors"
)

func TestLogin(t *testing.T) {
	host := os.Getenv("GITLAB_HOST")
	if host == "" { fatalOnErr(errors.New("environment variable `GITLAB_HOST` has to be set to run this test (e.g. http://hostname:port)"), nil, t)}
	user := os.Getenv("GITLAB_ROOT_USER")
	if user == "" { fatalOnErr(errors.New("environment variable `GITLAB_ROOT_USER`has to be set to run this test"), nil, t)}
	password := os.Getenv("GITLAB_ROOT_PASSWORD")
	if password == "" { fatalOnErr(errors.New("environment variable `GITLAB_ROOT_PASSWORD`has to be set to run this test"), nil, t)}

	args := []string{"login", "--host", host, "--user", user, "--password", password}
	cmd := exec.Command(golabBinary, args...)
	out, err := cmd.CombinedOutput()
	fatalOnErr(err, out, t)
}
