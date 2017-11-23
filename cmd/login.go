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
	"strings"
	"errors"

	"github.com/spf13/cobra"
    "github.com/howeyc/gopass"
	"github.com/xanzy/go-gitlab"
	"net/url"
	"github.com/spf13/viper"
	"io/ioutil"
	path2 "path"

)

var host string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a Gitlab server",
	Long: `Log in to a Gitlab server with your username and password.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if host == "" {
			return errors.New("required parameter `--host` or `-s` not given - exiting")
		}
		if username == "" {
			return errors.New("required parameter `--user` or `-u` not given - exiting")
		}
		if password == "" {
			var err error
			password, err = askForPassword()
			if err != nil { return err }
		}
		token, err := getPrivateToken(host, username, password)
		if err != nil { return err }
		err = writeGolabConf(host, token)
		if err != nil { return err}
		fmt.Printf("** successfully logged in to %s\n", host)
		return nil
	},
}

func writeGolabConf(host string, token string) error {
	conf := []byte(fmt.Sprintf("---\nurl: \"%s\"\ntoken: \"%s\"", host, token))
	pwd, err := os.Getwd()
	filename := path2.Join(pwd, ".golab.yml")
	err = ioutil.WriteFile(filename, conf, 0600)
	if err == nil {
		fmt.Printf("** golab config written to %s\n", filename)
	}
	return err
}

// see https://stackoverflow.com/questions/2137357/getpasswd-functionality-in-go
func askForPassword() (string, error) {
	fmt.Print("Enter Password: ")
	pass, err := gopass.GetPasswd()
	return strings.TrimSpace(string(pass)), err
}

func getPrivateToken(host string, username string, password string) (string, error) {
	opts := &gitlab.GetSessionOptions{
		Login: &username,
		Password: &password,
	}
	loginClient, err := getLoginClient(host)
	if err != nil { return "", err }
	session, _, err := loginClient.Session.GetSession(opts)
	if err != nil {
		return "", err
	}
	return session.PrivateToken, nil
}

func getLoginClient(host string) (*gitlab.Client, error) {
	baseUrl, err := url.Parse(host)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse given host '%s': %s", baseUrl, err))
	}
	loginClient := gitlab.NewClient(nil, viper.GetString("token"))
	loginClient.SetBaseURL(baseUrl.String() + "/api/v4")
	return loginClient, nil
}

func init() {
	loginCmd.PersistentFlags().StringVarP(&host, "host", "s", "", "(required) URL to Gitlab server eg. http://gitlab.org")
	loginCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "(required) username")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "(optional) password, if not given, you'll be prompted interactively")
	RootCmd.AddCommand(loginCmd)
}
