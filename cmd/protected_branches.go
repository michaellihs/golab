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
	"errors"
	"github.com/spf13/cobra"
)

// see https://docs.gitlab.com/ce/api/protected_branches.html#protected-branches-api
var protectedBranchesCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:   "protected-branches",
		Aliases: []string{"pb"},
		Short: "Protected branches",
		Long:  `Manage protected branches`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("you cannot run this command without a sub-command")
	},
}

// see https://docs.gitlab.com/ce/api/protected_branches.html#list-protected-branches
type protectedBranchesListFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
}

var protectedBranchesListCmd = &golabCommand{
	Parent: protectedBranchesCmd.Cmd,
	Flags:  &protectedBranchesListFlags{},
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List protected branches",
		Long:  `Gets a list of protected branches from a project.`,
	},
	Run: func(cmd golabCommand) error {
		// TODO implement in go-gitlab
		return errors.New("currently not implemented by go-gitlab")
		//flags := cmd.Flags.(*protectedBranchesListFlags)
		//branches, _, err := gitlabClient.Branches.ListProtectedBranches(*flags.Id)
		//if err != nil {
		//    return err
		//}
		//return OutputJson(branches)
	},
}

func init() {
	protectedBranchesCmd.Init()
	protectedBranchesListCmd.Init()
}