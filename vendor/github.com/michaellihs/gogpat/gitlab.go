package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	logrus "github.com/sirupsen/logrus"
)

const gitlabSignInPath = "/users/sign_in"
const gitlabPersonalAccessTokenURLPath = "/profile/personal_access_tokens"

// GitLabTokenRequest describes the token Request
type GitLabTokenRequest struct {
	URL       string
	Username  string
	Password  string
	Scope     Scope
	Date      string
	TokenName string
}

// CreateToken creates the token for a given gitlab instance
func CreateToken(gr GitLabTokenRequest) (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
		Jar:     jar,
	}

	loginURL, err := url.Parse(gr.URL + gitlabSignInPath)

	if err != nil {
		return "", err
	}

	postURL, err := url.Parse(gr.URL + gitlabPersonalAccessTokenURLPath)

	if err != nil {
		return "", err
	}

	v := url.Values{}
	v, err = addExpiry(v, gr.Date)
	if err != nil {
		return "", err
	}

	responseBody, err := visitFirst(loginURL, client)
	if err != nil {
		return "", err
	}
	v = getCSRFTokenFromBody(responseBody, v)
	responseBody = login(loginURL, gr, client, v)
	responseBody, err = postToken(postURL, gr, client, getCSRFTokenFromBody(responseBody, v))
	if err != nil {
		return "", err
	}
	token, err := findPersonalAccessToken(responseBody)
	if err != nil {
		return "", err
	}
	return token, nil
}

func visitFirst(u *url.URL, client *http.Client) (htmlBody string, err error) {
	response, err := client.Get(u.String())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func login(u *url.URL, gr GitLabTokenRequest, client *http.Client, values url.Values) (htmlBody string) {
	values.Add("user[login]", gr.Username)
	values.Add("user[password]", gr.Password)
	values.Add("user[remember_me]", "0")
	values.Add("utf8", "✓")
	resp, err := client.PostForm(u.String(), values)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func postToken(u *url.URL, gr GitLabTokenRequest, client *http.Client, values url.Values) (string, error) {
	values.Add("utf8", "✓")
	values.Add("personal_access_token[name]", gr.TokenName)
	values = NewScope(values, gr.Scope)
	resp, err := client.PostForm(u.String(), values)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func addExpiry(values url.Values, date string) (url.Values, error) {
	if date == "" {
		return values, nil
	}
	dateparts := strings.Split(date, "-")
	if len(dateparts) < 3 {
		return nil, errors.New("Date is too short it should be formatted like this 2017-12-03")
	}

	if len(dateparts[0]) < 4 {
		return nil, errors.New("Year is too short")
	}

	if len(dateparts[1]) < 2 {
		return nil, errors.New("Month is too short needs to be with zero digest")
	}
	if len(dateparts[2]) < 2 {
		return nil, errors.New("Day is too short needs to be with zero digest")
	}
	values.Add("personal_access_token[expires_at]", date)
	return values, nil
}

func getCSRFTokenFromBody(htmlBody string, values url.Values) url.Values {
	csrf := findCSRFToken(htmlBody)
	values.Add(csrf["csrf-param"], csrf["csrf-token"])
	return values
}

func findCSRFToken(htmlBody string) map[string]string {
	doc := soup.HTMLParse(htmlBody)
	csrfParam := doc.Find("meta", "name", "csrf-param").Attrs()["content"]
	csrfToken := doc.Find("meta", "name", "csrf-token").Attrs()["content"]
	m := make(map[string]string)
	m["csrf-param"] = csrfParam
	m["csrf-token"] = csrfToken
	return m
}

func findPersonalAccessToken(htmlBody string) (string, error) {
	doc := soup.HTMLParse(htmlBody)
	found := doc.Find("input", "id", "created-personal-access-token")
	if found.Pointer == nil {
		errorgrep := doc.Find("div", "id", "error_explanation")
		if errorgrep.Pointer == nil {
			return "", errors.New("Could find any error field")
		}
		return "", fmt.Errorf("%s", errorgrep.Pointer.LastChild.LastChild.LastChild.Data)
	}
	if val, ok := found.Attrs()["value"]; ok {
		return val, nil
	}
	return "", errors.New("could found something")
}
