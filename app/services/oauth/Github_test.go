package oauth

import (
	"net/url"
	"strings"
	"testing"
)

func TestBuildAuthorizeURL(t *testing.T) {
	clientID := "1234"
	redirectURI := "http://localhost"

	githubOauth := NewGithub(clientID, redirectURI)

	actualAuthURL, err := githubOauth.buildAuthorizeURL("0")
	if err != nil {
		t.Errorf("err when generating the Authorization URL: %+v", err)
	}
	if !testQueryParams(t, *actualAuthURL, clientID, redirectURI) {
		t.Fail()
	}
}

func testQueryParams(t *testing.T, actualAuthURL url.URL, clientID, redirectURI string) bool {
	result := true
	if res := actualAuthURL.String(); !strings.Contains(res, authorizeUrl) {
		t.Logf("Base URL is not for github authorize \nGot: %s", res)
		result = false
	}
	queryParams := actualAuthURL.Query()
	if res := queryParams.Get("client_id"); res != clientID {
		t.Logf("client_id does not match what was expected.\nGot: %s\nExp: %s", res, clientID)
		result = false
	}
	if res := queryParams.Get("redirect_uri"); res != redirectURI {
		t.Logf("redirect_uri does not match what was expected.\nGot: %s\nExp: %s", res, redirectURI)
		result = false
	}
	return result
}
