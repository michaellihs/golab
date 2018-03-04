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
	"errors"
	"net"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
)

type gitHelper struct {
}

func GitHelper() gitHelper {
	return gitHelper{}
}

func (g gitHelper) GetRemotes() (string, error) {
	out, err := exec.Command("git", "remote", "-v").Output()
	return string(out), err
}

func (g gitHelper) GetRemoteUrl(remotes string) (string, error) {
	re := regexp.MustCompile("origin\\s*(.+?)\\s*\\(fetch\\)")
	match := re.FindStringSubmatch(remotes)
	if len(match) > 1 {
		return match[1], nil
	}
	return "", errors.New("Could not find URL in " + remotes)
}

func (g gitHelper) GetWebUrl(remoteUrl string) (string, error) {
	if strings.HasPrefix(remoteUrl, "http") {
		return webifyHttpRemote(remoteUrl)
	} else if strings.HasPrefix(remoteUrl, "git@") {
		return webifyGitRemote(remoteUrl)
	} else if strings.HasPrefix(remoteUrl, "ssh") {
		return webifySshRemote(remoteUrl)
	} else {
		return "", errors.New("Cannot parse remote URL: " + remoteUrl)
	}
}

func webifySshRemote(remoteUrl string) (string, error) {
	u, err := url.Parse(remoteUrl)
	if err != nil {
		return "", err
	}
	host := ""
	if strings.Contains(u.Host, ":") {
		host, _, err = net.SplitHostPort(u.Host)
	} else {
		host = u.Host
	}
	return "https://" + host + strings.TrimSuffix(u.Path, ".git"), nil
}

func webifyGitRemote(remoteUrl string) (string, error) {
	url := strings.TrimRight(remoteUrl[4:], ".git")
	url = strings.Replace(url, ":", "/", 1)
	return "https://" + url, nil
}

func webifyHttpRemote(remoteUrl string) (string, error) {
	if strings.HasSuffix(remoteUrl, ".git") {
		return strings.TrimSuffix(remoteUrl, ".git"), nil
	}
	return remoteUrl, nil
}
