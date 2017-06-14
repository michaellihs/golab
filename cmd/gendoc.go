package cmd

import (
	"github.com/spf13/cobra/doc"
	"github.com/spf13/cobra"
	"errors"
	"github.com/spf13/viper"
)

var docPath string

var gendocCmd = &cobra.Command{
	Use: "gendoc",
	Short: "Render the Markdown Documentation for golab",
	Long: `Renders the Markdown Documentation for golab into <PATH>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if docPath == "" {
			return errors.New("Required parameter `-p` or `--path` not given. Exiting.")
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
	viper.BindPFlag("path", groupGetCmd.PersistentFlags().Lookup("path"))
}