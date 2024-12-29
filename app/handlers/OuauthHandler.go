package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Kaltner/oauth_test/app/services/oauth"
)

type OauthHandler struct {
	GithubService *oauth.Github
}

func NewOauthHandler(githubService *oauth.Github) OauthHandler {
	return OauthHandler{
		GithubService: githubService,
	}
}

func (o OauthHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	authURL, err := o.GithubService.Authorize()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	http.Redirect(w, r, authURL, http.StatusFound)
}

func (o OauthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	state, code, err := o.checkCallbackRequestQueryParams(queryParams)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}

	accessToken, err := o.GithubService.Callback(code, state)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Access Token: %s", accessToken)
	w.Write([]byte(accessToken))
}

func (o OauthHandler) checkCallbackRequestQueryParams(query url.Values) (string, string, error) {
	state, ok := query["state"]
	if !ok || len(state) == 0 {
		return "", "", errors.New("state query string is not set")
	}
	code, ok := query["code"]
	if !ok || len(code) == 0 {
		return "", "", errors.New("code query string is not set")
	}

	return state[0], code[0], nil
}
