package main

import (
	"errors"
	"fmt"
	"os"

	. "github.com/michaellihs/gogpat/gitlab"

	logrus "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// VERSION indicates which version of the binary is running.
var VERSION string

const gitlabDefaultURL = "https://gitlab.com"

// preload initializes any global options and configuration
// before the main or sub commands are run.
func preload(c *cli.Context) (err error) {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if len(c.Args()) < 1 {
		return errors.New("please supply filename(s)")
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "gogpat"
	app.ArgsUsage = "gogpat"
	app.Author = "@solidnerd"
	app.Email = "github@mietz.io"
	app.Usage = "gitlab personal access token cli."
	app.Before = preload
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "run in debug mode",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "create",
			Usage:    "creates a gitlab api token for the specified gitlab",
			HideHelp: true,
			Action:   create,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "api, a",
					Usage: "Access the authenticated user's API: Full access to GitLab as the user, including read/write on all their groups and projects",
				},

				cli.BoolFlag{
					Name:  "read_user, ru",
					Usage: "Read the authenticated user's personal information: Read-only access to the user's profile information, like username, public email and full name",
				},
				cli.BoolFlag{
					Name:  "read_registry, rr",
					Usage: "Read the authenticated user's personal information: Read-only access to the user's profile information, like username, public email and full name",
				},
				cli.BoolFlag{
					Name:  "sudo, s",
					Usage: "Perform API actions as any user in the system (if the authenticated user is an admin: Access to the Sudo feature, to perform API actions as any user in the system (only available for admins)",
				},
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Sets the user for the login",
				},
				cli.StringFlag{
					Name:  "password, p",
					Usage: "Sets the user for the login",
				},
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Sets the name of the personal token by default it's gogpat",
				},
				cli.StringFlag{
					Name:  "expiry, ex",
					Usage: "Sets the expiry date of the personal token it's should be in format like this 2017-12-22",
				},
			},
		}}
	app.Run(os.Args)
}

func create(c *cli.Context) {
	url := c.Args().Get(0)
	if url == "" {
		url = gitlabDefaultURL
	}
	user := c.String("user")
	if user == "" {
		logrus.Errorf("Please provide a valid username for the gitlab instance: %s", url)
	}
	password := c.String("password")
	if password == "" {
		logrus.Errorf("Please provide a valid password for the gitlab instance: %s", url)
	}
	name := c.String("name")
	api := c.Bool("api")
	readUser := c.Bool("read_user")
	readRegistry := c.Bool("read_registry")
	sudo := c.Bool("sudo")
	if api == false {
		logrus.Warnf("Scopes can't be blank setting by default --%s", c.FlagNames()[0])
		api = true
	}
	expiry := c.String("expiry")
	request := GitLabTokenRequest{
		URL:       url,
		Username:  user,
		Password:  password,
		Scope:     Scope{API: api, ReadUser: readUser, ReadRegistry: readRegistry, Sudo: sudo},
		Date:      expiry,
		TokenName: name,
	}
	token, err := CreateToken(request)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(token)
}
