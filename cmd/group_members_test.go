package cmd

import (
	"fmt"
	"net/http/httptest"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("int2AccessLevel", func() {
	It("returns expected access level", func() {
		Expect(int2AccessLevel(10)).To(Equal(gitlab.AccessLevel(gitlab.GuestPermissions)))
		Expect(int2AccessLevel(20)).To(Equal(gitlab.AccessLevel(gitlab.ReporterPermissions)))
		Expect(int2AccessLevel(30)).To(Equal(gitlab.AccessLevel(gitlab.DeveloperPermissions)))
		Expect(int2AccessLevel(40)).To(Equal(gitlab.AccessLevel(gitlab.MasterPermissions)))
		Expect(int2AccessLevel(50)).To(Equal(gitlab.AccessLevel(gitlab.OwnerPermission)))
		Expect(int2AccessLevel(34)).To(BeNil())
	})
})

var _ = Describe("group-members command", func() {

	var (
		mux    *http.ServeMux
		server *httptest.Server
	)

	BeforeEach(func() {
		// do this to reset command line flags
		resetCommandLineFlagSet()

		// mux is the HTTP request multiplexer used with the test server.
		mux = http.NewServeMux()

		// server is a test HTTP server used to provide mock API responses.
		server = httptest.NewServer(mux)

		// client is the Gitlab client being tested.
		gitlabClient = gitlab.NewClient(nil, "")
		gitlabClient.SetBaseURL(server.URL + "/api/v4")
	})

	Context("when the `get` sub command is executed", func() {
		Context("if no `--id` or `--user-id` parameters are given", func() {
			It("should exit with error", func() {
				// TODO think about a better way to reset vars from previous runs...
				id = 0;
				userId = 0
				_, _, err := executeCommand(RootCmd, "group-members", "get")
				if err == nil {
					Fail("No error was thrown, when no --id was given")
				}
				Expect(err.Error()).To(Equal("required parameter `-i` or `--id`not given - exiting"))

				_, _, err = executeCommand(RootCmd, "group-members", "get", "-i", "19")
				if err == nil {
					Fail("No error was thrown, when --user-id was given")
				}
				fmt.Println(err.Error())
				Expect(err.Error()).To(Equal("required parameter `-u` or `--user_id`not given - exiting"))
			})
		})

		It("should return the expected group-members entry", func() {
			defer server.Close()
			method := ""
			expected := readFixture("group-members-get")
			mux.HandleFunc("/api/v4/groups/30/members/40", func(w http.ResponseWriter, r *http.Request) {
				method = r.Method
				fmt.Fprintf(w, expected)
			})
			stdout, _, err := executeCommand(RootCmd, "group-members", "get", "-i", "30", "-u", "40")
			Expect(err).To(BeNil())
			Expect(stdout).NotTo(Equal(""))
			Expect(method).To(Equal("GET"))
			Expect(stdout).To(Equal(expected))
		})
	})

	Context("when the `ls` sub command is executed", func() {
		Context("if no `--id` parameter is given", func() {
			It("should exit with error", func() {
				// TODO think about a better way to reset vars from previous runs...
				id = 0
				userId = 0
				_, _, err := executeCommand(RootCmd, "group-members", "ls")
				fmt.Println(err.Error())
				Expect(err).NotTo(BeNil(), "No error was raised when missing -i parameter")
				Expect(err.Error()).To(Equal("required parameter `-i` or `--id`not given - exiting"))
			})
		})

		It("returns expected group members", func() {
			defer server.Close()
			method := ""
			expected := readFixture("group-ls")
			mux.HandleFunc("/api/v4/groups/30/members", func(w http.ResponseWriter, r *http.Request) {
				method = r.Method
				fmt.Fprintf(w, expected)
			})
			stdout, _, err := executeCommand(RootCmd, "group-members", "ls", "-i", "30")
			Expect(err).To(BeNil())
			Expect(stdout).NotTo(Equal(""))
			Expect(method).To(Equal("GET"))
			Expect(stdout).To(Equal(expected))
		})
	})
})
