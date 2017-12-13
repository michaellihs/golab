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
	"github.com/spf13/cobra/doc"
	"github.com/spf13/cobra"
	"errors"
)

var docPath string

var gendocCmd = &cobra.Command{
	Use: "gendoc",
	Short: "Render the Markdown Documentation for golab",
	Long: `Renders the Markdown Documentation for golab into <PATH>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if docPath == "" {
			return errors.New("required parameter `-p` or `--path` not given")
		}

		err := doc.GenMarkdownTree(RootCmd, docPath)

		return err
	},
}

func init() {
	initGendocCommand()
	RootCmd.AddCommand(gendocCmd)
}

func initGendocCommand() {
	gendocCmd.PersistentFlags().StringVarP(&docPath, "path", "p", "", "(required) Path into which to render Markdown documentation")
}