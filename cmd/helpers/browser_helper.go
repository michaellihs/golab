// Copyright © 2018 Michael Lihs
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

package helpers

import (
	"os/exec"
	"runtime"

	"errors"
)

type BrowserHelper struct{}

func NewBrowserHelper() BrowserHelper {
	return BrowserHelper{}
}

func (b BrowserHelper) Open(url string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("start", url).Run()
	} else if runtime.GOOS == "linux" {
		// see https://stackoverflow.com/questions/5116473/linux-command-to-open-url-in-default-browser
		return exec.Command("xdg-open", url).Run()
	} else if runtime.GOOS == "darwin" {
		return exec.Command("open", url).Run()
	}
	return errors.New("Unknown OS: " + runtime.GOOS)
}
