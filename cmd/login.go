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
	"fmt"
	"github.com/spf13/viper"
)

var host string
var token string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a Gitlab server",
	Long: `Log in to a Gitlab server with the URL given in <host> and <token>.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement me
		fmt.Println(fmt.Sprintf("Logging in to Gitlab server with host: %s and token: %s", host, token))
	},
}

func init() {
	loginCmd.PersistentFlags().StringVarP(&host, "host", "u", "", "URL to Gitlab server eg. http://gitlab.org")
	loginCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Access token for the Gitlab server account")
	viper.BindPFlag("host", loginCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("token", loginCmd.PersistentFlags().Lookup("token"))

	RootCmd.AddCommand(loginCmd)
}
