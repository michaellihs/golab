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
	Priority    *int    `flag_name:"priority" short:"p" type:"integer" required:"no" description:"The priority of the label. Must be greater or equal than zero or null to remove the priority."`
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

func init() {
	labelsCmd.Init()
	labelsListCmd.Init()
	labelsCreateCmd.Init()
}
