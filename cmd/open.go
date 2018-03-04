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

package cmd

import (
	"github.com/michaellihs/golab/cmd/helpers"
	"github.com/spf13/cobra"
)

var openCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "open",
		Aliases: []string{"o"},
		Short:   "Open Gitlab for project",
		Long:    `Open the Gitlab project page for the repository in current directory`,
	},
	Run: func(cmd golabCommand) error {
		bh := helpers.NewBrowserHelper()
		if url, err := getRemoteUrl(); err == nil {
			return bh.Open(url)
		} else {
			return err
		}
	},
}

func getRemoteUrl() (string, error) {
	gh := helpers.GitHelper()
	remotes, err := gh.GetRemotes()
	if err != nil {
		return "", err
	}
	remote, err := gh.GetRemoteUrl(remotes)
	if err != nil {
		return "", err
	}
	url, err := gh.GetWebUrl(remote)
	if err != nil {
		return "", err
	}
	return url, nil
}

func init() {
	openCmd.Init()
}
