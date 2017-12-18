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
	"github.com/spf13/cobra"
	"errors"
	"github.com/xanzy/go-gitlab"
)

// see https://docs.gitlab.com/ce/api/deploy_keys.html#deploy-keys-api
var deployKeysCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "deploy-keys",
		Aliases: []string{"dk"},
		Short:   "Deploy Keys API",
		Long:    `Manage deploy keys`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot run this command without further subcommands")
	},
}

// see https://docs.gitlab.com/ce/api/deploy_keys.html#list-all-deploy-keys
var deployKeysListAllCmd = &golabCommand{
	Parent: deployKeysCmd.Cmd,
	Cmd: &cobra.Command{
		Use:     "list-all",
		Aliases: []string{"lsa"},
		Short:   "List all deploy keys",
		Long:    `Get a list of all deploy keys across all projects of the GitLab instance. This endpoint requires admin access.`,
	},
	Run: func(cmd golabCommand) error {
		keys, _, err := gitlabClient.DeployKeys.ListAllDeployKeys()
		if err != nil {
			return err
		}
		return OutputJson(keys)
	},
}

// see https://docs.gitlab.com/ce/api/deploy_keys.html#list-project-deploy-keys
type deployKeysListAllForProjectFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL encoded path of a project"`
}

var deployKeysListAllForProjectCmd = &golabCommand{
	Parent: deployKeysCmd.Cmd,
	Flags:  &deployKeysListAllForProjectFlags{},
	Cmd: &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List project deploy keys",
		Long:    `Get a list of a project's deploy keys.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*deployKeysListAllForProjectFlags)
		keys, _, err := gitlabClient.DeployKeys.ListProjectDeployKeys(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(keys)
	},
}

// see https://docs.gitlab.com/ce/api/deploy_keys.html#single-deploy-key
type deployKeysGetSingleFlags struct {
	Id    *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId *int    `flag_name:"key_id" short:"k" type:"integer" required:"yes" description:"The ID of the deploy key"`
}

var deployKeysGetSingleCmd = &golabCommand{
	Parent: deployKeysCmd.Cmd,
	Flags:  &deployKeysGetSingleFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get single deploy key",
		Long:  `Get a single deploy key`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*deployKeysGetSingleFlags)
		key, _, err := gitlabClient.DeployKeys.GetDeployKey(*flags.Id, *flags.KeyId)
		if err != nil {
			return err
		}
		return OutputJson(key)
	},
}

// see https://docs.gitlab.com/ce/api/deploy_keys.html#add-deploy-key
type deployKeysAddFlags struct {
	Id      *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Title   *string `flag_name:"title" short:"t" type:"string" required:"yes" description:"New deploy key's title"`
	Key     *string `flag_name:"key" short:"k" type:"string" required:"yes" description:"New deploy key"`
	CanPush *bool   `flag_name:"can_push" short:"p" type:"boolean" required:"no" description:"Can deploy key push to the project's repository"`
}

var deployKeysAddCmd = &golabCommand{
	Parent: deployKeysCmd.Cmd,
	Flags:  &deployKeysAddFlags{},
	Opts:   &gitlab.AddDeployKeyOptions{},
	Cmd: &cobra.Command{
		Use:     "add",
		Aliases: []string{"create"},
		Short:   "Add deploy key",
		Long: `Creates a new deploy key for a project.

If the deploy key already exists in another project, it will be joined to current project only if original one is accessible by the same user.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*deployKeysAddFlags)
		opts := cmd.Opts.(*gitlab.AddDeployKeyOptions)
		key, _, err := gitlabClient.DeployKeys.AddDeployKey(*flags.Id, opts)
		if err != nil {
		    return err
		}
		return OutputJson(key)
	},
}

// see https://docs.gitlab.com/ce/api/deploy_keys.html#delete-deploy-key
type deplyKeysDeleteFlags struct {
	Id    *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId *int    `flag_name:"key_id" short:"k" type:"integer" required:"yes" description:"The ID of the deploy key"`
}

var deplyKeysDeleteCmd = &golabCommand{
	Parent: deployKeysCmd.Cmd,
	Flags:  &deplyKeysDeleteFlags{},
	Cmd:    &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete deploy key",
		Long:    `Removes a deploy key from the project. If the deploy key is used only for this project, it will be deleted from the system.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*deplyKeysDeleteFlags)
		_, err := gitlabClient.DeployKeys.DeleteDeployKey(*flags.Id, *flags.KeyId)
		return err
	},
}

func init() {
	deployKeysCmd.Init()
	deployKeysListAllCmd.Init()
	deployKeysListAllForProjectCmd.Init()
	deployKeysGetSingleCmd.Init()
	deployKeysAddCmd.Init()
	deplyKeysDeleteCmd.Init()
}
