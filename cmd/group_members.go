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

var accessLevel, source, target int

var expiresAt string

var remove bool

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
			ListOptions: gitlab.ListOptions{Page: 1, PerPage: 1000},
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

var groupMemberAddCmd = &cobra.Command{
	Use: "add",
	Short: "Add a member to a group",
	Long: `Add a member to a group

  Access Levels:

	10 = Guest Permissions
	20 = Reporter Permissions
	30 = Developer Permissions
	40 = Master Permissions
	50 = Owner Permissions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id`not given - exiting")
		}
		if userId == 0 {
			return errors.New("required parameter `-u` or `--user_id`not given - exiting")
		}
		if accessLevel == 0 {
			return errors.New("required parameter `-a` or `--access_level` not given - exiting")
		}
		opts := &gitlab.AddGroupMemberOptions{
			UserID:      &userId,
			AccessLevel: int2AccessLevel(accessLevel),
		}
		if expiresAt != "" {
			opts.ExpiresAt = &expiresAt
		}
		member, _, err := gitlabClient.GroupMembers.AddGroupMember(id, opts)
		if err != nil { return err }
		return OutputJson(member)
	},
}

var groupMemberEditCmd = &cobra.Command{
	Use: "edit",
	Short: "Edit a member of a group or project",
	Long: `Updates a member of a group or project.

  Access Levels:

	10 = Guest Permissions
	20 = Reporter Permissions
	30 = Developer Permissions
	40 = Master Permissions
	50 = Owner Permissions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
		if userId == 0 {
			return errors.New("required parameter `-u` or `-user_id` not given - exiting")
		}
		if accessLevel == 0 {
			return errors.New("required parameter `-a` or `-access_level` not given - exiting")
		}
		opts := &gitlab.EditGroupMemberOptions{
			AccessLevel: int2AccessLevel(accessLevel),
		}
		if expiresAt != "" {
			opts.ExpiresAt = &expiresAt
		}
		member, _, err := gitlabClient.GroupMembers.EditGroupMember(id, userId, opts)
		if err != nil { return err }
		return OutputJson(member)
	},
}

var groupMemberDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Remove a member from a group or project",
	Long: `Removes a user from a group or project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
		if userId == 0 {
			return errors.New("required parameter `-u` or `--user_id` not given - exiting")
		}
		_, err := gitlabClient.GroupMembers.RemoveGroupMember(id, userId)
		return err
	},
}

var groupMemberSyncCmd = &cobra.Command{
	Use: "sync",
	Short: "Synchronizes members of 2 groups",
	Long: `Synchronizes the members of 2 groups, by either

* merging them (default) - members that exist in target group but not in source group are kept
* removing them (--remove) - members that exist in target group but not in source group are deleted`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if source == 0 {
			return errors.New("required parameter `--source` not given - exiting")
		}
		if target == 0 {
			return errors.New("required parameter `--target` not given - exiting")
		}

		opts := &gitlab.ListGroupMembersOptions{
			ListOptions: gitlab.ListOptions{Page: 1, PerPage: 1000},
		}

		createNonExistingTargetUsers(source, target, opts)

		if remove {
			err := removeTargetMembers(target, source, opts)
			if err != nil { return err }
		}

		members, _, err := gitlabClient.Groups.ListGroupMembers(target, opts)
		if err != nil { return err }
		return OutputJson(members)
	},
}

func createNonExistingTargetUsers(source int, target int, opts *gitlab.ListGroupMembersOptions) error {
	sourceMembers, _, err := gitlabClient.Groups.ListGroupMembers(source, opts)
	if err != nil { return err }
	for _, sourceMember := range sourceMembers  {
		_, resp, err := gitlabClient.GroupMembers.GetGroupMember(target, sourceMember.ID)
		if resp.StatusCode == 404 {   // 404 means "Not Found" --> does not exist yet
			newMemberOpts := &gitlab.AddGroupMemberOptions{
				UserID: &sourceMember.ID,
				AccessLevel: &sourceMember.AccessLevel,
			}
			if sourceMember.ExpiresAt != nil {
				expires, err := isoTime2String(sourceMember.ExpiresAt)
				if err != nil { return err }
				newMemberOpts.ExpiresAt = &expires
			}
			gitlabClient.GroupMembers.AddGroupMember(target, newMemberOpts)
		} else if err != nil {
			return err
		}
	}
	return nil
}

func removeTargetMembers(target int, source int, opts *gitlab.ListGroupMembersOptions) error {
	targetMembers, _, err := gitlabClient.Groups.ListGroupMembers(target, opts)
	if err != nil { return err }
	for _, targetMember := range targetMembers {
		_, resp, err := gitlabClient.GroupMembers.GetGroupMember(source, targetMember.ID)
		if resp.StatusCode == 404 {
			_, err := gitlabClient.GroupMembers.RemoveGroupMember(target, targetMember.ID)
			if err != nil { return err }
		} else if err != nil {
			return err
		}
	}
	return nil
}

func int2AccessLevel(accessLevel int) *gitlab.AccessLevelValue {
	switch accessLevel {
	case 10: return gitlab.AccessLevel(gitlab.GuestPermissions)
	case 20: return gitlab.AccessLevel(gitlab.ReporterPermissions)
	case 30: return gitlab.AccessLevel(gitlab.DeveloperPermissions)
	case 40: return gitlab.AccessLevel(gitlab.MasterPermissions)
	case 50: return gitlab.AccessLevel(gitlab.OwnerPermission)
	default: panic("Unrecognized value for AccessLevel")
	}
}

func init() {
	initGroupMembersLsCmd()
	initGroupMembersGetCmd()
	initGroupMemberAddCmd()
	initGroupMemberUpdateCmd()
	initGroupMemberDeleteCmd()
	initGroupMemberSyncCmd()
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

func initGroupMemberAddCmd() {
	groupMemberAddCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to add new member to")
	groupMemberAddCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(required) id of user to be added as new group member")
	groupMemberAddCmd.PersistentFlags().IntVarP(&accessLevel, "access_level", "a", 0, "(required) access level of new group member")
	groupMemberAddCmd.PersistentFlags().StringVarP(&expiresAt, "expires_at", "e", "", "(optional) expiry date of membership (yyyy-mm-dd)")
	groupMembersCmd.AddCommand(groupMemberAddCmd)
}

func initGroupMemberUpdateCmd() {
	groupMemberEditCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to change membership for")
	groupMemberEditCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(required) id the user to change membership for")
	groupMemberEditCmd.PersistentFlags().IntVarP(&accessLevel, "access_level", "a", 0, "(required) a valid access level")
	groupMemberEditCmd.PersistentFlags().StringVarP(&expiresAt, "expires_at", "e", "", "(optional) expiry date of membership (yyy-mm-dd)")
	groupMembersCmd.AddCommand(groupMemberEditCmd)
}

func initGroupMemberDeleteCmd() {
	groupMemberDeleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) the id of the group to delete user from")
	groupMemberDeleteCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(required) the id of the user to be removed from group")
	groupMembersCmd.AddCommand(groupMemberDeleteCmd)
}

func initGroupMemberSyncCmd() {
	groupMemberSyncCmd.PersistentFlags().IntVarP(&source, "source", "s", 0, "(required) id of group to copy members from")
	groupMemberSyncCmd.PersistentFlags().IntVarP(&target, "target", "t", 0, "(required) id of group to copy members to")
	groupMemberSyncCmd.PersistentFlags().BoolVarP(&remove, "remove", "r", false, "(optional) remove members in target group that don't exist in source group")
	groupMembersCmd.AddCommand(groupMemberSyncCmd)
}
