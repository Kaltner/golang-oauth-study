package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Github struct {
	clientID    string
	redirectURI string
	stateList   map[string]bool
}

type gitHubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

const authorizeUrl = "https://github.com/login/oauth/authorize"
const tokenUrl = "https://github.com/login/oauth/access_token"

func NewGithub(clientID, redirectURI string) *Github {
	return &Github{
		clientID:    clientID,
		redirectURI: redirectURI,
		stateList:   make(map[string]bool),
	}
}
func (g *Github) Authorize() (string, error) {
	state := g.generateState()

	authorizeUrl, err := g.buildAuthorizeURL(state)
	if err != nil {
		return "", nil
	}
	return authorizeUrl.String(), nil
}

func (g *Github) Callback(code, state string) (string, error) {
	if !g.findState(state) {
		return "", errors.New("could not find the state of the request")
	}
	g.deleteState(state)
	accessToken, err := g.fecthAccessToken(code)
	if err != nil {
		return "", fmt.Errorf("failed to ge tthe accessToken: %s", err)
	}
	return accessToken, nil
}

func (g *Github) generateState() string {
	fmt.Printf("%+v\n", g.stateList)
	state := strings.ReplaceAll(uuid.NewString(), "-", "")
	g.stateList[state] = true
	fmt.Printf("%+v\n", g.stateList)
	return state
}

func (g *Github) buildAuthorizeURL(state string) (*url.URL, error) {
	responseType := "code"
	scope := "user public_repo"

	callbackUrl := fmt.Sprintf("%s/%s", g.redirectURI, "/callback")
	queryParams := url.Values{}
	queryParams.Add("response_type", responseType)
	queryParams.Add("client_id", g.clientID)
	queryParams.Add("redirect_uri", callbackUrl)
	queryParams.Add("scope", scope)
	queryParams.Add("state", state)

	fmt.Printf("%+v\n", queryParams)
	url, err := url.Parse(fmt.Sprintf("%s?%s", authorizeUrl, queryParams.Encode()))
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (g *Github) findState(state string) bool {
	_, ok := g.stateList[state]
	return ok
}

func (g *Github) deleteState(state string) {
	delete(g.stateList, state)
}

func (g *Github) fecthAccessToken(code string) (string, error) {
	grantType := "authorization_code"
	githubClientSecret, _ := os.LookupEnv("GITHUB_CLIENT_SECRET")

	values := url.Values{}
	values.Add("grant_type", grantType)
	values.Add("client_id", g.clientID)
	values.Add("client_secret", githubClientSecret)
	values.Add("redirect_uri", g.redirectURI)
	values.Add("code", code)

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(res.Body)
	var accessTokenResponse gitHubAccessTokenResponse
	err = decoder.Decode(&accessTokenResponse)
	if err != nil {
		return "", fmt.Errorf("failed to decode the JSON response: %q", err)
	}
	return accessTokenResponse.AccessToken, nil
}
