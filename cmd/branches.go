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
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

// see https://docs.gitlab.com/ce/api/branches.html#branches-api
var branchesCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "branches",
		Aliases: []string{"alias"},
		Short:   "Branches",
		Long:    `Manage repository branches`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot use this command without further sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#list-repository-branches
type branchesGetFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
}

var branchesGetCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesGetFlags{},
	Opts:   &gitlab.ListBranchesOptions{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "List repository branches",
		Long:  `Get a list of repository branches from a project, sorted by name alphabetically. This endpoint can be accessed without authentication if the repository is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesGetFlags)
		opts := cmd.Opts.(*gitlab.ListBranchesOptions)
		branches, _, err := gitlabClient.Branches.ListBranches(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(branches)
	},
}

func init() {
	branchesCmd.Init()
	branchesGetCmd.Init()
}
