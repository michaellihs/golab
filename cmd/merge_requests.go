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

// see https://docs.gitlab.com/ce/api/merge_requests.html#create-mr
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

// see https://docs.gitlab.com/ce/api/merge_requests.html#delete-a-merge-request
type mergeRequestsDeleteFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsDeleteCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsDeleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete a merge request",
		Long:  `Only for admins and project owners. Soft deletes the merge request in question.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsDeleteFlags)
		_, err := gitlabClient.MergeRequests.DeleteMergeRequest(*flags.Id, *flags.MergeRequestIid)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#accept-mr
type mergeRequestAcceptFlags struct {
	Id                        *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	MergeRequestIid           *int    `flag_name:"merge_request_iid" short:"m" type:"int" required:"yes" description:"Internal ID of MR"`
	MergeCommitMessage        *string `flag_name:"merge_commit_message" type:"string" required:"no" description:"Custom merge commit message"`
	ShouldRemoveSourceBranch  *bool   `flag_name:"should_remove_source_branch" short:"d" type:"bool" required:"no" description:"if true removes the source branch"`
	MergeWhenPipelineSucceeds *bool   `flag_name:"merge_when_pipeline_succeeds" type:"bool" required:"no" description:"if true the MR is merged when the pipeline succeeds"`
	Sha                       *string `flag_name:"sha" type:"string" required:"no" description:"if present, then this SHA must match the HEAD of the source branch, otherwise the merge will fail"`
}

var mergeRequestAcceptCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestAcceptFlags{},
	Opts:   &gitlab.AcceptMergeRequestOptions{},
	Cmd: &cobra.Command{
		Use:   "accept",
		Short: "Accept merge request",
		Long: `Merge changes submitted with MR using this API.

If it has some conflicts and can not be merged - you'll get a 405 and the error message 'Branch cannot be merged'

If merge request is already merged or closed - you'll get a 406 and the error message 'Method Not Allowed'

If the sha parameter is passed and does not match the HEAD of the source - you'll get a 409 and the error message 'SHA does not match HEAD of source branch'

If you don't have permissions to accept this merge request - you'll get a 401`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestAcceptFlags)
		opts := cmd.Opts.(*gitlab.AcceptMergeRequestOptions)
		mr, _, err := gitlabClient.MergeRequests.AcceptMergeRequest(*flags.Id, *flags.MergeRequestIid, opts)
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#cancel-merge-when-pipeline-succeeds
type mergeRequestsCancelPipelineSucceedsFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequetsCancelPipelineSucceedsCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsCancelPipelineSucceedsFlags{},
	Cmd: &cobra.Command{
		Use:   "cancel-when-pipeline-succeeds",
		Short: "Cancel Merge When Pipeline Succeeds",
		Long: `If you don't have permissions to accept this merge request - you'll get a 401

If the merge request is already merged or closed - you get 405 and error message 'Method Not Allowed'

In case the merge request is not set to be merged when the pipeline succeeds, you'll also get a 406 error.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsCancelPipelineSucceedsFlags)
		mr, _, err := gitlabClient.MergeRequests.CancelMergeWhenPipelineSucceeds(*flags.Id, *flags.MergeRequestIid)
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#list-issues-that-will-close-on-merge
type mergeRequestsClosedIssuesUponMergeFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsClosedIssuesUponMergeCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsClosedIssuesUponMergeFlags{},
	Cmd: &cobra.Command{
		Use:   "list-issues",
		Short: "List issues that will close on merge",
		Long:  `Get all the issues that would be closed by merging the provided merge request.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsClosedIssuesUponMergeFlags)
		issues, _, err := gitlabClient.MergeRequests.GetIssuesClosedOnMerge(*flags.Id, *flags.MergeRequestIid)
		if err != nil {
			return err
		}
		return OutputJson(issues)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#subscribe-to-a-merge-request
type mergeRequestsSubscribeFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsSubscribeCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsSubscribeFlags{},
	Cmd: &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to a merge request",
		Long:  `Subscribes the authenticated user to a merge request to receive notification. If the user is already subscribed to the merge request, the status code 304 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsSubscribeFlags)
		mr, resp, err := gitlabClient.MergeRequests.Subscribe(*flags.Id, *flags.MergeRequestIid)
		if resp.StatusCode == 304 {
			return errors.New("304: the user was already subscribed to the merge request")
		}
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#unsubscribe-from-a-merge-request
type mergeRequestsUnsubscribeFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsUnsubscribeCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsUnsubscribeFlags{},
	Cmd: &cobra.Command{
		Use:   "unsubscribe",
		Short: "Unsubscribe from a merge request",
		Long:  `Unsubscribes the authenticated user from a merge request to not receive notification. If the user is not subscribed to the merge request, the status code 304 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsUnsubscribeFlags)
		mr, resp, err := gitlabClient.MergeRequests.Unsubscribe(*flags.Id, *flags.MergeRequestIid)
		if resp.StatusCode == 304 {
			return errors.New("304: the user was not subscribed to the merge request")
		}
		if err != nil {
			return err
		}
		return OutputJson(mr)
	},
}

// see https://docs.gitlab.com/ce/api/merge_requests.html#create-a-todo
type mergeRequestsCreateTodoFlags struct {
	Id              *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
	MergeRequestIid *int    `flag_name:"iid" short:"m" type:"integer" required:"yes" description:"The internal ID of the merge request"`
}

var mergeRequestsCreateTodoCmd = &golabCommand{
	Parent: mergeRequestsCmd,
	Flags:  &mergeRequestsCreateTodoFlags{},
	Cmd: &cobra.Command{
		Use:     "create-todo",
		Aliases: []string{"todo"},
		Short:   "Create a todo",
		Long:    `Manually creates a todo for the current user on a merge request. If there already exists a todo for the user on that merge request, status code 304 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*mergeRequestsCreateTodoFlags)
		mr, resp, err := gitlabClient.MergeRequests.CreateTodo(*flags.Id, *flags.MergeRequestIid)
		if resp.StatusCode == 304 {
			return errors.New("304: there already exists a todo on that merge request for the current user")
		}
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
	mergeRequestsDeleteCmd.Init()
	mergeRequestAcceptCmd.Init()
	mergeRequetsCancelPipelineSucceedsCmd.Init()
	mergeRequestsClosedIssuesUponMergeCmd.Init()
	mergeRequestsSubscribeCmd.Init()
	mergeRequestsUnsubscribeCmd.Init()
	mergeRequestsCreateTodoCmd.Init()
	RootCmd.AddCommand(mergeRequestsCmd)
}
