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