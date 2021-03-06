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
	"io/ioutil"
)

var _ = Describe("int2AccessLevel", func() {
	It("returns expected access level", func() {
		Expect(int2AccessLevel(10)).To(Equal(gitlab.AccessLevel(gitlab.GuestPermissions)))
		Expect(int2AccessLevel(20)).To(Equal(gitlab.AccessLevel(gitlab.ReporterPermissions)))
		Expect(int2AccessLevel(30)).To(Equal(gitlab.AccessLevel(gitlab.DeveloperPermissions)))
		Expect(int2AccessLevel(40)).To(Equal(gitlab.AccessLevel(gitlab.MasterPermissions)))
		Expect(int2AccessLevel(50)).To(Equal(gitlab.AccessLevel(gitlab.OwnerPermission)))
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
				Expect(err).NotTo(BeNil(), "No error was raised when missing -i parameter")
				Expect(err.Error()).To(Equal("required parameter `-i` or `--id`not given - exiting"))
			})
		})

		It("returns expected group members", func() {
			defer server.Close()
			method := ""
			expected := readFixture("group-members-ls")
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

	Context("when teh `add` sub command is executed", func() {
		Context("if no `--id`, `--user_id` or `--access_level` parameter is given", func() {
			It("should exit with error", func() {
				// TODO think about a better way to reset vars from previous runs...
				id = 0
				userId = 0
				accessLevel = 0
				_, _, err := executeCommand(RootCmd, "group-members", "add")
				Expect(err).NotTo(BeNil(), "No error was raised when missing -i parameter")
				Expect(err.Error()).To(Equal("required parameter `-i` or `--id` not given - exiting"))

				_, _, err = executeCommand(RootCmd, "group-members", "add", "-i", "30")
				Expect(err).NotTo(BeNil(), "No error was raised when missing -u parameter")
				Expect(err.Error()).To(Equal("required parameter `-u` or `--user_id` not given - exiting"))

				_, _, err = executeCommand(RootCmd, "group-members", "add", "-i", "30", "-u", "30")
				Expect(err).NotTo(BeNil(), "No error was raised when missing -u parameter")
				Expect(err.Error()).To(Equal("required parameter `-a` or `--access_level` not given - exiting"))
			})
		})

		It("creates group member as expected", func() {
			defer server.Close()
			// TODO think about a better way to reset vars from previous runs...
			method := ""
			body := ""
			expected := readFixture("group-members-add")
			mux.HandleFunc("/api/v4/groups/30/members", func(w http.ResponseWriter, r *http.Request) {
				method = r.Method
				bodyBytes, _ := ioutil.ReadAll(r.Body)
				body = string(bodyBytes)
				fmt.Fprintf(w, expected)
			})
			stdout, _, err := executeCommand(RootCmd, "group-members", "add", "-i", "30", "-u", "40", "-a", "50", "-e", "2016-09-23")
			Expect(err).To(BeNil())
			Expect(stdout).NotTo(Equal(""))
			Expect(method).To(Equal("POST"))
			Expect(body).To(Equal(`{"user_id":40,"access_level":50,"expires_at":"2016-09-23"}`))
			Expect(stdout).To(Equal(expected))
		})
	})
})
