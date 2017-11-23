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