package cmd

import (
	"fmt"
	"net/http/httptest"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("project command", func() {

	var (
		mux    *http.ServeMux
		server *httptest.Server
	)

	// TODO think about how we get this de-duplicated
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

	Context("when the `ls` command is exectued", func() {
		It("returns projects as expected", func() {
			defer server.Close()
			method := ""
			expected := readFixture("project-ls")
			mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
				method = r.Method
				fmt.Fprintf(w, expected)
			})
			stdout, _, err := executeCommand(RootCmd, "project", "ls")
			Expect(err).To(BeNil())
			Expect(stdout).NotTo(Equal(""))
			Expect(method).To(Equal("GET"))
			Expect(stdout).To(Equal(expected))
		})
	})

})