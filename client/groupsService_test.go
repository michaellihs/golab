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

package client_test

import (
	. "github.com/michaellihs/golab/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"fmt"
)

var _ = Describe("GroupsService", func() {

	It("sends expected URL to API for getting a project", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "GET")
			testURL(r, "/api/v3/groups/michaellihs2")
			testHeaderContainsExpectedToken(r, "test-token")
			fmt.Fprint(w, `{"id": 4, "name": "michaellihs2"}`)
		})

		groupsService := &GroupsService{Client: gitlabClient}
		group, err := groupsService.GetGroup("michaellihs2")

		Expect(err).To(BeNil())
		Expect(group.ID).To(Equal(4))
		Expect(group.Name).To(Equal("michaellihs2"))
	})

	Context("Namespacify", func() {
		serveMux, httpTestServer, gitlabClient := setup()

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"id": 888, "name": "path"}`)
		})

		groupsService := &GroupsService{Client: gitlabClient}

		It("returns the input value if a number is given", func() {
			namespace_id, err := groupsService.Namespacify("1234")

			Expect(err).To(BeNil())
			Expect(namespace_id).To(Equal(1234))
		})

		It("returns 0 if empty string is given as a parameter", func() {
			namespace_id, err := groupsService.Namespacify("")

			Expect(err).To(BeNil())
			Expect(namespace_id).To(Equal(0))
		})

		It("returns the namespace_id from group details if path is given", func() {
			// TODO if we put this in the "Context" section, server shuts down too early...
			defer teardown(httpTestServer)
			namespace_id, err := groupsService.Namespacify("path")

			Expect(err).To(BeNil())
			Expect(namespace_id).To(Equal(888))
		})
	})

})
