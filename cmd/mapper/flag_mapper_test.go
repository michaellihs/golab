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

package mapper

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

func mockCmd() *Command {
	mock := &Command{
		Use:   "mock",
		Short: "mock",
		Long:  `mock`,
		RunE: func(cmd *Command, args []string) error {
			return nil
		},
	}
	return mock
}

var _ = Describe("FlagMapper", func() {

	type testFlags struct {
		Flag1 *bool     `flag_name:"flag1" short:"f" type:"bool" required:"yes" description:"first flag"`
		Flag2 *string   `flag_name:"flag2" type:"string" required:"no" description:"second flag"`
		Flag3 *int      `flag_name:"flag3" type:"string" required:"no" description:"third flag"`
		Flag4 *[]string `flag_name:"flag4" type:"[]string" required:"no" description:"fourth flag"`
		Flag5 []int     `flag_name:"flag5" type:"[]int" required:"no" description:"fifth flag"`
	}

	type testOpts struct {
		Flag1 *bool
		Flag2 *string
		Flag3 *int
		Flag4 *[]string
		Flag5 []int
	}

	type testOptsNonMatching struct {
		Flag1 *bool
		Flag2 *string
		Flag3 *int
		Flag4 *string // non matching with flags
	}

	type testFlagsWithTransformation struct {
		Flag1 *string `flag_name:"visibility" type:"bool" required:"no" description:"first flag" transform:"string2visibility"`
	}

	type optsRequireTransformation struct {
		Flag1 *gitlab.VisibilityValue
	}

	type testFlagsWithLabelsTransformation struct {
		Labels *string `flag_name:"labels" type:"string" required:"no" description:"labels" transform:"string2Labels"`
	}

	type optsRequireLabelsTransformation struct {
		Labels gitlab.Labels
	}

	type testFlagsWithPropertyNotInOpts struct {
		Id   *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"id"`
		Name *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"name"`
	}

	type testOptsWithMissingProperty struct {
		Name *string `flag_name:"name" short:"n" type:"string" required:"yes" description:"name"`
	}

	It("provides a constructor that takes a cobra command as parameter", func() {
		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		Expect(reflect.TypeOf(flagMapper).String()).To(Equal("mapper.FlagMapper"))
	})

	It("sets expected flags on command", func() {
		flags := &testFlags{}

		cmd := mockCmd()
		var mapper = New(cmd)
		mapper.SetFlags(flags)

		Expect(cmd.Flag("flag1")).NotTo(BeNil())
		Expect(cmd.Flag("flag1").Name).To(Equal("flag1"))
		Expect(cmd.Flag("flag1").Usage).To(Equal("(required) first flag"))
		Expect(cmd.Flag("flag1").Shorthand).To(Equal("f"))

		Expect(cmd.Flag("flag2")).NotTo(BeNil())
		Expect(cmd.Flag("flag2").Name).To(Equal("flag2"))
		Expect(cmd.Flag("flag2").Usage).To(Equal("(optional) second flag"))
		Expect(cmd.Flag("flag2").Shorthand).To(Equal(""))

		Expect(cmd.Flag("flag3")).NotTo(BeNil())
		Expect(cmd.Flag("flag3").Name).To(Equal("flag3"))
		Expect(cmd.Flag("flag3").Usage).To(Equal("(optional) third flag"))
		Expect(cmd.Flag("flag3").Shorthand).To(Equal(""))

		Expect(cmd.Flag("flag4")).NotTo(BeNil())
		Expect(cmd.Flag("flag4").Name).To(Equal("flag4"))
		Expect(cmd.Flag("flag4").Usage).To(Equal("(optional) fourth flag"))
		Expect(cmd.Flag("flag4").Shorthand).To(Equal(""))

		Expect(cmd.Flag("flag5")).NotTo(BeNil())
		Expect(cmd.Flag("flag5").Name).To(Equal("flag5"))
		Expect(cmd.Flag("flag5").Usage).To(Equal("(optional) fifth flag"))
		Expect(cmd.Flag("flag5").Shorthand).To(Equal(""))
	})

	It("sets no flags if given flags are nil", func() {
		mockCmd := mockCmd()
		mapper := New(mockCmd)
		mapper.SetFlags(nil)
	})

	It("maps flags and opts as expected with AutoMap", func() {
		mockCmd := mockCmd()
		var mapper = InitializedMapper(mockCmd, &testFlags{}, &testOpts{})

		executeCommand(mockCmd, "mock", "--flag1", "true", "--flag2", "string", "--flag3", "4", "--flag4", "v1, v2, v3", "--flag5", "1,2,3,4")
		_, _, err := mapper.AutoMap()
		flags := mapper.MappedFlags().(*testFlags)
		opts := mapper.MappedOpts().(*testOpts)

		Expect(err).To(BeNil())
		Expect(*flags.Flag1).To(Equal(true))
		Expect(*opts.Flag1).To(Equal(true))
	})

	It("maps valid args to given opts struct as expected", func() {
		flags := &testFlags{}
		opts := &testOpts{}
		cmd := mockCmd()
		var mapper = New(cmd)
		mapper.SetFlags(flags)

		executeCommand(cmd, "mock", "--flag1", "true", "--flag2", "string", "--flag3", "4", "--flag4", "v1, v2, v3", "--flag5", "1,2,3,4")
		mapper.Map(flags, opts)

		Expect(*opts.Flag1).To(Equal(true))
		Expect(*opts.Flag2).To(Equal("string"))
		Expect(*opts.Flag3).To(Equal(4))
		// TODO there seems to be bug in cobra, when parsing array flags
		Expect(*opts.Flag4).Should(ConsistOf("v1, v2, v3"))
		Expect(opts.Flag5).Should(ConsistOf(1, 2, 3, 4))
	})

	It("maps args to given flags struct as expected", func() {
		flags := &testFlags{}
		opts := &testOpts{}
		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		flagMapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "--flag1", "false", "--flag2", "string", "--flag3", "4", "--flag4", "v1, v2, v3")
		flagMapper.Map(flags, opts)

		Expect(*flags.Flag1).To(Equal(true))
		Expect(*flags.Flag2).To(Equal("string"))
		Expect(*flags.Flag3).To(Equal(4))
		Expect(*flags.Flag4).Should(ConsistOf("v1, v2, v3"))
	})

	It("maps nil flags as expected", func() {
		cmd := mockCmd()
		mapper := InitializedMapper(cmd, nil, nil)
		opts, flags, err := mapper.AutoMap()
		Expect(opts).To(BeNil())
		Expect(flags).To(BeNil())
		Expect(err).To(BeNil())
	})

	It("skips args with non-matching types as expected", func() {
		flags := &testFlags{}
		opts := &testOptsNonMatching{}
		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		flagMapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "--flag1", "true", "--flag2", "string", "--flag3", "4", "--flag4", "v1, v2, v3")
		flagMapper.Map(flags, opts)

		Expect(opts.Flag4).To(BeNil())
	})

	It("calls a transform function as expected", func() {
		flags := &testFlagsWithTransformation{}
		opts := &optsRequireTransformation{}
		cmd := mockCmd()
		var flagMapper = InitializedMapper(cmd, flags, opts)

		executeCommand(cmd, "mock", "--visibility", "private")
		flagMapper.AutoMap()

		Expect(*opts.Flag1).To(Equal(*gitlab.Visibility(gitlab.PrivateVisibility)))
	})

	It("transforms string to gitlab.Labels as expected", func() {
		flags := &testFlagsWithLabelsTransformation{}
		opts := &optsRequireLabelsTransformation{}
		cmd := mockCmd()
		var mapper = InitializedMapper(cmd, flags, opts)

		executeCommand(cmd, "mock", "--labels", "label1,label2,label3")
		mapper.AutoMap()

		Expect(opts.Labels).Should(ConsistOf("label1", "label2", "label3"))
	})

	It("transforms string to time.Time value as expected", func() {
		type str2timeValFlags struct {
			Time *string `flag_name:"time" type:"string" required:"no" description:"time" transform:"string2TimeVal"`
		}
		type str2timeValOpts struct {
			Time time.Time
		}
		flags := &str2timeValFlags{}
		opts := &str2timeValOpts{}
		cmd := mockCmd()
		var mapper = InitializedMapper(cmd, flags, opts)

		executeCommand(cmd, "mock", "--time", "2017-12-13")
		mapper.AutoMap()

		Expect(opts.Time).NotTo(BeNil())
		Expect(opts.Time.Day()).To(Equal(13))
		Expect(opts.Time.Month()).To(Equal(time.Month(12)))
		Expect(opts.Time.Year()).To(Equal(2017))
	})

	It("transforms JSON to commit actions as expected", func() {
		type json2CommitActionsFlags struct {
			Actions *string `flag_name:"actions" transform:"json2CommitActions" type:"array" required:"yes" description:"A JSON encoded array of action hashes to commit as a batch."`
		}
		type json2CommitActionsOpts struct {
			Actions []*gitlab.CommitAction
		}
		json := `[
			{
			  "action": "create",
			  "file_path": "foo/bar",
			  "content": "some content"
			},
			{
			  "action": "delete",
			  "file_path": "foo/bar2"
			},
			{
			  "action": "move",
			  "file_path": "foo/bar3",
			  "previous_path": "foo/bar4",
			  "content": "some content"
			},
			{
			  "action": "update",
			  "file_path": "foo/bar5",
			  "content": "new content"
			}
	  	]`
		flags := &json2CommitActionsFlags{}
		opts := &json2CommitActionsOpts{}
		cmd := mockCmd()
		var mapper = InitializedMapper(cmd, flags, opts)

		executeCommand(cmd, "mock", "--actions", json)
		mapper.AutoMap()

		Expect(len(opts.Actions)).To(Equal(4))
	})

	It("silently ignores properties in flags that are not available in opts", func() {
		flags := &testFlagsWithPropertyNotInOpts{}
		opts := &testOptsWithMissingProperty{}
		mockCmd := mockCmd()
		var mapper = New(mockCmd)
		mapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "-i", "34", "-n", "name")
		mapper.Map(flags, opts)

		Expect(*opts.Name).To(Equal("name"))
	})

	It("can handle nil opts", func() {
		flags := &testFlagsWithPropertyNotInOpts{}
		mockCmd := mockCmd()
		var mapper = New(mockCmd)
		mapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "-i", "34", "-n", "name")
		mapper.Map(flags, nil)

		Expect(*flags.Name).To(Equal("name"))
	})

	It("returns an error during mapping, if required flag is not set in mapping", func() {
		flags := &testFlags{}
		mockCmd := mockCmd()
		var mapper = New(mockCmd)
		mapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "-n", "name")
		err := mapper.Map(flags, nil)

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("required flag --flag1 was empty"))
	})

	It("returns an error during mapping, if required flag is not set in auto-mapping", func() {
		flags := &testFlags{}
		mockCmd := mockCmd()
		var mapper = InitializedMapper(mockCmd, flags, nil)

		executeCommand(mockCmd, "mock", "-n", "name")
		_, _, err := mapper.AutoMap()

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("required flag --flag1 was empty"))
	})

})

// TODO put the following methods into a testhelper

// see https://github.com/spf13/cobra/blob/master/command_test.go for basic implementation of this method
func executeCommand(root *Command, args ...string) (stdout string, output string, err error) {
	stdout, output, err = executeCommandC(root, args...)
	return strings.TrimRight(stdout, "\n"), output, err
}

// see https://github.com/spf13/cobra/blob/master/command_test.go for basic implementation of this method
func executeCommandC(root *Command, args ...string) (stdout string, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf) // this only redirects stderr, not stdout!
	root.SetArgs(args)

	// for capturing stdout, see https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	_, err = root.ExecuteC()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	stdout = <-outC

	return stdout, buf.String(), err
}
