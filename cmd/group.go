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
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO how can we namespace parameters to commands (this results in an error, that id is already declared in `cmd/project`
//var id string

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Gitlab Groups",
	Long: `Show, create, update and delete Gitlab groups.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("group called")
	},
}

var groupGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a group",
	Long: `Get detailed information for a group identified by either ID or the namespace / path of the group`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			fmt.Println("Required parameter `-i` or `--id` not given. Exiting.")
			os.Exit(1)
		}
		group, err := gitlabClient.Groups.GetGroup(id)
		if err != nil {
			fmt.Println("An error occurred (exiting): " + err.Error())
			os.Exit(1)
		}
		result, _ := json.MarshalIndent(group, "", "  ")
		fmt.Println(string(result))
	},
}

func init() {
	initGroupGetCommand()
	RootCmd.AddCommand(groupCmd)
}

func initGroupGetCommand() {
	groupGetCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "(required) Either ID or namespace of group")
	viper.BindPFlag("id", groupGetCmd.PersistentFlags().Lookup("id"))
	groupCmd.AddCommand(groupGetCmd)
}
