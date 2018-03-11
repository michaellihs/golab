// Copyright © 2018 Michael Lihs
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

// see https://docs.gitlab.com/ce/api/environments.html
var environmentsCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "environments",
		Aliases: []string{"environment", "env"},
		Short:   "Manage environments",
		Long:    `Manage environments`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("you cannot call this command without any subcommand")
	},
}

// see https://docs.gitlab.com/ce/api/environments.html#list-environments
type environmentsListFlag struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
}

var environmentsListCmd = &golabCommand{
	Parent: environmentsCmd.Cmd,
	Flags:  &environmentsListFlag{},
	Opts:   &gitlab.ListEnvironmentsOptions{},
	Cmd: &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List environments",
		Long:    `Get all environments for a project`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*environmentsListFlag)
		opts := cmd.Opts.(*gitlab.ListEnvironmentsOptions)
		environments, _, err := gitlabClient.Environments.ListEnvironments(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(environments)
	},
}

// see https://docs.gitlab.com/ce/api/environments.html#create-a-new-environment
type environmentsCreateFlags struct {
	Id          *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Name        *string `flag_name:"name" type:"string" required:"yes" description:"The name of the environment"`
	ExternalURL *string `flag_name:"external_url" type:"string" required:"no" description:"Place to link to for this environment"`
}

var environmentsCreateCmd = &golabCommand{
	Parent: environmentsCmd.Cmd,
	Flags:  &environmentsCreateFlags{},
	Opts:   &gitlab.CreateEnvironmentOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create a new environment",
		Long:  `Creates a new environment for the given project with the given name and external URL.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*environmentsCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateEnvironmentOptions)
		environment, _, err := gitlabClient.Environments.CreateEnvironment(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(environment)
	},
}

// see https://docs.gitlab.com/ce/api/environments.html#delete-an-environment
type environmentsDeleteFlags struct {
	Id            *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	EnvironmentId *int    `flag_name:"environment_id" short:"e" type:"integer" required:"yes" description:"The ID of the environment"`
}

var environmentsDeleteCmd = &golabCommand{
	Parent: environmentsCmd.Cmd,
	Flags:  &environmentsDeleteFlags{},
	Cmd: &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete an environment",
		Long:    `Deletes an environment with a given ID.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*environmentsDeleteFlags)
		_, err := gitlabClient.Environments.DeleteEnvironment(*flags.Id, *flags.EnvironmentId)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/environments.html#edit-an-existing-environment
type environmentsEditFlags struct {
	Id            *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	EnvironmentId *int    `flag_name:"environment_id" short:"e" type:"integer" required:"yes" description:"The ID of the environment"`
	Name          *string `flag_name:"name" type:"string" required:"no" description:"The new name of the environment"`
	ExternalURL   *string `flag_name:"external_url" type:"string" required:"yes" description:"The new external_url"`
}

var environmentsEditCmd = &golabCommand{
	Parent: environmentsCmd.Cmd,
	Flags:  &environmentsEditFlags{},
	Opts:   &gitlab.EditEnvironmentOptions{},
	Cmd: &cobra.Command{
		Use:     "edit",
		Aliases: []string{"update"},
		Short:   "Edit an existing environment",
		Long: `Updates an existing environment's name and/or external_url.

It returns 200 if the environment was successfully updated. In case of an error, a status code 400 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*environmentsEditFlags)
		opts := cmd.Opts.(*gitlab.EditEnvironmentOptions)
		environment, _, err := gitlabClient.Environments.EditEnvironment(*flags.Id, *flags.EnvironmentId, opts)
		if err != nil {
			return err
		}
		return OutputJson(environment)
	},
}

// see https://docs.gitlab.com/ce/api/environments.html#stop-an-environment
/* TODO not implemented in go-gitlab yet
type environmentsStopFlags struct {
	Id            *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	EnvironmentId *int    `flag_name:"environment_id" short:"e" type:"integer" required:"yes" description:"The ID of the environment"`
}

var environmentsStopCmd = &golabCommand{
	Parent: environmentsCmd.Cmd,
	Flags:  &environmentsStopFlags{},
	Cmd: &cobra.Command{
		Use:     "stop",
		Aliases: []string{""},
		Short:   "Stop an environment",
		Long:    `It returns 200 if the environment was successfully stopped, and 404 if the environment does not exist.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*environmentsStopFlags)
		return gitlabClient.Environments.StopEnvironemt(*flags.Id, *flags.EnvironmentId)
	},
}
*/

func init() {
	environmentsCmd.Init()
	environmentsListCmd.Init()
	environmentsCreateCmd.Init()
	environmentsEditCmd.Init()
	environmentsDeleteCmd.Init()
	//environmentsStopCmd.Init()
}
