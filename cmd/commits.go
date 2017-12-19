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
	"github.com/xanzy/go-gitlab"
)

// see https://docs.gitlab.com/ce/api/commits.html#commits-api
var commitsCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:   "commits",
		Short: "Manage Commits",
		Long:  `Manage Commits`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("this command cannot be used without any sub-command")
	},
}

// see https://docs.gitlab.com/ce/api/commits.html#list-repository-commits
type commitsListFlags struct {
	Id      *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	RefName *string `flag_name:"ref_name" short:"r" type:"string" required:"no" description:"The name of a repository branch or tag or if not given the default branch"`
	Since   *string `flag_name:"since" transform:"string2TimeVal" short:"s" type:"string" required:"no" description:"Only commits after or on this date will be returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ"`
	Until   *string `flag_name:"until" transform:"string2TimeVal" short:"u" type:"string" required:"no" description:"Only commits before or on this date will be returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ"`
}

var commitsListCmd = &golabCommand{
	Parent: commitsCmd.Cmd,
	Flags:  &commitsListFlags{},
	Opts:   &gitlab.ListCommitsOptions{},
	Cmd: &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List repository commits",
		Long:    `Get a list of repository commits in a project`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*commitsListFlags)
		opts := cmd.Opts.(*gitlab.ListCommitsOptions)
		c, _, err := gitlabClient.Commits.ListCommits(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(c)
	},
}

// see https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
type commitsCreateFlags struct {
	Id            *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Branch        *string `flag_name:"branch" type:"string" required:"yes" description:"Name of the branch to commit into. To create a new branch, also provide start_branch."`
	CommitMessage *string `flag_name:"commit_message" type:"string" required:"yes" description:"Commit message"`
	Actions       *string `flag_name:"actions" transform:"json2CommitActions" type:"array" required:"yes" description:"A JSON encoded array of action hashes to commit as a batch."`
	AuthorEmail   *string `flag_name:"author_email" type:"string" required:"no" description:"Specify the commit author's email address"`
	AuthorName    *string `flag_name:"author_name" type:"string" required:"no" description:"Specify the commit author's name"`
	// TODO this field is currently not available in go-gitlab
	// StartBranch   *string `flag_name:"start_branch" type:"string" required:"no" description:"Name of the branch to start the new commit from"`
}

var commitsCreateCmd = &golabCommand{
	Parent: commitsCmd.Cmd,
	Flags:  &commitsCreateFlags{},
	Opts:   &gitlab.CreateCommitOptions{},
	Cmd: &cobra.Command{
		Use:     "create",
		Aliases: []string{""},
		Short:   "Create a commit with multiple files and actions",
		Long: `Create a commit by posting a JSON payload

JSON encoded Actions:

	[
		{
		  "action": "create",
		  "file_path": "foo/bar",
		  "content": "some content"
		},
		{
		  "action": "delete",
		  "file_path": "foo/bar2"
		},
		{
		  "action": "move",
		  "file_path": "foo/bar3",
		  "previous_path": "foo/bar4",
		  "content": "some content"
		},
		{
		  "action": "update",
		  "file_path": "foo/bar5",
		  "content": "new content"
		}
    ]`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*commitsCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateCommitOptions)
		c, _, err := gitlabClient.Commits.CreateCommit(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(c)
	},
}

func init() {
	commitsCmd.Init()
	commitsListCmd.Init()
	commitsCreateCmd.Init()
}
