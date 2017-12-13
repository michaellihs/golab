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
	"errors"
)

var zshCompletionPath string

var zshCompletionCmd = &cobra.Command{
	Use: "zsh-completion",
	Short: "Generate ZSH completion file",
	Long: `Generate ZSH completion file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if zshCompletionPath == "" {
			return errors.New("required parameter `-p` or `--path` not given - exiting")
		}
		err := RootCmd.GenZshCompletionFile(zshCompletionPath)
		return err
	},
}

func init() {
	initZshCompletionCmd()
}

func initZshCompletionCmd() {
	zshCompletionCmd.PersistentFlags().StringVarP(&zshCompletionPath, "path", "p", "", "(required) Path into which to render ZSH completion")
	RootCmd.AddCommand(zshCompletionCmd)
}