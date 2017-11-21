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

var groupMembersCmd = &cobra.Command{
	Use: "group-members",
	Short: "Access group members",
	Long: `Show members and access level of groups`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("check usage of `group-members` with `golab group-members -h`")
	},
}

var groupMembersLsCmd = &cobra.Command{
	Use: "ls",
	Short: "List all members of a group",
	Long: `Gets a list of groupmembers viewable by the authenticated user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id`not given - exiting")
		}
		opts := &gitlab.ListGroupMembersOptions{
			// TODO think about generic way of adding pagination...
			// ListOptions:
		}
		members, _, err := gitlabClient.Groups.ListGroupMembers(id, opts)
		if err != nil { return err }
		return OutputJson(members)
	},
}

var groupMemberGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get a member of a group",
	Long: `Get a member of a group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id`not given - exiting")
		}
		if userId == 0 {
			return errors.New("required parameter `-u` or `--user_id`not given - exiting")
		}
		member, _, err := gitlabClient.GroupMembers.GetGroupMember(id, userId)
		if err != nil {
			return err
		}
		return OutputJson(member)
	},
}

func init() {
	initGroupMembersLsCmd()
	initGroupMembersGetCmd()
	RootCmd.AddCommand(groupMembersCmd)
}

func initGroupMembersLsCmd() {
	groupMembersLsCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to show members for")
	groupMembersCmd.AddCommand(groupMembersLsCmd)
}

func initGroupMembersGetCmd() {
	groupMemberGetCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to get member from")
	groupMemberGetCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0,"(required) id of user to get group member infos")
	groupMembersCmd.AddCommand(groupMemberGetCmd)
}
