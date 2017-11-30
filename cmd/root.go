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
	"net/url"
	"encoding/json"
	"net/http"
	"crypto/tls"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"github.com/hashicorp/go-rootcerts"
	"github.com/hashicorp/go-cleanhttp"
)

var cfgFile string

var gitlabClient *gitlab.Client

var RootCmd = &cobra.Command{
	Use:   "golab",
	Short: "Gitlab CLI written in Go",
	Long:  `This application provides a Command Line Interface for Gitlab.`,
}

func Execute() {
	initRootCommand()
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

func initRootCommand() {
	// TODO this is an ugly hack to prevent re-initialization when mocked in testing
	if gitlabClient == nil {
		cobra.OnInitialize(initConfig)
		cobra.OnInitialize(initGitlabClient)
	}

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "(optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)")
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".golab") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AddConfigPath(".")      // adding current directory as first search path
 	viper.AutomaticEnv()              // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func initGitlabClient() {
	baseUrl, err := url.Parse(viper.GetString("url"))
	if err != nil {
		fmt.Printf("Could not parse given URL '%s': %s", baseUrl, err)
	}

	httpClient, err := initHttpClient()
	if err != nil {
		panic("Error in initializing http client " + err.Error())
	}

	gitlabClient = gitlab.NewClient(httpClient, viper.GetString("token"))
	gitlabClient.SetBaseURL(baseUrl.String() + "/api/v4")
}

func initHttpClient() (*http.Client, error) {
	// see https://github.com/hashicorp/go-rootcerts
	tlsConfig := &tls.Config{}
	err := rootcerts.ConfigureTLS(tlsConfig, &rootcerts.Config{
		CAFile: os.Getenv("GOLAB_CAFILE"),
		CAPath: os.Getenv("GOLAB_CAPATH"),
	})
	if err != nil {
		return nil, err
	}
	c := cleanhttp.DefaultClient()
	t := cleanhttp.DefaultTransport()

	// use this line for debugging certificates
	//fmt.Println(tlsConfig.RootCAs)

	t.TLSClientConfig = tlsConfig
	c.Transport = t
	return c, nil

	// TODO this is an ugly hack to prevent SSL verification... see https://github.com/andygrunwald/go-jira/issues/52
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	//}
	//return &http.Client{Transport: tr}, nil
}

func isoTime2String(time *gitlab.ISOTime) (string, error) {
	bytes, err := time.MarshalJSON()
	return string(bytes), err
}
