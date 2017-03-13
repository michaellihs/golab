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
	. "github.com/michaellihs/golab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"fmt"
)

var _ = Describe("ProjectsService", func() {

	It("sends expected GET request for getting a project and maps it as expected", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "GET")
			testURL(r, "/api/v3/projects/1234")
			testHeaderContainsExpectedToken(r, "test-token")
			fmt.Fprint(w, `{"id":1234, "name":"testproject"}`)
		})

		projectsService := &ProjectsService{Client: gitlabClient}
		project, err := projectsService.Get("1234")

		Expect(err).To(BeNil())
		Expect(project.ID).To(Equal(1234))
		Expect(project.Name).To(Equal("testproject"))
	})

	It("sends expected GET request for listing projects and maps them as expected", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "GET")
			testURL(r, "/api/v3/projects")
			testHeaderContainsExpectedToken(r, "test-token")
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		})

		projectsService := &ProjectsService{Client: gitlabClient}

		projects, err := projectsService.List()

		Expect(err).To(BeNil())
		Expect(projects).To(Equal(&[]Project{{ID: 1}, {ID: 2}}))
	})

	It("sends the expected POST request for creating a new project", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "POST")
			testURL(r, "/api/v3/projects")
			testHeaderContainsExpectedToken(r, "test-token")
			testJsonBody(r, values{"name": "testproject"})
			fmt.Fprint(w, `{"id" : 4, "name" : "testproject"}`)
		})

		projectsService := &ProjectsService{Client: gitlabClient}

		project, err := projectsService.Create(&ProjectParams{Name: "testproject"})

		Expect(err).To(BeNil())
		Expect(project.Name).To(Equal("testproject"))
		Expect(project.ID).To(Equal(4))
	})

	It("sends the expected DELETE request for deleting a project by its ID", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "DELETE")
			testURL(r, "/api/v3/projects/1234")
			testHeaderContainsExpectedToken(r, "test-token")
			fmt.Fprint(w, "true")

			// TODO: test case of error:
			// {
			//     "message": "404 Project Not Found"
			// }
		})

		projectsService := &ProjectsService{Client: gitlabClient}

		result, err := projectsService.Delete("1234")

		Expect(err).To(BeNil())
		Expect(result).To(Equal(true))
	})

	It("sends the expected DELETE request for deleting a project by 'namespace/project-name'", func() {
		serveMux, httpTestServer, gitlabClient := setup()
		defer teardown(httpTestServer)

		serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			testMethod(r, "DELETE")
			testURL(r, "/api/v3/projects/michaellihs%2Fhasenfurz")
			testHeaderContainsExpectedToken(r, "test-token")
			fmt.Fprint(w, "true")

			// TODO: test case of error:
			// {
			//     "message": "404 Project Not Found"
			// }
		})

		projectsService := &ProjectsService{Client: gitlabClient}

		result, err := projectsService.Delete("michaellihs/hasenfurz")

		Expect(err).To(BeNil())
		Expect(result).To(Equal(true))
	})

})
