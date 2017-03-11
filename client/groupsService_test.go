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

})
