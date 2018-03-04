// Copyright © 2018 Michael Lihs
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

package helpers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GitHelper", func() {

	gh := GitHelper()

	var _ = Describe("GetWebUrl", func() {

		It("returns expected URL for http(s) origin", func() {
			gitRemotes := `origin  https://github.com/michaellihs/golab.git (fetch)
                           origin  https://github.com/michaellihs/golab.git (push)`
			remote, err := gh.GetRemoteUrl(gitRemotes)
			Expect(err).To(BeNil())
			Expect(gh.GetWebUrl(remote)).To(Equal("https://github.com/michaellihs/golab"))
		})

		It("returns expected URL for ssh origin", func() {
			gitRemotes := `origin  ssh://github.com/michaellihs/golab.git (fetch)
                           origin  ssh://github.com/michaellihs/golab.git (push)`
			remote, err := gh.GetRemoteUrl(gitRemotes)
			Expect(err).To(BeNil())
			Expect(gh.GetWebUrl(remote)).To(Equal("https://github.com/michaellihs/golab"))
		})

		It("returns expected URL for ssh origin with user", func() {
			gitRemotes := `origin  ssh://user@github.com/michaellihs/golab.git (fetch)
                           origin  ssh://user@github.com/michaellihs/golab.git (push)`
			remote, err := gh.GetRemoteUrl(gitRemotes)
			Expect(err).To(BeNil())
			Expect(gh.GetWebUrl(remote)).To(Equal("https://github.com/michaellihs/golab"))
		})

		It("returns expected URL for ssh origin with user and port", func() {
			gitRemotes := `origin  ssh://user@github.com:1234/michaellihs/golab.git (fetch)
                           origin  ssh://user@github.com:1234/michaellihs/golab.git (push)`
			remote, err := gh.GetRemoteUrl(gitRemotes)
			Expect(err).To(BeNil())
			Expect(gh.GetWebUrl(remote)).To(Equal("https://github.com/michaellihs/golab"))
		})

		It("returns expected URL for git origin", func() {
			gitRemotes := `origin  git@github.com:michaellihs/golab.git (fetch)
                           origin  git@github.com:michaellihs/golab.git (push)`
			remote, err := gh.GetRemoteUrl(gitRemotes)
			Expect(err).To(BeNil())
			Expect(gh.GetWebUrl(remote)).To(Equal("https://github.com/michaellihs/golab"))
		})

	})

})
