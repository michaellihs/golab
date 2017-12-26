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

// For detailed API specification, see https://docs.gitlab.com/ce/api/users.html
// TODO currently there is no support for GPG keys in the go-gitlab library

package cmd

import (
	"errors"
	"fmt"

	"strconv"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var email, password, username, state, expires, scopes, name string
var id, userId, tokenId, emailId int

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Gitlab users",
	Long:  `Allows create, update and deletion of a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of the subcommands, see `golab user -h`")
	},
}

// see https://docs.gitlab.com/ce/api/users.html#single-user
type userGetFlags struct {
	Id       *int    `flag_name:"id" short:"i" type:"int" required:"no" description:"The ID of a user"`
	Username *string `flag_name:"username" short:"u" type:"string" required:"no" description:"Username of a user"`
}

var userGetCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get a single user",
		Long:  `Get a single user. You can either provide --id or --username.`,
	},
	Run: func(cmd golabCommand) error {
		id, err := getUserId(getIdOrUsername(cmd.Flags.(*userGetFlags)))
		if err != nil {
			return err
		}
		user, _, err := gitlabClient.Users.GetUser(id)
		if err != nil {
			return err
		}
		return OutputJson(user)
	},
}

func getIdOrUsername(flags *userGetFlags) (int, string) {
	id := 0
	if flags.Id != nil {
		id = *flags.Id
	}
	username := ""
	if flags.Username != nil {
		username = *flags.Username
	}
	return id, username
}

// see https://docs.gitlab.com/ce/api/users.html#for-admins
type userGetAsAdminFlags struct {
	Username     *string `flag_name:"username" short:"u" type:"string" required:"no" description:"Username of the user to look up"`
	ExternalUID  *string `flag_name:"external_uid" type:"string" required:"no" description:"External UID of the user to look up (only together with provider)"`
	Provider     *string `flag_name:"provider" type:"string" required:"no" description:"External provider of user to look up"`
	External     *bool   `flag_name:"external" type:"bool" required:"no" description:"If set to true only external users will be returned"`
	CratedBefore *string `flag_name:"created_before" transform:"string2Time" type:"string" required:"no" description:"Search users created before, e.g. 2001-01-02"`
	CreatedAfter *string `flag_name:"created_after" transform:"string2Time" type:"string" required:"no" description:"Search users created after, e.g. 2001-01-02"`
}

var userGetAsAdminCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userGetAsAdminFlags{},
	Opts:   &gitlab.GetUsersAsAdminOptions{},
	Cmd: &cobra.Command{
		Use:   "get-as-admin",
		Short: "Lookup users by username",
		Long:  `Lookup users by username`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.GetUsersAsAdminOptions)
		users, _, err := gitlabClient.Users.GetUsersAsAdmin(opts)
		if err != nil {
			return err
		}
		return OutputJson(users)

	},
}

// see https://docs.gitlab.com/ce/api/users.html#list-users
type listUsersFlags struct {
	Active               *bool   `flag_name:"active" type:"bool" required:"no" description:"Filter users based on state active"`
	Blocked              *bool   `flag_name:"blocked" type:"bool" required:"no" description:"Filter users based on state blocked"`
	Search               *string `flag_name:"search" type:"string" required:"no" description:"Search for users by email or username (admin only)"`
	Username             *string `flag_name:"username" type:"string" required:"no" description:"Lookup users by username (admin only)"`
	ExternUid            *string `flag_name:"extern_uid" type:"string" required:"no" description:"Lookup users by external UID and provider (admin only)"`
	Provider             *string `flag_name:"provider" type:"string" required:"no" description:"Lookup users by external UID and provider (admin only)"`
	External             *bool   `flag_name:"external" type:"bool" required:"no" description:"Search for users who are external (admin only)"`
	CreatedBefore        *string `flag_name:"created_before" type:"string" required:"no" description:"Search users by creation date time range, e.g. 2001-01-02T00:00:00.060Z (admin only)"`
	CreatedAfter         *string `flag_name:"created_after" type:"string" required:"no" description:"Search users by creation date time range, e.g. 2001-01-02T00:00:00.060Z (admin only)"`
	CustomAttributeKey   *string `flag_name:"custom_attribute_key" type:"string" required:"no" description:"Filter by custom attribute key (admin only)"`
	CustomAttributeValue *string `flag_name:"custom_attribute_value" type:"string" required:"no" description:"Filter by custom attribute value (admin only)"`
}

var userLsCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &listUsersFlags{},
	Opts:   &gitlab.ListUsersOptions{},
	Paged:  true,
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List users",
		Long:  `Get a list of users.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.ListUsersOptions)
		users, _, err := gitlabClient.Users.ListUsers(opts)
		if err != nil {
			return err
		}
		return OutputJson(users)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#user-creation
type userCreateFlags struct {
	Email            *string `flag_name:"email" short:"e" type:"string" required:"yes" description:"Email"`
	Password         *string `flag_name:"password" short:"p" type:"string" required:"no" description:"Password"`
	ResetPassword    *bool   `flag_name:"reset_password" type:"bool" required:"no" description:"Send user password reset link - true or false(default)"`
	Username         *string `flag_name:"username" short:"u" type:"string" required:"yes" description:"Username"`
	Name             *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"Name"`
	Skype            *string `flag_name:"skype" type:"string" required:"no" description:"Skype ID"`
	Linkedin         *string `flag_name:"linkedin" type:"string" required:"no" description:"LinkedIn"`
	Twitter          *string `flag_name:"twitter" type:"string" required:"no" description:"Twitter account"`
	WebsiteUrl       *string `flag_name:"website_url" type:"string" required:"no" description:"Website URL"`
	Organization     *string `flag_name:"organization" type:"string" required:"no" description:"Organization name"`
	ProjectsLimit    *int    `flag_name:"projects_limit" type:"int" required:"no" description:"Number of projects user can create"`
	ExternUid        *string `flag_name:"extern_uid" type:"string" required:"no" description:"External UID"`
	Provider         *string `flag_name:"provider" type:"string" required:"no" description:"External provider name"`
	Bio              *string `flag_name:"bio" type:"string" required:"no" description:"User's biography"`
	Location         *string `flag_name:"location" type:"string" required:"no" description:"User's location"`
	Admin            *bool   `flag_name:"admin" type:"bool" required:"no" description:"User is admin - true or false (default)"`
	CanCreateGroup   *bool   `flag_name:"can_create_group" type:"bool" required:"no" description:"User can create groups - true or false"`
	SkipConfirmation *bool   `flag_name:"skip_confirmation" type:"bool" required:"no" description:"Skip confirmation - true or false (default)"`
	External         *bool   `flag_name:"external" type:"bool" required:"no" description:"Flags the user as external - true or false(default)"`
	// TODO currently not supported by go-gitlab
	//Avatar           *string `flag_name:"avatar" type:"string" required:"no" description:"Image file for user's avatar"`
}

var userCreateCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userCreateFlags{},
	Opts:   &gitlab.CreateUserOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "User creation",
		Long:  `Creates a new user. Note only administrators can create new users. Either password or reset_password should be specified (reset_password takes priority).`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.CreateUserOptions)
		OutputJson(opts)
		user, _, err := gitlabClient.Users.CreateUser(opts)
		if err != nil {
			return err
		}
		return OutputJson(user)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#user-deletion
type userDeleteFlags struct {
	Id         *string `flag_name:"id" short:"i" type:"int" required:"yes" description:"User ID or user name of user to be deleted"`
	HardDelete *bool   `flag_name:"hard_delete" short:"d" type:"bool" required:"no" description:"If true, contributions that would usually be moved to the ghost user will be deleted instead, as well as groups owned solely by this user."`
}

var userDeleteCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userDeleteFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "User deletion",
		Long:  `Deletes a user. Available only for administrators. This returns a 204 No Content status code if the operation was successfully or 404 if the resource was not found.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userDeleteFlags)
		id, err := userIdFromFlag(*flags.Id)
		if err != nil {
			return err
		}
		_, err = gitlabClient.Users.DeleteUser(id)
		return err
	},
}

func userIdFromFlag(intStrId string) (int, error) {
	idInt, err := strconv.Atoi(intStrId)
	username := intStrId
	if err != nil {
		idInt = 0
	} else {
		username = ""
	}
	return getUserId(idInt, username)
}

// see https://docs.gitlab.com/ce/api/users.html#user-modification
type userModifyFlags struct {
	Id               *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"User ID or user name of user to be deleted"`
	Email            *string `flag_name:"email" short:"e" type:"string" required:"no" description:"Email"`
	Password         *string `flag_name:"password" short:"p" type:"string" required:"no" description:"Password"`
	Username         *string `flag_name:"username" short:"u" type:"string" required:"no" description:"Username"`
	Name             *string `flag_name:"name" short:"n" type:"string" required:"no" description:"Name"`
	Skype            *string `flag_name:"skype" type:"string" required:"no" description:"Skype ID"`
	Linkedin         *string `flag_name:"linkedin" type:"string" required:"no" description:"LinkedIn"`
	Twitter          *string `flag_name:"twitter" type:"string" required:"no" description:"Twitter account"`
	WebsiteUrl       *string `flag_name:"website_url" type:"string" required:"no" description:"Website URL"`
	Organization     *string `flag_name:"organization" type:"string" required:"no" description:"Organization name"`
	ProjectsLimit    *int    `flag_name:"projects_limit" type:"int" required:"no" description:"Number of projects user can create"`
	ExternUid        *string `flag_name:"extern_uid" type:"string" required:"no" description:"External UID"`
	Provider         *string `flag_name:"provider" type:"string" required:"no" description:"External provider name"`
	Bio              *string `flag_name:"bio" type:"string" required:"no" description:"User's biography"`
	Location         *string `flag_name:"location" type:"string" required:"no" description:"User's location"`
	Admin            *bool   `flag_name:"admin" type:"bool" required:"no" description:"User is admin - true or false (default)"`
	CanCreateGroup   *bool   `flag_name:"can_create_group" type:"bool" required:"no" description:"User can create groups - true or false"`
	SkipConfirmation *bool   `flag_name:"skip_confirmation" type:"bool" required:"no" description:"Skip confirmation - true or false (default)"`
	External         *bool   `flag_name:"external" type:"bool" required:"no" description:"Flags the user as external - true or false(default)"`
	// TODO currently not supported by go-gitlab
	//Avatar           *string `flag_name:"avatar" type:"string" required:"no" description:"Image file for user's avatar"`
}

var userModifyCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userModifyFlags{},
	Opts:   &gitlab.ModifyUserOptions{},
	Cmd: &cobra.Command{
		Use:   "modify",
		Short: "User modification",
		Long:  `Modifies an existing user. Only administrators can change attributes of a user.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userModifyFlags)
		opts := cmd.Opts.(*gitlab.ModifyUserOptions)
		id, err := userIdFromFlag(*flags.Id)
		if err != nil {
			return err
		}
		user, _, err := gitlabClient.Users.ModifyUser(id, opts)
		if err != nil {
			return err
		}
		return OutputJson(user)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#list-ssh-keys
var userSshKeysCmd = &golabCommand{
	Parent: userCmd,
	Flags:  nil,
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "ssh-keys",
		Short: "Manage a user's SSH keys",
		Long:  `SSH key management for users`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("command cannot be run without sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/users.html#list-ssh-keys
type userSshKeysListFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"no" description:"id of user to show SSH keys for - if none is given, logged in user will be used"`
}

var userSshKeysListCmd = &golabCommand{
	Parent: userSshKeysCmd.Cmd,
	Flags:  &userSshKeysListFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List SSH keys",
		Long:  `Get a list of (currently authenticated user's) SSH keys.`,
	},
	Run: func(cmd golabCommand) error {
		var sshKeys []*gitlab.SSHKey
		var err error
		flags := cmd.Flags.(*userSshKeysListFlags)
		if flags.Id != nil {
			id, err := userIdFromFlag(*flags.Id)
			if err != nil {
				return err
			}
			sshKeys, _, err = gitlabClient.Users.ListSSHKeysForUser(id)
		} else {
			sshKeys, _, err = gitlabClient.Users.ListSSHKeys()
		}
		if err != nil {
			return err
		}
		return OutputJson(sshKeys)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#single-ssh-key
type userSshKeysGetFlags struct {
	Id *int `flag_name:"id" short:"i" required:"yes" description:"key id of SSH key to be shown"`
}

var userSshKeysGetCmd = &golabCommand{
	Parent: userSshKeysCmd.Cmd,
	Flags:  &userSshKeysGetFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Single SSH key",
		Long:  `Get a single SSH key for a given user.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userSshKeysGetFlags)
		sshKey, _, err := gitlabClient.Users.GetSSHKey(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(sshKey)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#add-ssh-key
type userSshKeysAddFlags struct {
	User  *string `flag_name:"user" short:"u" type:"string" required:"yes" description:"User ID or user name of user to delete SSH key from"`
	Title *string `flag_name:"title" short:"t" type:"string" required:"yes" description:"New SSH Key's title"`
	Key   *string `flag_name:"key" short:"k" type:"string" required:"yes" description:"Public SSH key"`
}

var userSshKeysAddCmd = &golabCommand{
	Parent: userSshKeysCmd.Cmd,
	Flags:  &userSshKeysAddFlags{},
	Opts:   &gitlab.AddSSHKeyOptions{},
	Cmd: &cobra.Command{
		Use:   "add",
		Short: "Add SSH key",
		Long:  `Creates a new key (owned by the currently authenticated user, if no user id was given)`,
	},
	Run: func(cmd golabCommand) error {
		var err error
		var key *gitlab.SSHKey
		flags := cmd.Flags.(*userSshKeysAddFlags)
		opts := cmd.Opts.(*gitlab.AddSSHKeyOptions)
		if flags.User == nil {
			key, _, err = gitlabClient.Users.AddSSHKey(opts)
		} else {
			userId, err := userIdFromFlag(*flags.User)
			if err != nil {
				return err
			}
			key, _, err = gitlabClient.Users.AddSSHKeyForUser(userId, opts)
		}
		if err != nil {
			return err
		}
		return OutputJson(key)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#delete-ssh-key-for-current-user
type userSshKeysDeleteFlags struct {
	KeyId *int    `flag_name:"key_id" short:"k" required:"yes" description:"key id of SSH key to be deleted"`
	User  *string `flag_name:"user" short:"u" type:"string" required:"yes" description:"User ID or user name of user to delete SSH key from"`
}

var userSshKeysDeleteCmd = &golabCommand{
	Parent: userSshKeysCmd.Cmd,
	Flags:  &userSshKeysDeleteFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete SSH key",
		Long:  `Deletes key owned by a specified user (available only for admin) or by currently logged in user.`,
	},
	Run: func(cmd golabCommand) error {
		var err error
		flags := cmd.Flags.(*userSshKeysDeleteFlags)
		if flags.User == nil {
			_, err = gitlabClient.Users.DeleteSSHKey(*flags.KeyId)
		} else {
			userId, err := userIdFromFlag(*flags.User)
			if err != nil {
				return err
			}
			_, err = gitlabClient.Users.DeleteSSHKeyForUser(userId, *flags.KeyId)
		}
		return err
	},
}

// see https://docs.gitlab.com/ce/api/users.html#get-user-activities-admin-only
type userActivitiesFlags struct {
	From *string `flag_name:"from" transform:"string2IsoTime" type:"string" required:"no" description:"Date string in the format YEAR-MONTH-DAY, e.g. 2016-03-11. Defaults to 6 months ago."`
}

var userActivitiesCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userActivitiesFlags{},
	Opts:   &gitlab.GetUserActivitiesOptions{},
	Cmd: &cobra.Command{
		Use:   "activities",
		Short: "Get user activities (admin only)",
		Long: `Get the last activity date for all users, sorted from oldest to newest.

The activities that update the timestamp are:

* Git HTTP/SSH activities (such as clone, push)
* User logging in into GitLab

By default, it shows the activity for all users in the last 6 months, but this can be amended by using the from parameter.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.GetUserActivitiesOptions)
		userActivities, _, err := gitlabClient.Users.GetUserActivities(opts)
		if err != nil {
			return err
		}
		return OutputJson(userActivities)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#get-all-impersonation-tokens-of-a-user
var userImpersonationTokenCmd = &golabCommand{
	Parent: userCmd,
	Flags:  nil,
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "impersonation-token",
		Short: "Impersonation token",
		Long:  `Manage a user's impersonation token`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("you cannot use this command without one of its sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/users.html#get-all-impersonation-tokens-of-a-user
type userImpersonationTokenGetAllFlags struct {
	UserId *string `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"The ID of the user or the name of the user to get tokens for"`
	State  *string `flag_name:"state" short:"s" type:"string" required:"no" description:"filter tokens based on state (all, active, inactive)"`
}

var userImpersonationTokenGetAllCmd = &golabCommand{
	Parent: userImpersonationTokenCmd.Cmd,
	Flags:  &userImpersonationTokenGetAllFlags{},
	Opts:   &gitlab.GetAllImpersonationTokensOptions{},
	Cmd: &cobra.Command{
		Use:   "get-all",
		Short: "Get all impersonation tokens of a user",
		Long:  `It retrieves every impersonation token of the user.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userImpersonationTokenGetAllFlags)
		opts := cmd.Opts.(*gitlab.GetAllImpersonationTokensOptions)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		tokens, _, err := gitlabClient.Users.GetAllImpersonationTokens(userId, opts)
		if err != nil {
			return err
		}
		return OutputJson(tokens)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#get-an-impersonation-token-of-a-user
type userImpersonationTokenGetFlags struct {
	UserId               *string `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"The ID of the user or the username for which to get a token"`
	ImpersonationTokenId *int    `flag_name:"impersonation_token_id" short:"t" type:"integer" required:"yes" description:"The ID of the impersonation token"`
}

var userImpersonationTokenGetCmd = &golabCommand{
	Parent: userImpersonationTokenCmd.Cmd,
	Flags:  &userImpersonationTokenGetFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get an impersonation token of a user",
		Long:  `It shows a user's impersonation token (admins only).`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userImpersonationTokenGetFlags)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		token, _, err := gitlabClient.Users.GetImpersonationToken(userId, *flags.ImpersonationTokenId)
		if err != nil {
			return err
		}
		return OutputJson(token)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#create-an-impersonation-token
type userImpersonationTokenCreateFlags struct {
	UserId    *string   `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"The ID of the user"`
	Name      *string   `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the impersonation token"`
	ExpiresAt *string   `flag_name:"expires_at" short:"e" type:"string" transform:"string2Time" required:"no" description:"The expiration date of the impersonation token in ISO format (YYYY-MM-DD)"`
	Scopes    *[]string `flag_name:"scopes" short:"s" type:"array" required:"yes" description:"The array of scopes of the impersonation token (api, read_user)"`
}

var userImpersonationTokenCreateCmd = &golabCommand{
	Parent: userImpersonationTokenCmd.Cmd,
	Flags:  &userImpersonationTokenCreateFlags{},
	Opts:   &gitlab.CreateImpersonationTokenOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create an impersonation token (admin only)",
		Long:  `It creates a new impersonation token. Note that only administrators can do this. You are only able to create impersonation tokens to impersonate the user and perform both API calls and Git reads and writes. The user will not see these tokens in their profile settings page. Requires admin permissions.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userImpersonationTokenCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateImpersonationTokenOptions)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		token, _, err := gitlabClient.Users.CreateImpersonationToken(userId, opts)
		if err != nil {
			return err
		}
		return OutputJson(token)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#revoke-an-impersonation-token
type userImpersonationTokenRevokeFlags struct {
	UserId               *string `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"The ID of the user or username of user to revoke token for"`
	ImpersonationTokenId *int    `flag_name:"impersonation_token_id" short:"t" type:"integer" required:"yes" description:"The ID of the impersonation token"`
}

var userImpersonationTokenRevokeCmd = &golabCommand{
	Parent: userImpersonationTokenCmd.Cmd,
	Flags:  &userImpersonationTokenRevokeFlags{},
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "revoke",
		Short: "Revoke an impersonation token (admin only)",
		Long:  `It revokes an impersonation token. Requires admin permissions.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userImpersonationTokenRevokeFlags)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		_, err = gitlabClient.Users.RevokeImpersonationToken(userId, *flags.ImpersonationTokenId)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/users.html#list-emails
var userEmailsCmd = &golabCommand{
	Parent: userCmd,
	Cmd: &cobra.Command{
		Use:   "emails",
		Short: "User emails",
		Long:  `Manage user's email'`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot run this command without any further sub-command")
	},
}

// see https://docs.gitlab.com/ce/api/users.html#list-emails
type userEmailsListFlags struct {
	UserId *string `flag_name:"user_id" short:"u" type:"string" required:"no" description:"The ID of the user or username of user to list emails for. If none is given, emails of currently logged in user are shown"`
}

var userEmailsListCmd = &golabCommand{
	Parent: userEmailsCmd.Cmd,
	Flags:  &userEmailsListFlags{},
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List emails",
		Long: `If no user_id is given: get a list of currently authenticated user's emails.
If a user_id is given: Get a list of a specified user's emails. Available only for admin`,
	},
	Run: func(cmd golabCommand) error {
		var err error
		var emails []*gitlab.Email
		flags := cmd.Flags.(*userEmailsListFlags)
		if flags.UserId == nil {
			emails, _, err = gitlabClient.Users.ListEmails()
		} else {
			userId, err := userIdFromFlag(*flags.UserId)
			if err != nil {
				return err
			}
			emails, _, err = gitlabClient.Users.ListEmailsForUser(userId)
		}
		if err != nil {
			return err
		}
		return OutputJson(emails)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#single-email
type userEmailsGetFlags struct {
	EmailId *int `flag_name:"email_id" short:"e" type:"int" required:"yes" description:"email ID"`
}

var userEmailsGetCmd = &golabCommand{
	Parent: userEmailsCmd.Cmd,
	Flags:  &userEmailsGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Single email",
		Long:  `Get a single email.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userEmailsGetFlags)
		email, _, err := gitlabClient.Users.GetEmail(*flags.EmailId)
		if err != nil {
			return err
		}
		return OutputJson(email)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#add-email
type userEmailsAddFlags struct {
	UserId *string `flag_name:"user_id" short:"u" type:"string" required:"no" description:"id or username of user to add email to"`
	Email  *string `flag_name:"email" short:"e" type:"string" required:"yes" description:"email address"`
}

var userEmailsAddCmd = &golabCommand{
	Parent: userEmailsCmd.Cmd,
	Flags:  &userEmailsAddFlags{},
	Opts:   &gitlab.AddEmailOptions{},
	Cmd: &cobra.Command{
		Use:   "add",
		Short: "Create new email (for user)",
		Long: `If no user_id is given: Creates a new email owned by the currently authenticated user.
If a user_id is given: Create new email owned by specified user. Available only for admin

Will return created email on success.`,
	},
	Run: func(cmd golabCommand) error {
		var email *gitlab.Email
		var err error
		flags := cmd.Flags.(*userEmailsAddFlags)
		opts := cmd.Opts.(*gitlab.AddEmailOptions)
		if flags.UserId == nil {
			email, _, err = gitlabClient.Users.AddEmail(opts)
		} else {
			userId, err := userIdFromFlag(*flags.UserId)
			if err != nil {
				return err
			}
			email, _, err = gitlabClient.Users.AddEmailForUser(userId, opts)
		}
		if err != nil {
			return err
		}
		return OutputJson(email)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#delete-email-for-current-user
type userEmailsDeleteFlags struct {
	UserId  *string `flag_name:"user_id" short:"u" type:"string" required:"no" description:"id or username of user to delete email from"`
	EmailId *int    `flag_name:"email_id" short:"e" type:"string" required:"yes" description:"id of email to be deleted"`
}

var userEmailsDeleteCmd = &golabCommand{
	Parent: userEmailsCmd.Cmd,
	Flags:  &userEmailsDeleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete email for (current) user",
		Long: `If no user_id is given: Deletes email owned by currently authenticated user.
If a user_id is given: Deletes email owned by a specified user. Available only for admin.`,
	},
	Run: func(cmd golabCommand) error {
		var err error
		flags := cmd.Flags.(*userEmailsDeleteFlags)
		if flags.UserId == nil {
			_, err = gitlabClient.Users.DeleteEmail(*flags.EmailId)
		} else {
			userId, err := userIdFromFlag(*flags.UserId)
			if err != nil {
				return err
			}
			_, err = gitlabClient.Users.DeleteEmailForUser(userId, *flags.EmailId)
		}
		return err
	},
}

// see https://docs.gitlab.com/ce/api/users.html#block-user
type userBlockFlags struct {
	UserId *string `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"id or username of user to block"`
}

var userBlockCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userBlockFlags{},
	Cmd: &cobra.Command{
		Use:   "block",
		Short: "Block user",
		Long:  `Blocks the specified user. Available only for admin.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userBlockFlags)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		return gitlabClient.Users.BlockUser(userId)
	},
}

// see https://docs.gitlab.com/ce/api/users.html#unblock-user
type userUnblockFlags struct {
	UserId *string `flag_name:"user_id" short:"u" type:"string" required:"yes" description:"id or username of user to unblock"`
}

var userUnblockCmd = &golabCommand{
	Parent: userCmd,
	Flags:  &userUnblockFlags{},
	Cmd: &cobra.Command{
		Use:   "unblock",
		Short: "Unblock user",
		Long:  `Unblocks the specified user. Available only for admin`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*userUnblockFlags)
		userId, err := userIdFromFlag(*flags.UserId)
		if err != nil {
			return err
		}
		return gitlabClient.Users.UnblockUser(userId)
	},
}

func getUserId(id int, username string) (int, error) {
	if (id == 0 && username == "") || (id != 0 && username != "") {
		return 0, errors.New("you either have to provide an id or a username")
	}
	if username != "" {
		users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{Username: &username})
		if err != nil {
			return 0, err
		}
		if len(users) != 1 {
			return 0, errors.New("Number of users found for username: " + strconv.Itoa(len(users)))
		}
		id = users[0].ID
	}
	return id, nil
}

func init() {
	userGetCmd.Init()
	userGetAsAdminCmd.Init()
	userLsCmd.Init()
	userCreateCmd.Init()
	userModifyCmd.Init()
	userDeleteCmd.Init()
	userSshKeysCmd.Init()
	userSshKeysListCmd.Init()
	userSshKeysGetCmd.Init()
	userSshKeysAddCmd.Init()
	userSshKeysDeleteCmd.Init()
	userActivitiesCmd.Init()
	userImpersonationTokenCmd.Init()
	userImpersonationTokenGetAllCmd.Init()
	userImpersonationTokenGetCmd.Init()
	userImpersonationTokenCreateCmd.Init()
	userImpersonationTokenRevokeCmd.Init()
	userEmailsCmd.Init()
	userEmailsListCmd.Init()
	userEmailsGetCmd.Init()
	userEmailsAddCmd.Init()
	userEmailsDeleteCmd.Init()
	userBlockCmd.Init()
	userUnblockCmd.Init()
	RootCmd.AddCommand(userCmd)
}
