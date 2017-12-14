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

// see https://docs.gitlab.com/ce/api/namespaces.html#namespaces-api
var namespacesCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:   "namespaces",
		Short: "Manage namespaces",
		Long:  `Usernames and groupnames fall under a special category called namespaces.

For users and groups supported API calls see the users and groups documentation respectively.

Pagination is used.`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("this command cannot be used without sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/namespaces.html#list-namespaces
type namespaceListFlags struct {
	// only for pagination
}

var namespacesListCmd = &golabCommand{
	Parent: namespacesCmd.Cmd,
	Flags:  &namespaceListFlags{},
	Opts:   &gitlab.ListNamespacesOptions{},
	Paged:  true,
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List namespaces",
		Long:  `Get a list of the namespaces of the authenticated user. If the user is an administrator, a list of all namespaces in the GitLab instance is shown.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.ListNamespacesOptions)
		ns, _, err := gitlabClient.Namespaces.ListNamespaces(opts)
		if err != nil {
		    return err
		}
		return OutputJson(ns)
	},
}

// see https://docs.gitlab.com/ce/api/namespaces.html#search-for-namespace
type namespacesSearchFlags struct {
	Search *string `flag_name:"search" type:"string" required:"no" description:"Returns a list of namespaces the user is authorized to see based on the search criteria"`
}

var namespacesSearchCmd = &golabCommand{
	Parent: namespacesCmd.Cmd,
	Flags:  &namespacesSearchFlags{},
	Cmd: &cobra.Command{
		Use:   "search",
		Short: "Search for namespace",
		Long:  `Get all namespaces that match a string in their name or path.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*namespacesSearchFlags)
		ns, _, err := gitlabClient.Namespaces.SearchNamespace(*flags.Search)
		if err != nil {
		    return err
		}
		return OutputJson(ns)
	},
}

func init() {
	namespacesCmd.Init()
	namespacesListCmd.Init()
	namespacesSearchCmd.Init()
}
