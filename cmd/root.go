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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"net/url"
	"encoding/json"
)

var cfgFile string

var gitlabClient *gitlab.Client

var RootCmd = &cobra.Command{
	Use:   "golab",
	Short: "Gitlab CLI written in Go",
	Long: `This application provides a Command Line Interface for Gitlab.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func OutputJson(object interface{}) error {
	result, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGitlabClient)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "(optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".golab") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AddConfigPath(".")      // adding current directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}

func initGitlabClient() {
	baseUrl, err := url.Parse(viper.GetString("url"))
	if err != nil {
		fmt.Printf("Could not parse given URL '%s': %s", baseUrl, err)
	}
	// TODO this is an ugly hack to prevent re-initialization when mocked in testing
	if gitlabClient == nil {
		gitlabClient = gitlab.NewClient(nil, viper.GetString("token"))
		gitlabClient.SetBaseURL(baseUrl.String() + "/api/v4")
	}
}

func isoTime2String(time *gitlab.ISOTime) (string, error) {
	bytes, err := time.MarshalJSON()
	return string(bytes), err
}