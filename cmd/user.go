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
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/spf13/viper"
	"strconv"
)

var user, email, password, skype, linkedin, twitter, websiteUrl, organization, username, externUid, provider, bio, location, adminString, canCreateGroupString, externalString string
var keyId, projectsLimit int
var admin, canCreateGroup, skipConfirmation, external, active, blocked bool

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Gitlab users",
	Long: `Allows create, update and deletion of a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of the subcommands, see `golab user -h`")
	},
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get user details",
	Long: `Get detailed information for given user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := getUserId(id, username)
		user, _, err := gitlabClient.Users.GetUser(id)
		if err != nil {	return err }
		return OutputJson(user)
	},
}

var lsCmd = &cobra.Command{
	Use: "ls",
	Short: "Get list of all users",
	Long: `Get a list of all users on the Gitlab server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listUserOptions := &gitlab.ListUsersOptions{}
		if active {
			listUserOptions.Active = &active
		}
		if blocked {
			listUserOptions.Blocked = &blocked
		}
		users, _, err := gitlabClient.Users.ListUsers(listUserOptions)
		if err != nil {	return err }
		return OutputJson(users)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long: `Allows creation of a new user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO fix binding of parameters
		if projectsLimit == -1 { projectsLimit = 10 }
		createUserOptions := &gitlab.CreateUserOptions{
			Admin: &admin,
			Bio: &bio,
			CanCreateGroup: &canCreateGroup,
			SkipConfirmation: &skipConfirmation,
			Email: &email,
			Linkedin: &linkedin,
			Name: &name,
			Password: &password,
			ProjectsLimit: &projectsLimit,
			Skype: &skype,
			Twitter: &twitter,
			Username: &username,
			WebsiteURL: &websiteUrl,
		}
		if provider != "" {
			createUserOptions.Provider = &provider
			createUserOptions.ExternUID = &externUid
		}
		user, _, err := gitlabClient.Users.CreateUser(createUserOptions)
		if err != nil {	return err }
		err = OutputJson(user)
		return err
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Long: `Delete a user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := getUserId(id, user)
		if err != nil {	return err }
		resp , err := gitlabClient.Users.DeleteUser(id)
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
		if err != nil {	return err }
		currUser, _, err := gitlabClient.Users.GetUser(id)
		if err != nil {	return err }
		modifyUserOptions := &gitlab.ModifyUserOptions{}
		modifyUserOptions.Admin = boolFromParamAndCurrSetting(adminString, currUser.IsAdmin)
		// TODO changing email has no effect at the moment...
		if email != "" { modifyUserOptions.Email = &email }
		if username != "" { modifyUserOptions.Username = &username }
		if provider != "" {
			modifyUserOptions.Provider = &provider
			modifyUserOptions.ExternUID = &externUid
		}
		if name != "" { modifyUserOptions.Name = &name}
		if password != "" { modifyUserOptions.Password = &password }
		if skype != "" { modifyUserOptions.Skype = &skype }
		if twitter != "" { modifyUserOptions.Twitter = &twitter }
		if linkedin != "" { modifyUserOptions.Linkedin = &linkedin }
		if websiteUrl != "" { modifyUserOptions.WebsiteURL = &websiteUrl }
		// TODO currently not supported by go-gitlab API
		//if organization != "" { modifyUserOptions.Organization = &organization }
		// TODO currently not supported by go-gitlab API
		//if projectsLimit != -1 { modifyUserOptions.projectsLimit = &projectsLimit }
		if externUid != "" { modifyUserOptions.ExternUID = &externUid }
		if provider != "" { modifyUserOptions.Provider = &provider }
		if bio != "" { modifyUserOptions.Bio = &bio }
		// TODO currently not supported by go-gitlab API
		//if location != "" { modifyUserOptions.Location = location }
		modifyUserOptions.CanCreateGroup = boolFromParamAndCurrSetting(canCreateGroupString, currUser.CanCreateGroup)
		// TODO currently not supported by go-gitlab API
		//if external != "" { modifyUserOptions.External = &external }

		user, _, err := gitlabClient.Users.ModifyUser(id, modifyUserOptions)
		if err != nil {	return err }
		return OutputJson(user)
	},
}

var listSshKeysCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "Manage a user's ssh keys",
	Long:  `Allows management of a user's ssh keys (create, list, delete)'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id != 0 {
			sshKeys, _, err := gitlabClient.Users.ListSSHKeysForUser(id)
			if err != nil { return err }
			return OutputJson(sshKeys)
		} else {
			sshKeys, _, err := gitlabClient.Users.ListSSHKeys()
			if err != nil { return err }
			return OutputJson(sshKeys)
		}
	},
}

var getSshKeyCmd = &cobra.Command{
	Use: "get",
	Short: "Single SSH key",
	Long: `Get a single ssh key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if keyId != 0 {
			sshKey, _, err := gitlabClient.Users.GetSSHKey(keyId)
			if err != nil { return err }
			return OutputJson(sshKey)
		}
		return errors.New("you have to provide an id for a ssh key")
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
	initUserGetCommand()
	initUserLsCommand()
	initUserCreateCommand()
	initUserModifyCommand()
	initUserDeleteCommand()
	initListSshKeysCmd()
	RootCmd.AddCommand(userCmd)
}

func initUserGetCommand() {
	getCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "(mandatory if id is unset) username of the user to look up")
	getCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(mandatory if username is unset) id of the user to look up")
	viper.BindPFlag("username", getCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("id", getCmd.PersistentFlags().Lookup("id"))
	userCmd.AddCommand(getCmd)
}

func initUserLsCommand() {
	lsCmd.PersistentFlags().BoolVarP(&active, "active", "a", false, "(optional) show only active users")
	lsCmd.PersistentFlags().BoolVarP(&blocked,"blocked", "b", false, "(optional) show only blocked users")
	userCmd.AddCommand(lsCmd)
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
	// TODO currently not supported by go-gitlab API
	// modifyCmd.PersistentFlags().StringVarP(&organization, "organization", "", "", "(optional) user's new organization name")
	modifyCmd.PersistentFlags().IntVarP(&projectsLimit, "projects_limit", "", -1, "(optional) user's new projects limit")
	modifyCmd.PersistentFlags().StringVarP(&externUid, "extern_uid", "", "", "(optional) user's new external UID")
	modifyCmd.PersistentFlags().StringVarP(&provider, "provider", "", "", "(optional) user's new external provider name")
	modifyCmd.PersistentFlags().StringVarP(&bio, "bio", "", "", "(optional) user's new biography")
	modifyCmd.PersistentFlags().StringVarP(&location, "location", "", "", "(optional) user's new location")
	modifyCmd.PersistentFlags().StringVarP(&adminString, "admin", "a", "", "(optional) user is admin - true or false")
	modifyCmd.PersistentFlags().StringVarP(&canCreateGroupString, "can_create_group", "", "", "(optional) user can create groups - true or false")
	// TODO currently not supported by go-gitlab API
	// modifyCmd.PersistentFlags().StringVarP(&externalString, "external", "", "", "(optional) flags the user as external - true or false")
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
	// TODO currently not supported by go-gitlab API
	// viper.BindPFlag("projects_limit", modifyCmd.PersistentFlags().Lookup("projects_limit"))
	viper.BindPFlag("extern_uid", modifyCmd.PersistentFlags().Lookup("extern_uid"))
	viper.BindPFlag("provider", modifyCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("bio", modifyCmd.PersistentFlags().Lookup("bio"))
	viper.BindPFlag("location", modifyCmd.PersistentFlags().Lookup("location"))
	viper.BindPFlag("admin", modifyCmd.PersistentFlags().Lookup("admin"))
	viper.BindPFlag("can_create_group", modifyCmd.PersistentFlags().Lookup("can_create_group"))
	// TODO currently not supported by go-gitlab API
	// viper.BindPFlag("external", modifyCmd.PersistentFlags().Lookup("external"))
	userCmd.AddCommand(modifyCmd)
}

func initUserDeleteCommand() {
	deleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(mandatory if no username is set) id of the user to be deleted")
	deleteCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "(mandatory if no id is set) username of the user to be deleted")
	viper.BindPFlag("id", deleteCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("user", deleteCmd.PersistentFlags().Lookup("user"))
	userCmd.AddCommand(deleteCmd)
}

func initListSshKeysCmd() {
	listSshKeysCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(optional) id of user to show ssh-keys for - if none is given, logged in user will be used")
	viper.BindPFlag("id", listSshKeysCmd.PersistentFlags().Lookup("id"))

	getSshKeyCmd.PersistentFlags().IntVarP(&keyId, "key_id", "k", 0, "(mandatory) key id of ssh key to be shown")
	viper.BindPFlag("key_id", getSshKeyCmd.PersistentFlags().Lookup("key_id"))

	listSshKeysCmd.AddCommand(getSshKeyCmd)

	userCmd.AddCommand(listSshKeysCmd)
}
