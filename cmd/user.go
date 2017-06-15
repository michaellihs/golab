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

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Gitlab users",
	Long: `Allows create, update and deletion of a user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
		if err != nil {
			return err
		}
		err = OutputJson(users)
		return err
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long: `Allows creation of a new user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Long: `Delete a user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("delete called")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user",
	Long: `Allows updating a user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update called")
	},
}

func init() {
	userCmd.AddCommand(updateCmd)
	userCmd.AddCommand(createCmd)
	userCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(userCmd)
}
