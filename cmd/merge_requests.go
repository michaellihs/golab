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
	"github.com/michaellihs/golab/cmd/mapper"
	"github.com/xanzy/go-gitlab"
)

var listMergeRequestFlagsMapper mapper.FlagMapper

var mergeRequestsCmd = &cobra.Command{
	Use:   "merge-requests",
	Short: "Manage Merge Requests",
	Long:  `Show, create, edit and delte Merge Requests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("this command cannot be run without a sub-command")
	},
}

type listMergeRequestsFlags struct {
	State           *string `flag_name:"state" type:"string" required:"no" description:"Return all merge requests or just those that are opened, closed, or merged"`
	OrderBy         *string `flag_name:"order_by" type:"string" required:"no" description:"Return requests ordered by created_at or updated_at fields. Default is created_at"`
	Sort            *string `flag_name:"sort" type:"string" required:"no" description:"Return requests sorted in asc or desc order. Default is desc"`
	Milestone       *string `flag_name:"milestone" type:"string" required:"no" description:"Return merge requests for a specific milestone"`
	View            *string `flag_name:"view" type:"string" required:"no" description:"If simple, returns the iid, URL, title, description, and basic state of merge request"`
	Labels          *string `flag_name:"labels" type:"string" required:"no" description:"Return merge requests matching a comma separated list of labels"`
	CreatedAfter    *string `flag_name:"created_after" type:"datetime" required:"no" description:"Return merge requests created after the given time (inclusive)"`
	CreatedBefore   *string `flag_name:"created_before" type:"datetime" required:"no" description:"Return merge requests created before the given time (inclusive)"`
	Scope           *string `flag_name:"scope" type:"string" required:"no" description:"Return merge requests for the given scope: created-by-me, assigned-to-me or all. Defaults to created-by-me"`
	AuthorId        *int    `flag_name:"author_id" type:"integer" required:"no" description:"Returns merge requests created by the given user id. Combine with scope=all or scope=assigned-to-me"`
	AssigneeId      *int    `flag_name:"assignee_id" type:"integer" required:"no" description:"Returns merge requests assigned to the given user id"`
	MyReactionEmoji *string `flag_name:"my_reaction_emoji" type:"string" required:"no" description:"Return merge requests reacted by the authenticated user by the given emoji (Introduced in GitLab 10.0)"`
}

var mergeRequestListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List merge requests",
	Long: `Get all merge requests the authenticated user has access to. By default it returns only merge requests created by the current user. To get all merge requests, use parameter scope=all.

The state parameter can be used to get only merge requests with a given state (opened, closed, or merged) or all of them (all). The pagination parameters page and per_page can be used to restrict the list of merge requests.

Note: the changes_count value in the response is a string, not an integer. This is because when an MR has too many changes to display and store, it will be capped at 1,000. In that case, the API will return the string "1000+" for the changes count.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listMergeRequestFlagsMapper.AutoMap()
		opts := listMergeRequestFlagsMapper.MappedOpts().(*gitlab.ListMergeRequestsOptions)
		mergeRequests, _, err := gitlabClient.MergeRequests.ListMergeRequests(opts)
		if err != nil {
			return err
		}
		return OutputJson(mergeRequests)
	},
}

func init() {
	initListMergeRequestCmd()
	RootCmd.AddCommand(mergeRequestsCmd)
}

func initListMergeRequestCmd() {
	listMergeRequestFlagsMapper = mapper.InitializedMapper(mergeRequestListCmd, &listMergeRequestsFlags{}, &gitlab.ListMergeRequestsOptions{})
	mergeRequestsCmd.AddCommand(mergeRequestListCmd)
}
