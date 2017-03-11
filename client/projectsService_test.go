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

		// TODO introduce errors here
		projects := projectsService.List()

		//Expect(err).To(Equal(nil))
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
