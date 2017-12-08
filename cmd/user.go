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
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

var key, title, user, email, password, skype, linkedin, twitter, websiteUrl, organization, username, externUid, provider, bio, location, adminString, canCreateGroupString, externalString, state, expires, scopes, name string
var id, userId, keyId, projectsLimit, tokenId, emailId int
var admin, canCreateGroup, skipConfirmation, external, active, blocked bool

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
	Opts:   nil,
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Single user.",
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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Allows creation of a new user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO fix binding of parameters
		if projectsLimit == -1 {
			projectsLimit = 10
		}
		createUserOptions := &gitlab.CreateUserOptions{
			Admin:            &admin,
			Bio:              &bio,
			CanCreateGroup:   &canCreateGroup,
			SkipConfirmation: &skipConfirmation,
			Email:            &email,
			Linkedin:         &linkedin,
			Name:             &name,
			Password:         &password,
			ProjectsLimit:    &projectsLimit,
			Skype:            &skype,
			Twitter:          &twitter,
			Username:         &username,
			WebsiteURL:       &websiteUrl,
		}
		if provider != "" {
			createUserOptions.Provider = &provider
			createUserOptions.ExternUID = &externUid
		}
		user, _, err := gitlabClient.Users.CreateUser(createUserOptions)
		if err != nil {
			return err
		}
		err = OutputJson(user)
		return err
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Long:  `Delete a user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := getUserId(id, user)
		if err != nil {
			return err
		}
		resp, err := gitlabClient.Users.DeleteUser(id)
		// TODO following the documentation, the user's data should be returned, but {} is returned...
		// TODO see https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/users.md#user-deletion
		return OutputJson(resp.Body)
	},
}

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify a user",
	Long: `Allows modification of a user's properties

Currently there are some restrictions:
* the email address cannot be modified
* the organization cannot be modified
* projects limit cannot be modified
* location cannot be modified
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := getUserId(id, user)
		if err != nil {
			return err
		}
		currUser, _, err := gitlabClient.Users.GetUser(id)
		if err != nil {
			return err
		}
		modifyUserOptions := &gitlab.ModifyUserOptions{}
		modifyUserOptions.Admin = boolFromParamAndCurrSetting(adminString, currUser.IsAdmin)
		// TODO changing email has no effect at the moment...
		if email != "" {
			modifyUserOptions.Email = &email
		}
		if username != "" {
			modifyUserOptions.Username = &username
		}
		if provider != "" {
			modifyUserOptions.Provider = &provider
			modifyUserOptions.ExternUID = &externUid
		}
		if name != "" {
			modifyUserOptions.Name = &name
		}
		if password != "" {
			modifyUserOptions.Password = &password
		}
		if skype != "" {
			modifyUserOptions.Skype = &skype
		}
		if twitter != "" {
			modifyUserOptions.Twitter = &twitter
		}
		if linkedin != "" {
			modifyUserOptions.Linkedin = &linkedin
		}
		if websiteUrl != "" {
			modifyUserOptions.WebsiteURL = &websiteUrl
		}
		if organization != "" {
			modifyUserOptions.Organization = &organization
		}
		if projectsLimit != -1 {
			modifyUserOptions.ProjectsLimit = &projectsLimit
		}
		if externUid != "" {
			modifyUserOptions.ExternUID = &externUid
		}
		if provider != "" {
			modifyUserOptions.Provider = &provider
		}
		if bio != "" {
			modifyUserOptions.Bio = &bio
		}
		if location != "" {
			modifyUserOptions.Location = &location
		}
		modifyUserOptions.CanCreateGroup = boolFromParamAndCurrSetting(canCreateGroupString, currUser.CanCreateGroup)
		if externalString == "true" {
			external = true
			modifyUserOptions.External = &external
		}
		if externalString == "false" {
			external = false
			modifyUserOptions.External = &external
		}

		user, _, err := gitlabClient.Users.ModifyUser(id, modifyUserOptions)
		if err != nil {
			return err
		}
		return OutputJson(user)
	},
}

var listSshKeysCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "Manage a user's ssh keys",
	Long:  `Allows management of a user's ssh keys (create, list, delete). If no sub-command is given, it lists ssh keys of currently authenticated user / user specified by user id.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id != 0 {
			sshKeys, _, err := gitlabClient.Users.ListSSHKeysForUser(id)
			if err != nil {
				return err
			}
			return OutputJson(sshKeys)
		} else {
			sshKeys, _, err := gitlabClient.Users.ListSSHKeys()
			if err != nil {
				return err
			}
			return OutputJson(sshKeys)
		}
	},
}

var getSshKeyCmd = &cobra.Command{
	Use:   "get",
	Short: "Single SSH key",
	Long:  `Get a single ssh key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if keyId != 0 {
			sshKey, _, err := gitlabClient.Users.GetSSHKey(keyId)
			if err != nil {
				return err
			}
			return OutputJson(sshKey)
		}
		return errors.New("you have to provide an id for a ssh key")
	},
}

var addSshKeyCmd = &cobra.Command{
	Use:   "add",
	Short: "Add SSH key",
	Long:  `Creates a new key (owned by the currently authenticated user, if no user id was given)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if key == "" || title == "" {
			return errors.New("you have to provide a key and a title")
		}
		sshKeyOps := &gitlab.AddSSHKeyOptions{
			Key:   &key,
			Title: &title,
		}
		if userId != 0 {
			sshKey, _, err := gitlabClient.Users.AddSSHKeyForUser(userId, sshKeyOps)
			if err != nil {
				return err
			}
			return OutputJson(sshKey)
		} else {
			sshKey, _, err := gitlabClient.Users.AddSSHKey(sshKeyOps)
			if err != nil {
				return err
			}
			return OutputJson(sshKey)
		}
	},
}

var deleteSshKeyCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete SSH key",
	Long:  `If no user id is given, deletes key owned by currently authenticated user. If a user id is given, deletes key owned by specified user. Available only for admins.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 {
			_, err := gitlabClient.Users.DeleteSSHKey(keyId)
			return err
		} else {
			_, err := gitlabClient.Users.DeleteSSHKeyForUser(userId, keyId)
			return err
		}
	},
}

var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "Get the last activity date for all users, sorted from oldest to newest.",
	Long: `The activities that update the timestamp are:

* Git HTTP/SSH activities (such as clone, push)
* User logging in into GitLab

By default, it shows the activity for all users in the last 6 months, but this can be amended by using the from parameter.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userActivities, _, err := gitlabClient.Users.GetUserActivities()
		if err != nil {
			return err
		}
		return OutputJson(userActivities)
	},
}

var impersinationTokenCmd = &cobra.Command{
	Use:   "impersonation-token",
	Short: "Manage impersonation tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("you cannot use this command without one of its sub-commands")
	},
}

var getImpersonationTokenCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all impersonation tokens of a user",
	Long:  `It retrieves every impersonation token of the user. Use the pagination parameters page and per_page to restrict the list of impersonation tokens.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 {
			return errors.New("required parameter `user` is missing")
		}
		if tokenId == 0 {
			opts := &gitlab.GetAllImpersonationTokensOptions{}
			if state != "" {
				opts.State = &state
			}
			token, _, err := gitlabClient.Users.GetAllImpersonationTokens(userId, opts)
			if err != nil {
				return err
			}
			return OutputJson(token)
		} else {
			tokens, _, err := gitlabClient.Users.GetImpersonationToken(userId, tokenId)
			if err != nil {
				return err
			}
			return OutputJson(tokens)
		}
	},
}

var createImpersonationTokenCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an impersonation token",
	Long:  `It creates a new impersonation token. Note that only administrators can do this. You are only able to create impersonation tokens to impersonate the user and perform both API calls and Git reads and writes. The user will not see these tokens in their profile settings page.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 {
			return errors.New("required parameter `user` is missing")
		}
		parsedScopes := strings.Split(scopes, ",")
		opts := &gitlab.CreateImpersonationTokenOptions{
			Name:   &name,
			Scopes: &parsedScopes,
		}
		if expires != "" {
			parsedExpires, err := time.Parse("2006-01-02", expires)
			if err != nil {
				return err
			}
			opts.ExpiresAt = &parsedExpires
		}
		token, _, err := gitlabClient.Users.CreateImpersonationToken(userId, opts)
		if err != nil {
			return err
		}
		return OutputJson(token)
	},
}

var revokeImpersonationTokenCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke an impersonation token",
	Long:  `It revokes an impersonation token.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 || tokenId == 0 {
			return errors.New("both, user_id and impersonation_token_id have to be given as parameters")
		}
		_, err := gitlabClient.Users.RevokeImpersonationToken(userId, tokenId)
		return err
	},
}

var emailsCmd = &cobra.Command{
	Use:   "emails",
	Short: "Manage emails for users",
	Long:  `List, add and delete emails for users`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("use one of the sub-commands, see `golab user emails -h`")
	},
}

var emailsListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List emails (for user)",
	Long: `If no user_id is given: get a list of currently authenticated user's emails.
If a user_id is given: Get a list of a specified user's emails. Available only for admin`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 {
			emails, _, err := gitlabClient.Users.ListEmails()
			if err != nil {
				return err
			}
			return OutputJson(emails)
		} else {
			emails, _, err := gitlabClient.Users.ListEmailsForUser(userId)
			if err != nil {
				return err
			}
			return OutputJson(emails)
		}
	},
}

var emailsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single email",
	Long:  `Get a single email for given email_id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		email, _, err := gitlabClient.Users.GetEmail(emailId)
		if err != nil {
			return err
		}
		return OutputJson(email)
	},
}

var emailsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add email (for user)",
	Long: `If no user_id is given: Creates a new email owned by the currently authenticated user.
If a user_id is given: Create new email owned by specified user. Available only for admin

Will return created email on success.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := &gitlab.AddEmailOptions{
			Email: &email,
		}
		if userId == 0 {
			resp, _, err := gitlabClient.Users.AddEmail(opts)
			if err != nil {
				return err
			}
			return OutputJson(resp)
		} else {
			resp, _, err := gitlabClient.Users.AddEmailForUser(userId, opts)
			if err != nil {
				return err
			}
			return OutputJson(resp)
		}
	},
}

var emailsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete email for current / given user",
	Long: `If no user_id is given: Deletes email owned by currently authenticated user.
If a user_id is given: Deletes email owned by a specified user. Available only for admin.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userId == 0 {
			_, err := gitlabClient.Users.DeleteEmail(emailId)
			return err
		} else {
			_, err := gitlabClient.Users.DeleteEmailForUser(userId, emailId)
			return err
		}
	},
}

func boolFromParamAndCurrSetting(paramString string, currentSetting bool) *bool {
	var result bool
	if paramString == "true" || paramString == "1" {
		result = true
	} else if paramString == "false" || paramString == "0" {
		result = false
	} else {
		result = currentSetting
	}
	return &result
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
	userLsCmd.Init()
	initUserCreateCommand()
	initUserModifyCommand()
	initUserDeleteCommand()
	initSshKeysCmd()
	initImpersonationTokenCmd()
	initEmailsCmd()
	userCmd.AddCommand(activitiesCmd)
	RootCmd.AddCommand(userCmd)
}

func initUserCreateCommand() {
	createCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "(mandatory) Email")
	createCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "(mandatory) Password")
	createCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "(mandatory) Username")
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(mandatory) Name")
	createCmd.PersistentFlags().StringVarP(&skype, "skype", "", "", "(optional) Skype ID")
	createCmd.PersistentFlags().StringVarP(&linkedin, "linkedin", "", "", "(optional) LinkedIn")
	createCmd.PersistentFlags().StringVarP(&twitter, "twitter", "", "", "(optional) Twitter account")
	createCmd.PersistentFlags().StringVarP(&websiteUrl, "website_url", "", "", "(optional) Website URL")
	createCmd.PersistentFlags().StringVarP(&organization, "organization", "", "", "(optional) Organization name")
	createCmd.PersistentFlags().IntVarP(&projectsLimit, "projects_limit", "", 10, "(optional) Number of projects user can create (10 is default)")
	createCmd.PersistentFlags().StringVarP(&externUid, "extern_uid", "", "", "(optional) External UID")
	createCmd.PersistentFlags().StringVarP(&provider, "provider", "", "", "(optional) External provider name")
	createCmd.PersistentFlags().StringVarP(&bio, "bio", "", "", "(optional) User's biography")
	createCmd.PersistentFlags().StringVarP(&location, "location", "", "", "(optional) User's location")
	createCmd.PersistentFlags().BoolVarP(&admin, "admin", "a", false, "(optional) User is admin - true or false (default)")
	createCmd.PersistentFlags().BoolVarP(&canCreateGroup, "can_create_group", "", false, "(optional) User can create groups - true or false (default)")
	createCmd.PersistentFlags().BoolVarP(&skipConfirmation, "skipConfirmation", "", false, "(optional) Skip confirmation")
	createCmd.PersistentFlags().BoolVarP(&external, "external", "", false, "(optional) Flags the user as external - true or false(default)")
	viper.BindPFlag("email", createCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("password", createCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("username", createCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("name", createCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("skype", createCmd.PersistentFlags().Lookup("skype"))
	viper.BindPFlag("linkedin", createCmd.PersistentFlags().Lookup("linkedin"))
	viper.BindPFlag("twitter", createCmd.PersistentFlags().Lookup("twitter"))
	viper.BindPFlag("website_url", createCmd.PersistentFlags().Lookup("website_url"))
	viper.BindPFlag("organization", createCmd.PersistentFlags().Lookup("organization"))
	viper.BindPFlag("projects_limit", createCmd.PersistentFlags().Lookup("projects_limit"))
	viper.BindPFlag("extern_uid", createCmd.PersistentFlags().Lookup("extern_uid"))
	viper.BindPFlag("provider", createCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("bio", createCmd.PersistentFlags().Lookup("bio"))
	viper.BindPFlag("location", createCmd.PersistentFlags().Lookup("location"))
	viper.BindPFlag("admin", createCmd.PersistentFlags().Lookup("admin"))
	viper.BindPFlag("can_create_group", createCmd.PersistentFlags().Lookup("can_create_group"))
	viper.BindPFlag("skipConfirmation", createCmd.PersistentFlags().Lookup("skipConfirmation"))
	viper.BindPFlag("external", createCmd.PersistentFlags().Lookup("external"))
	userCmd.AddCommand(createCmd)
}

func initUserModifyCommand() {
	modifyCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(mandatory) id of the user to be modified")
	modifyCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "(optional) user's new email address")
	modifyCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "(optional) user's new password")
	modifyCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "(optional) user's new username")
	modifyCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(optional) user's new name")
	modifyCmd.PersistentFlags().StringVarP(&skype, "skype", "", "", "(optional) user's new Skype ID")
	modifyCmd.PersistentFlags().StringVarP(&linkedin, "linkedin", "", "", "(optional) user's new LinkedIn account")
	modifyCmd.PersistentFlags().StringVarP(&twitter, "twitter", "", "", "(optional) user's new Twitter account")
	modifyCmd.PersistentFlags().StringVarP(&websiteUrl, "website_url", "", "", "(optional) user's new website URL")
	modifyCmd.PersistentFlags().StringVarP(&organization, "organization", "", "", "(optional) user's new organization name")
	modifyCmd.PersistentFlags().IntVarP(&projectsLimit, "projects_limit", "", -1, "(optional) user's new projects limit")
	modifyCmd.PersistentFlags().StringVarP(&externUid, "extern_uid", "", "", "(optional) user's new external UID")
	modifyCmd.PersistentFlags().StringVarP(&provider, "provider", "", "", "(optional) user's new external provider name")
	modifyCmd.PersistentFlags().StringVarP(&bio, "bio", "", "", "(optional) user's new biography")
	modifyCmd.PersistentFlags().StringVarP(&location, "location", "", "", "(optional) user's new location")
	modifyCmd.PersistentFlags().StringVarP(&adminString, "admin", "a", "", "(optional) user is admin - true or false")
	modifyCmd.PersistentFlags().StringVarP(&canCreateGroupString, "can_create_group", "", "", "(optional) user can create groups - true or false")
	modifyCmd.PersistentFlags().StringVarP(&externalString, "external", "", "", "(optional) flags the user as external - true or false")
	viper.BindPFlag("id", modifyCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("email", modifyCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("password", modifyCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("username", modifyCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("name", modifyCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("skype", modifyCmd.PersistentFlags().Lookup("skype"))
	viper.BindPFlag("linkedin", modifyCmd.PersistentFlags().Lookup("linkedin"))
	viper.BindPFlag("twitter", modifyCmd.PersistentFlags().Lookup("twitter"))
	viper.BindPFlag("website_url", modifyCmd.PersistentFlags().Lookup("website_url"))
	viper.BindPFlag("organization", modifyCmd.PersistentFlags().Lookup("organization"))
	viper.BindPFlag("projects_limit", modifyCmd.PersistentFlags().Lookup("projects_limit"))
	viper.BindPFlag("extern_uid", modifyCmd.PersistentFlags().Lookup("extern_uid"))
	viper.BindPFlag("provider", modifyCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("bio", modifyCmd.PersistentFlags().Lookup("bio"))
	viper.BindPFlag("location", modifyCmd.PersistentFlags().Lookup("location"))
	viper.BindPFlag("admin", modifyCmd.PersistentFlags().Lookup("admin"))
	viper.BindPFlag("can_create_group", modifyCmd.PersistentFlags().Lookup("can_create_group"))
	viper.BindPFlag("external", modifyCmd.PersistentFlags().Lookup("external"))
	userCmd.AddCommand(modifyCmd)
}

func initUserDeleteCommand() {
	deleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(mandatory if no username is set) id of the user to be deleted")
	deleteCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "(mandatory if no id is set) username of the user to be deleted")
	viper.BindPFlag("id", deleteCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("user", deleteCmd.PersistentFlags().Lookup("user"))
	userCmd.AddCommand(deleteCmd)
}

func initSshKeysCmd() {
	listSshKeysCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(optional) id of user to show ssh-keys for - if none is given, logged in user will be used")
	viper.BindPFlag("id", listSshKeysCmd.PersistentFlags().Lookup("id"))

	getSshKeyCmd.PersistentFlags().IntVarP(&keyId, "key_id", "k", 0, "(mandatory) key id of ssh key to be shown")
	viper.BindPFlag("key_id", getSshKeyCmd.PersistentFlags().Lookup("key_id"))

	addSshKeyCmd.PersistentFlags().IntVarP(&userId, "user", "u", 0, "(optional) id of user to add key for")
	addSshKeyCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "(mandatory) public ssh key")
	addSshKeyCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "(mandatory) title for ssh public key")
	viper.BindPFlag("user", getSshKeyCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("key", getSshKeyCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("title", getSshKeyCmd.PersistentFlags().Lookup("title"))

	deleteSshKeyCmd.PersistentFlags().IntVarP(&userId, "user", "u", 0, "(optional) id of user to delete key for")
	deleteSshKeyCmd.PersistentFlags().IntVarP(&keyId, "key_id", "k", 0, "(optional) id of ssh key to be deleted")
	viper.BindPFlag("user", deleteSshKeyCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("key_id", deleteSshKeyCmd.PersistentFlags().Lookup("key_id"))

	listSshKeysCmd.AddCommand(getSshKeyCmd, addSshKeyCmd, deleteSshKeyCmd)

	userCmd.AddCommand(listSshKeysCmd)
}

func initImpersonationTokenCmd() {
	getImpersonationTokenCmd.PersistentFlags().IntVarP(&userId, "user", "u", 0, "(required) id of user to get token(s) for")
	getImpersonationTokenCmd.PersistentFlags().IntVarP(&tokenId, "impersonation_token_id", "t", 0, "(optional) id of token")
	getImpersonationTokenCmd.PersistentFlags().StringVarP(&state, "state", "s", "", "(optional) state of token to be used as a filter (has no effect, if user is given)")
	viper.BindPFlag("user", getImpersonationTokenCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("impersonation_token_id", getImpersonationTokenCmd.PersistentFlags().Lookup("impersonation_token_id"))
	viper.BindPFlag("state", getImpersonationTokenCmd.PersistentFlags().Lookup("state"))

	createImpersonationTokenCmd.PersistentFlags().IntVarP(&userId, "user", "u", 0, "(required) the id of the user")
	createImpersonationTokenCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) the name of the impersonation token")
	createImpersonationTokenCmd.PersistentFlags().StringVarP(&expires, "expires_at", "e", "", "(optional) the expiration date of the impersonation token in ISO format (YYYY-MM-DD)")
	createImpersonationTokenCmd.PersistentFlags().StringVarP(&scopes, "scopes_array", "s", "", "(required) the comma-separated array of scopes of the impersonation token ( allowed values: `api`, `read_user`)")

	revokeImpersonationTokenCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(required) id of user to revoke token for")
	revokeImpersonationTokenCmd.PersistentFlags().IntVarP(&tokenId, "impersonation_token_id", "t", 0, "(required) id of token to be revoked")

	impersinationTokenCmd.AddCommand(getImpersonationTokenCmd, createImpersonationTokenCmd, revokeImpersonationTokenCmd)
	userCmd.AddCommand(impersinationTokenCmd)
}

func initEmailsCmd() {
	emailsListCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(optional) id of user to list emails for")
	viper.BindPFlag("user_id", emailsListCmd.PersistentFlags().Lookup("user_id"))

	emailsGetCmd.PersistentFlags().IntVarP(&emailId, "email_id", "i", 0, "(required) id of email")
	viper.BindPFlag("email_id", emailsGetCmd.PersistentFlags().Lookup("email_id"))

	emailsAddCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(optional) id of user to create email for")
	emailsAddCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "(required) email address to be created")
	viper.BindPFlag("user_id", emailsAddCmd.PersistentFlags().Lookup("user_id"))
	viper.BindPFlag("email", emailsAddCmd.PersistentFlags().Lookup("email"))

	emailsDeleteCmd.PersistentFlags().IntVarP(&userId, "user_id", "u", 0, "(optional) id of user to delete email from")
	emailsDeleteCmd.PersistentFlags().IntVarP(&emailId, "email_id", "i", 0, "(required) id of email to be deleted")
	viper.BindPFlag("user_id", emailsDeleteCmd.PersistentFlags().Lookup("user_id"))
	viper.BindPFlag("email_id", emailsDeleteCmd.PersistentFlags().Lookup("email_id"))

	emailsCmd.AddCommand(emailsListCmd, emailsAddCmd, emailsDeleteCmd)
	userCmd.AddCommand(emailsCmd)
}
