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
