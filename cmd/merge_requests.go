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

// see https://docs.gitlab.com/ce/api/merge_requests.html#merge-requests-api
var mergeRequestsCmd = &cobra.Command{
	Use:     "merge-requests",
	Aliases: []string{"mr"},
	Short:   "Manage Merge Requests",
	Long:    `Show, create, edit and delte Merge Requests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("this command cannot be run without a sub-command")
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#list-merge-requests
type mergeRequestsListFlags struct {
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

var mergeRequestsListCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsListFlags{},
	Opts:   &gitlab.ListMergeRequestsOptions{},
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List merge requests",
		Long: `Get all merge requests the authenticated user has access to. By default it returns only merge requests created by the current user. To get all merge requests, use parameter scope=all.

The state parameter can be used to get only merge requests with a given state (opened, closed, or merged) or all of them (all). The pagination parameters page and per_page can be used to restrict the list of merge requests.

Note: the changes_count value in the response is a string, not an integer. This is because when an MR has too many changes to display and store, it will be capped at 1,000. In that case, the API will return the string "1000+" for the changes count.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.ListMergeRequestsOptions)
		mergeRequests, _, err := gitlabClient.MergeRequests.ListMergeRequests(opts)
		if err != nil {
			return err
		}
		return OutputJson(mergeRequests)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#list-project-merge-requests
type mergeRequestsListForProjectFlags struct {
	Id              *int    `flag_name:"id" type:"integer" required:"yes" description:"The ID of a project"`
	IIDs            []int   `flag_name:"iids" type:"Array[integer]" required:"no" description:"Return the request having the given iid"`
	State           *string `flag_name:"state" type:"string" required:"no" description:"Return all merge requests or just those that are opened, closed, or merged"`
	OrderBy         *string `flag_name:"order_by" type:"string" required:"no" description:"Return requests ordered by created_at or updated_at fields. Default is created_at"`
	Sort            *string `flag_name:"sort" type:"string" required:"no" description:"Return requests sorted in asc or desc order. Default is desc"`
	Milestone       *string `flag_name:"milestone" type:"string" required:"no" description:"Return merge requests for a specific milestone"`
	View            *string `flag_name:"view" type:"string" required:"no" description:"If simple, returns the iid, URL, title, description, and basic state of merge request"`
	Labels          *string `flag_name:"labels" type:"[]string" transform:"string2Labels" required:"no" description:"Return merge requests matching a comma separated list of labels"`
	CreatedAfter    *string `flag_name:"created_after" type:"datetime" required:"no" description:"Return merge requests created after the given time (inclusive)"`
	CreatedBefore   *string `flag_name:"created_before" type:"datetime" required:"no" description:"Return merge requests created before the given time (inclusive)"`
	Scope           *string `flag_name:"scope" type:"string" required:"no" description:"Return merge requests for the given scope: created-by-me, assigned-to-me or all (Introduced in GitLab 9.5)"`
	AuthorID        *int    `flag_name:"author_id" type:"integer" required:"no" description:"Returns merge requests created by the given user id (Introduced in GitLab 9.5)"`
	AssigneeID      *int    `flag_name:"assignee_id" type:"integer" required:"no" description:"Returns merge requests assigned to the given user id (Introduced in GitLab 9.5)"`
	MyReactionEmoji *string `flag_name:"my_reaction_emoji" type:"string" required:"no" description:"Return merge requests reacted by the authenticated user by the given emoji (Introduced in GitLab 10.0)"`
}

var mergeRequestsListForProjectCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsListForProjectFlags{},
	Opts:   &gitlab.ListProjectMergeRequestsOptions{},
	Cmd: &cobra.Command{
		Use:   "project-ls",
		Short: "List project merge requests",
		Long:  `Get all merge requests for this project. The state parameter can be used to get only merge requests with a given state (opened, closed, or merged) or all of them (all). The pagination parameters page and per_page can be used to restrict the list of merge requests.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsListForProjectFlags)
		opts := cmd.Opts.(*gitlab.ListProjectMergeRequestsOptions)
		mrs, _, err := gitlabClient.MergeRequests.ListProjectMergeRequests(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(mrs)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#get-single-mr
type mergeRequestGetFlags struct {
	Id  *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	Iid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestGetCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get single Merge Request",
		Long:  `Shows information about a single merge request.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestGetFlags)
		mr, _, err := gitlabClient.MergeRequests.GetMergeRequest(*flags.Id, *flags.Iid)
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#get-single-mr-commits
type mergeRequestGetCommitsFlags struct {
	Id  *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	Iid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsGetCommitsCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestGetCommitsFlags{},
	Cmd: &cobra.Command{
		Use:   "get-commits",
		Short: "Get single Merge Request commits",
		Long:  `Get a list of merge request commits.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestGetCommitsFlags)
		commits, _, err := gitlabClient.MergeRequests.GetMergeRequestCommits(*flags.Id, *flags.Iid)
		if err != nil {
			return err
		}
		return OutputJson(commits)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#get-single-mr-changes
type mergeRequestsGetChangesFlags struct {
	Id  *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	Iid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsGetChangesCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsGetChangesFlags{},
	Cmd: &cobra.Command{
		Use:   "get-changes",
		Short: "Get single Merge Request changes",
		Long:  `Shows information about the merge request including its files and changes.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsGetChangesFlags)
		changes, _, err := gitlabClient.MergeRequests.GetMergeRequestChanges(*flags.Id, *flags.Iid)
		if err != nil {
			return err
		}
		return OutputJson(changes)
	},
}

// see 
type mergeRequestsCreateFlags struct {
	Id                 *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	SourceBranch       *string `flag_name:"source_branch" short:"s" type:"string" required:"yes" description:"The source branch"`
	TargetBranch       *string `flag_name:"target_branch" short:"t" type:"string" required:"yes" description:"The target branch"`
	Title              *string `flag_name:"title" short:"n" type:"string" required:"yes" description:"Title of MR"`
	AssigneeId         *int    `flag_name:"assignee_id" short:"a" type:"integer" required:"no" description:"Assignee user ID"`
	Description        *string `flag_name:"description" short:"d" type:"string" required:"no" description:"Description of MR"`
	TargetProjectId    *int    `flag_name:"target_project_id" type:"integer" required:"no" description:"The target project (numeric id)"`
	Labels             *string `flag_name:"labels" type:"[]string" transform:"string2Labels" required:"no" description:"Labels for MR as a comma-separated list"`
	MilestoneId        *int    `flag_name:"milestone_id" type:"integer" required:"no" description:"The ID of a milestone"`
	RemoveSourceBranch *bool   `flag_name:"remove_source_branch" type:"boolean" required:"no" description:"Flag indicating if a merge request should remove the source branch when merging"`
}

var mergeRequestsCreateCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsCreateFlags{},
	Opts:   &gitlab.CreateMergeRequestOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create merge request",
		Long:  `Creates a new merge request.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateMergeRequestOptions)
		mr, _, err := gitlabClient.MergeRequests.CreateMergeRequest(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#update-mr
type mergeRequestUpdateFlags struct {
	Id                 *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	MergeRequestIid    *int    `flag_name:"merge_request_iid" short:"m" type:"integer" required:"yes" description:"The ID of a merge request"`
	TargetBranch       *string `flag_name:"target_branch" type:"string" required:"no" description:"The target branch"`
	Title              *string `flag_name:"title" type:"string" required:"no" description:"Title of MR"`
	AssigneeId         *int    `flag_name:"assignee_id" type:"integer" required:"no" description:"Assignee user ID"`
	Description        *string `flag_name:"description" type:"string" required:"no" description:"Description of MR"`
	StateEvent         *string `flag_name:"state_event" type:"string" required:"no" description:"New state (close/reopen)"`
	Labels             *string `flag_name:"labels" type:"[]string" transform:"string2Labels" required:"no" description:"Labels for MR as a comma-separated list"`
	MilestoneId        *int    `flag_name:"milestone_id" type:"integer" required:"no" description:"The ID of a milestone"`
	RemoveSourceBranch *bool   `flag_name:"remove_source_branch" type:"boolean" required:"no" description:"Flag indicating if a merge request should remove the source branch when merging"`
	DiscussionLocked   *bool   `flag_name:"discussion_locked" type:"boolean" required:"no" description:"Flag indicating if the merge request's discussion is locked. If the discussion is locked only project members can add, edit or resolve comments."`
}

var mergeRequestUpdateCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestUpdateFlags{},
	Opts:   &gitlab.UpdateMergeRequestOptions{},
	Cmd: &cobra.Command{
		Use:   "update",
		Short: "Update merge request",
		Long:  `Updates an existing merge request. You can change the target branch, title, or even close the MR.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestUpdateFlags)
		opts := cmd.Opts.(*gitlab.UpdateMergeRequestOptions)
		mr, _, err := gitlabClient.MergeRequests.UpdateMergeRequest(*flags.Id, *flags.MergeRequestIid, opts)
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

func init() {
	mergeRequestsListCmd.Init()
	mergeRequestsListForProjectCmd.Init()
	mergeRequestGetCmd.Init()
	mergeRequestsGetCommitsCmd.Init()
	mergeRequestsGetChangesCmd.Init()
	mergeRequestsCreateCmd.Init()
	mergeRequestUpdateCmd.Init()
	RootCmd.AddCommand(mergeRequestsCmd)
}
