package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/michaellihs/golab/client"

	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

func setup() (*http.ServeMux, *httptest.Server, *GitlabClient) {
	serveMux := http.NewServeMux()
	httpTestServer := httptest.NewServer(serveMux)

	hostUrl, _ := url.Parse(httpTestServer.URL)
	gitlabClient := NewClient(hostUrl, "test-token", http.DefaultClient)

	return serveMux, httpTestServer, gitlabClient
}

func teardown(server *httptest.Server) {
	server.Close()
}

func testURL(r *http.Request, want string) {
	Expect(r.RequestURI).To(Equal(want))
}

func testMethod(r *http.Request, want string) {
	Expect(r.Method).To(Equal(want))
}

func testHeaderContainsExpectedToken(r *http.Request, token string) {
	Expect(r.Header.Get("PRIVATE-TOKEN")).To(Equal(token))
}

type values map[string]string

func testFormValues(r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	err := r.ParseForm()
	Expect(err).To(BeNil())

	got := r.Form
	Expect(got).To(Equal(want))
}

func testJsonBody(r *http.Request, values values) {
	// somehow there seems to be a \n at the end
	body, _ := ioutil.ReadAll(r.Body)
	want, _ := json.Marshal(values)
	Expect(string(body)).Should(BeIdenticalTo(string(want) + "\n"))
}