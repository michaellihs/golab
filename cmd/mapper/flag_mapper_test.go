package mapper

import (
	"reflect"
	"strings"
	"bytes"
	"os"
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/spf13/cobra"
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
		Flag1 *bool     `flag_name:"flag1" short:"f" type:"bool" required:"no" description:"first flag"`
		Flag2 *string   `flag_name:"flag2" type:"string" required:"no" description:"second flag"`
		Flag3 *int      `flag_name:"flag3" type:"string" required:"no" description:"second flag"`
		Flag4 *[]string `flag_name:"flag4" type:"string" required:"no" description:"second flag"`
	}

	type testOpts struct {
		Flag1 *bool
		Flag2 *string
		Flag3 *int
		Flag4 *[]string
	}

	It("provides a constructor that takes a cobra command as parameter", func() {
		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		Expect(reflect.TypeOf(flagMapper).String()).To(Equal("mapper.FlagMapper"))
	})

	It("sets expected flags on command", func() {
		flags := &testFlags{}

		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		flagMapper.SetFlags(flags)

		Expect(mockCmd.Flag("flag1")).NotTo(BeNil())
		Expect(mockCmd.Flag("flag1").Name).To(Equal("flag1"))
		Expect(mockCmd.Flag("flag1").Shorthand).To(Equal("f"))

		Expect(mockCmd.Flag("flag2")).NotTo(BeNil())
		Expect(mockCmd.Flag("flag2").Name).To(Equal("flag2"))
		Expect(mockCmd.Flag("flag2").Shorthand).To(Equal(""))

		Expect(mockCmd.Flag("flag3")).NotTo(BeNil())
		Expect(mockCmd.Flag("flag3").Name).To(Equal("flag3"))
		Expect(mockCmd.Flag("flag3").Shorthand).To(Equal(""))

		Expect(mockCmd.Flag("flag4")).NotTo(BeNil())
		Expect(mockCmd.Flag("flag4").Name).To(Equal("flag4"))
		Expect(mockCmd.Flag("flag4").Shorthand).To(Equal(""))
	})

	It("maps args to given struct as expected", func() {
		flags := &testFlags{}
		opts := &testOpts{}
		mockCmd := mockCmd()
		var flagMapper = New(mockCmd)
		flagMapper.SetFlags(flags)

		executeCommand(mockCmd, "mock", "--flag1", "true", "--flag2", "string", "--flag3", "4", "--flag4", "v1, v2, v3")
		flagMapper.Map(flags, opts)

		Expect(*opts.Flag1).To(Equal(true))
		Expect(*opts.Flag2).To(Equal("string"))
		Expect(*opts.Flag3).To(Equal(4))
		// TODO bug, when parsing array flags
		Expect(*opts.Flag4).Should(ConsistOf("v1, v2, v3"))
	})

})

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