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

// see https://docs.gitlab.com/ce/api/labels.html#labels-api
var labelsCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "labels",
		Aliases: []string{"label"},
		Short:   "Manage labels",
		Long:    `Manage labels`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("you cannot call this command without any subcommand")
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#list-labels
type labelsListFlag struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
}

var labelsListCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsListFlag{},
	Cmd: &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List labels",
		Long:    `Get all labels for a project`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsListFlag)
		labels, _, err := gitlabClient.Labels.ListLabels(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(labels)
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#create-a-new-label
type labelsCreateFlags struct {
	Id          *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Name        *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the label"`
	Color       *string `flag_name:"color" short:"c" type:"string" required:"yes" description:"The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the CSS color names"`
	Description *string `flag_name:"description" short:"d" type:"string" required:"no" description:"The description of the label"`
	// TODO this is currently not available in go-gitlab
	Priority *int `flag_name:"priority" short:"p" type:"integer" required:"no" description:"The priority of the label. Must be greater or equal than zero or null to remove the priority."`
}

var labelsCreateCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsCreateFlags{},
	Opts:   &gitlab.CreateLabelOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create a new label",
		Long:  `Creates a new label for the given repository with the given name and color.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateLabelOptions)
		label, _, err := gitlabClient.Labels.CreateLabel(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(label)
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#delete-a-label
type labelsDeleteFlags struct {
	Id   *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Name *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the label"`
}

var labelsDeleteCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsDeleteFlags{},
	Opts:   &gitlab.DeleteLabelOptions{},
	Cmd: &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete a label",
		Long:    `Deletes a label with a given name.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsDeleteFlags)
		opts := cmd.Opts.(*gitlab.DeleteLabelOptions)
		_, err := gitlabClient.Labels.DeleteLabel(*flags.Id, opts)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#edit-an-existing-label
type labelsEditFlags struct {
	Id   *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Name *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the existing label"`
	// TODO think about an optional tag, that provides the "required / optional" message
	NewName     *string `flag_name:"new_name" short:"u" type:"string" required:"no" description:"(required, if color is not provided) The new name of the label"`
	Color       *string `flag_name:"color" short:"c" type:"string" required:"no" description:"(required, if new_name is not provided) The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the CSS color names"`
	Description *string `flag_name:"description" short:"d" type:"string" required:"no" description:"The new description of the label"`
	// TODO this is currently not available in go-gitlab
	Priority *int `flag_name:"priority" short:"p" type:"integer" required:"no" description:"The new priority of the label. Must be greater or equal than zero or null to remove the priority."`
}

var labelsEditCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsEditFlags{},
	Opts:   &gitlab.UpdateLabelOptions{},
	Cmd: &cobra.Command{
		Use:     "edit",
		Aliases: []string{"update"},
		Short:   "Edit an existing label",
		Long:    `Updates an existing label with new name or new color. At least one parameter is required, to update the label.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsEditFlags)
		if flags.NewName == nil && flags.Color == nil {
			return errors.New("either --new_name or --color is required")
		}
		opts := cmd.Opts.(*gitlab.UpdateLabelOptions)
		label, _, err := gitlabClient.Labels.UpdateLabel(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(label)
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#subscribe-to-a-label
type labelsSubscribeFlags struct {
	Id      *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	LabelId *string `flag_name:"label_id" short:"l" type:"string" required:"yes" description:"The ID or title of a project's label"`
}

var labelsSubsribeCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsSubscribeFlags{},
	Cmd: &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to a label",
		Long:  `Subscribes the authenticated user to a label to receive notifications. If the user is already subscribed to the label, the status code 304 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsSubscribeFlags)
		l, _, err := gitlabClient.Labels.SubscribeToLabel(*flags.Id, *flags.LabelId)
		if err != nil {
		    return err
		}
		return OutputJson(l)
	},
}

// see https://docs.gitlab.com/ce/api/labels.html#unsubscribe-from-a-label
type labelsUnsubscribeFlags struct {
	Id      *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	LabelId *string `flag_name:"label_id" short:"l" type:"string" required:"yes" description:"The ID or title of a project's label"`
}

var labelsUnsubscribeCmd = &golabCommand{
	Parent: labelsCmd.Cmd,
	Flags:  &labelsUnsubscribeFlags{},
	Cmd: &cobra.Command{
		Use:   "unsubscribe",
		Short: "Unsubscribe from a label",
		Long:  `Unsubscribes the authenticated user from a label to not receive notifications from it. If the user is not subscribed to the label, the status code 304 is returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*labelsUnsubscribeFlags)
		_, err  := gitlabClient.Labels.UnsubscribeFromLabel(*flags.Id, *flags.LabelId)
		return err
	},
}

func init() {
	labelsCmd.Init()
	labelsListCmd.Init()
	labelsCreateCmd.Init()
	labelsDeleteCmd.Init()
	labelsEditCmd.Init()
	labelsSubsribeCmd.Init()
	labelsUnsubscribeCmd.Init()
}
