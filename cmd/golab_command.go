// Copyright © 2018 Michael Lihs
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
	"reflect"

	"github.com/michaellihs/golab/cmd/mapper"
	"github.com/spf13/cobra"
)

type golabCommand struct {
	Parent *cobra.Command
	Flags  interface{}
	Opts   interface{}
	Paged  bool
	Run    func(cmd golabCommand) error
	Mapper mapper.FlagMapper
	Cmd    *cobra.Command
}

func (c golabCommand) Execute() error {
	_, _, err := c.Mapper.AutoMap()
	if err != nil {
		return err
	}
	c.Flags = c.Mapper.MappedFlags()
	c.Opts = c.Mapper.MappedOpts()
	if err = applyPagination(c); err != nil {
		return err
	}
	return c.Run(c)
}

func applyPagination(c golabCommand) error {
	if c.Paged {
		optsReflected := reflect.ValueOf(c.Opts).Elem()
		page, err := c.Cmd.Flags().GetInt("page")
		if err != nil {
			return err
		}
		optsReflected.FieldByName("ListOptions").FieldByName("Page").Set(reflect.ValueOf(page))
		perPage, err := c.Cmd.Flags().GetInt("per_page")
		if err != nil {
			return err
		}
		optsReflected.FieldByName("ListOptions").FieldByName("PerPage").Set(reflect.ValueOf(perPage))
	}
	return nil
}

func (c golabCommand) Init() error {
	c.Cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return c.Execute()
	}
	c.Mapper = mapper.InitializedMapper(c.Cmd, c.Flags, c.Opts)
	setPaginationFlags(c)
	c.Parent.AddCommand(c.Cmd)
	return nil // TODO do something useful with the error return
}

func setPaginationFlags(c golabCommand) {
	if c.Paged {
		c.Cmd.PersistentFlags().Int("page", 0, "(optional) Page of results to retrieve")
		c.Cmd.PersistentFlags().Int("per_page", 0, "(optional) The number of results to include per page (max 100)")
	}
}
