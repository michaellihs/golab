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