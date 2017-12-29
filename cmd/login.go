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

	"io/ioutil"
	path2 "path"

	"github.com/howeyc/gopass"
	. "github.com/michaellihs/gogpat/gogpat"
	"github.com/spf13/cobra"
)

// loginCmd implements a user login with username and password
// that is not available in the Gitlab API. We use the
// Gitlab UI and some hacks to scrape a personal access token
// for a user identified by username and password.
type loginFlags struct {
	Host     *string `flag_name:"host" short:"s" type:"string" required:"yes" description:"Hostname (http://gitlab.my-domain.com) of the gitlab server"`
	User     *string `flag_name:"user" short:"u" type:"string" required:"yes" description:"Username for the login"`
	Password *string `flag_name:"password" short:"p" type:"string" required:"no" description:"Password for the login"`
}

var loginCmd = &golabCommand{
	Parent: RootCmd,
	Flags:  &loginFlags{},
	Cmd: &cobra.Command{
		Use:   "login",
		Short: "Login to Gitlab",
		Long:  `Login to Gitlab using username and password`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*loginFlags)
		if *flags.Password == "" {
			var err error
			password, err = askForPassword()
			if err != nil {
				return err
			}
			flags.Password = &password
		}
		req := GitLabTokenRequest{
			URL:      *flags.Host,
			Username: *flags.User,
			Password: *flags.Password,
			Scope:    Scope{API: true},
		}
		token, err := CreateToken(req)
		if err != nil {
			return err
		}
		// TODO add path to write config to
		err = writeGolabConf(*flags.Host, token)
		if err != nil {
			return err
		}
		fmt.Printf("** successfully logged in to %s\n", *flags.Host)
		return nil
	},
}

// personalAccessTokenCmd creates a personal access token for
// a user identified by username and password. This is not
// available from the Gitlab API so we login to the Gitlab UI
// and scrape the token from the generated HTML.
type personalAccessTokenFlags struct {
	Host         *string `flag_name:"host" short:"s" type:"string" required:"yes" description:"Hostname (http://gitlab.my-domain.com) of the gitlab server"`
	User         *string `flag_name:"user" short:"u" type:"string" required:"yes" description:"Username for the login"`
	Password     *string `flag_name:"password" short:"p" type:"string" required:"no" description:"Password for the login"`
	API          *bool   `flag_name:"api" short:"a" type:"bool" required:"no" description:"Access the authenticated user's API (default: false)"`
	ReadUser     *bool   `flag_name:"read_user" type:"bool" required:"no" description:"Read the authenticated user's personal information (default: false)"`
	ReadRegistry *bool   `flag_name:"read_registry" type:"bool" required:"no" description:"Grant access to the docker registry (default: false)"`
	Sudo         *bool   `flag_name:"sudo" type:"bool" required:"no" description:"Perform API actions as any user in the system (default: false)"`
	Date         *string `flag_name:"expires" type:"string" required:"no" description:"Expiration date of token"`
	TokenName    *string `flag_name:"token_name" type:"string" required:"no" description:"Name of token"`
}

var personalAccessTokenCmd = &golabCommand{
	Parent: RootCmd,
	Flags:  &personalAccessTokenFlags{},
	Cmd: &cobra.Command{
		Use:     "personal-access-token",
		Aliases: []string{"pat"},
		Short:   "Create a personal access token",
		Long:    `Create a personal access token for a user identified by username and password`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*personalAccessTokenFlags)

		if *flags.Password == "" {
			var err error
			password, err = askForPassword()
			if err != nil {
				return err
			}
			flags.Password = &password
		}

		req := GitLabTokenRequest{
			URL:      *flags.Host,
			Username: *flags.User,
			Password: *flags.Password,
			Scope:    Scope{},
		}

		if flags.API != nil {
			req.Scope.API = *flags.API
		}

		if flags.ReadRegistry != nil {
			req.Scope.ReadRegistry = *flags.ReadRegistry
		}

		if flags.ReadUser != nil {
			req.Scope.ReadUser = *flags.ReadUser
		}

		if flags.Sudo != nil {
			req.Scope.Sudo = *flags.Sudo
		}

		if flags.Date != nil {
			req.Date = *flags.Date
		}

		if flags.TokenName != nil {
			req.TokenName = *flags.TokenName
		} else {
			req.TokenName = "golab-generated"
		}

		token, err := CreateToken(req)
		if err != nil {
			return err
		}
		fmt.Println(token)

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

func init() {
	loginCmd.Init()
	personalAccessTokenCmd.Init()
}
