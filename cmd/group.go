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

	. "github.com/michaellihs/golab/cmd/helpers"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Gitlab Groups",
	Long:  `Show, create, update and delete Gitlab groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("check usage of group with `golab group -h`")
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#list-groups
type groupLsFlags struct {
	SkipGroups   *[]string `flag_name:"skip_groups" type:"array" required:"no" description:"Skip the group IDs passed"`
	AllAvailable *bool     `flag_name:"all_available" type:"bool" required:"no" description:"Show all the groups you have access to (defaults to false for authenticated users)"`
	Search       *string   `flag_name:"search" type:"string" required:"no" description:"Return the list of authorized groups matching the search criteria"`
	OrderBy      *string   `flag_name:"order_by" type:"string" required:"no" description:"Order groups by name or path. Default is name"`
	Sort         *string   `flag_name:"sort" type:"string" required:"no" description:"Order groups in asc or desc order. Default is asc"`
	Statistics   *bool     `flag_name:"statistics" type:"bool" required:"no" description:"Include group statistics (admins only)"`
	Owned        *bool     `flag_name:"owned" type:"boolean" required:"no" description:"Limit to groups owned by the current user"`
}

var groupLsCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupLsFlags{},
	Opts:   &gitlab.ListGroupsOptions{},
	Paged:  true,
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List groups",
		Long:  `Get a list of visible groups for the authenticated user.`,
	},
	Run: func(cmd golabCommand) error {
		groups, _, err := gitlabClient.Groups.ListGroups(cmd.Opts.(*gitlab.ListGroupsOptions))
		if err != nil {
			return err
		}
		return OutputJson(groups)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#list-a-group-39-s-projects
type listGroupProjectsFlags struct {
	Id         *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the group owned by the authenticated user"`
	Archived   *bool   `flag_name:"archived" type:"bool" required:"no" description:"Limit by archived status"`
	Visibility *string `flag_name:"visibility" type:"string" transform:"str2Visibility" required:"no" description:"Limit by visibility public, internal, or private"`
	OrderBy    *string `flag_name:"order_by" type:"string" required:"no" description:"Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at"`
	Sort       *string `flag_name:"sort" type:"string" required:"no" description:"Return projects sorted in asc or desc order. Default is desc"`
	Search     *string `flag_name:"search" type:"string" required:"no" description:"Return list of authorized projects matching the search criteria"`
	Simple     *bool   `flag_name:"simple" type:"bool" required:"no" description:"Return only the ID, URL, name, and path of each project"`
	Owned      *bool   `flag_name:"owned" type:"bool" required:"no" description:"Limit by projects owned by the current user"`
	Starred    *bool   `flag_name:"starred" type:"bool" required:"no" description:"Limit by projects starred by the current user"`
}

var groupProjectsCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &listGroupProjectsFlags{},
	Opts:   &gitlab.ListGroupProjectsOptions{},
	Paged:  true,
	Cmd: &cobra.Command{
		Use:   "projects",
		Short: "List a group's projects",
		Long:  `Get a list of projects in this group. When accessed without authentication, only public projects are returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*listGroupProjectsFlags)
		projects, _, err := gitlabClient.Groups.ListGroupProjects(*flags.Id, cmd.Opts.(*gitlab.ListGroupProjectsOptions))
		if err != nil {
			return err
		}
		return OutputJson(projects)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#details-of-a-group
type groupGetFlags struct {
	Id *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the group owned by the authenticated user"`
}

var groupGetCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Details of a group",
		Long:  `Get all details of a group. This command can be accessed without authentication if the group is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*groupGetFlags)
		group, _, err := gitlabClient.Groups.GetGroup(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#new-group
type groupCreateFlags struct {
	Name                 *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the group"`
	Path                 *string `flag_name:"path" short:"p" type:"string" required:"yes" description:"The path of the group"`
	Description          *string `flag_name:"description" type:"string" required:"no" description:"The group's description"`
	Visibility           *string `flag_name:"visibility" type:"string" transform:"str2Visibility" required:"no" description:"The group's visibility. Can be private, internal, or public."`
	LfsEnabled           *bool   `flag_name:"lfs_enabled" type:"bool" required:"no" description:"Enable/disable Large File Storage (LFS) for the projects in this group"`
	RequestAccessEnabled *bool   `flag_name:"request_access_enabled" type:"bool" required:"no" description:"- Allow users to request member access."`
	ParentId             *int    `flag_name:"parent_id" type:"int" required:"no" description:"The parent group id for creating nested group."`
}

var groupCreateCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupCreateFlags{},
	Opts:   &gitlab.CreateGroupOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "New group",
		Long:  `Creates a new project group. Available only for users who can create groups.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.CreateGroupOptions)
		group, _, err := gitlabClient.Groups.CreateGroup(opts)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#transfer-project-to-group
type transferProjectFlags struct {
	Id        *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the group owned by the authenticated user"`
	ProjectId *int    `flag_name:"project_id" short:"p" type:"string" required:"yes" description:"The ID or URL-encoded path of a project"`
	// TODO go-gitlab does not support ID or URL-encoded path here
	// ProjectId *string `flag_name:"project_id" short:"p" type:"string" required:"yes" description:"The ID or URL-encoded path of a project"`
}

var transferProjectCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &transferProjectFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "transfer-project",
		Short: "Transfer project to group",
		Long:  `Transfer a project to the Group namespace. Available only for admin.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*transferProjectFlags)
		group, _, err := gitlabClient.Groups.TransferGroup(*flags.Id, *flags.ProjectId)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#update-group
type groupUpdateFlags struct {
	Id                   *int    `flag_name:"id" type:"integer" required:"yes" description:"The ID of the group"`
	Name                 *string `flag_name:"name" type:"string" required:"no" description:"The name of the group"`
	Path                 *string `flag_name:"path" type:"string" required:"no" description:"The path of the group"`
	Description          *string `flag_name:"description" type:"string" required:"no" description:"The description of the group"`
	Visibility           *string `flag_name:"visibility" type:"string" transform:"str2Visibility" required:"no" description:"The visibility level of the group. Can be private, internal, or public."`
	LfsEnabled           *bool   `flag_name:"lfs_enabled" type:"boolean" required:"no" description:"Enable/disable Large File Storage (LFS) for the projects in this group"`
	RequestAccessEnabled *bool   `flag_name:"request_access_enabled" type:"boolean" required:"no" description:"Allow users to request member access."`
}

var groupUpdateCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupUpdateFlags{},
	Opts:   &gitlab.UpdateGroupOptions{},
	Cmd: &cobra.Command{
		Use:   "update",
		Short: "Update group",
		Long:  `Updates the project group. Only available to group owners and administrators.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*groupUpdateFlags)
		opts := cmd.Opts.(*gitlab.UpdateGroupOptions)
		group, _, err := gitlabClient.Groups.UpdateGroup(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#remove-group
type groupDeleteFlags struct {
	Id *string `flag_name:"id" type:"string" required:"yes" description:"The ID or URL encoded path of a user group"`
}

var groupDeleteCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupDeleteFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Remove group",
		Long:  `Removes group with all projects inside.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*groupDeleteFlags)
		_, err := gitlabClient.Groups.DeleteGroup(*flags.Id)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#search-for-group
type groupSearchFlags struct {
	Search *string `flag_name:"search" short:"s" type:"string" required:"yes" description:"Search phrase"`
}

var groupSearchCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupSearchFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "search",
		Short: "Search for group",
		Long:  `Get all groups that match your string in their name or path.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*groupSearchFlags)
		groups, _, err := gitlabClient.Groups.SearchGroup(*flags.Search)
		if err != nil {
			return err
		}
		return OutputJson(groups)
	},
}

func init() {
	groupLsCmd.Init()
	groupProjectsCmd.Init()
	groupGetCmd.Init()
	groupCreateCmd.Init()
	transferProjectCmd.Init()
	groupUpdateCmd.Init()
	groupDeleteCmd.Init()
	groupSearchCmd.Init()
	RootCmd.AddCommand(groupCmd)
}
