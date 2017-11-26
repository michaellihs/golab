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

var _ = Describe("ls command", func() {

	var (
		mux    *http.ServeMux
		server *httptest.Server
	)

	BeforeEach(func() {
		// mux is the HTTP request multiplexer used with the test server.
		mux = http.NewServeMux()

		// server is a test HTTP server used to provide mock API responses.
		server = httptest.NewServer(mux)

		// client is the Gitlab client being tested.
		gitlabClient = gitlab.NewClient(nil, "")
		gitlabClient.SetBaseURL(server.URL + "/api/v4")
	})

	It("should exit with error, if no `--id` param is given", func() {
		// we don't want config file to be read (mocking)
		_, _, err := executeCommand(RootCmd, "group-members", "ls")
		if err == nil {
			Fail(fmt.Sprintf("Unexpected output: %v", err))
		}
		Expect(err.Error()).To(Equal("required parameter `-i` or `--id`not given - exiting"))
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
		if err != nil {
			Fail(fmt.Sprintf("unexpected error: %v", err))
		}
		if stdout == "" {
			Fail("we expected some output...")
		}
		Expect(method).To(Equal("GET"))
		Expect(stdout).To(Equal(expected))
	})
})
