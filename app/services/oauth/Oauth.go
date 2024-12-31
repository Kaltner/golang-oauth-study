package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Oauth struct {
	clientID         string
	clientSecret     string
	redirectURI      string
	authorizationURI string
	tokenURI         string
	provider         OauthProvider
}

type OauthProvider string

const (
	GithubProvider OauthProvider = "github"
	TwitchProvider OauthProvider = "twitch"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (o *Oauth) Authorize(state, scope string) (string, error) {
	authorizeUrl, err := o.buildAuthorizeURL(state, scope)
	if err != nil {
		return "", nil
	}
	return authorizeUrl.String(), nil
}

func (o *Oauth) Callback(code string) (string, error) {
	accessToken, err := o.fecthAccessToken(code)
	if err != nil {
		return "", fmt.Errorf("failed to ge the accessToken: %s", err)
	}
	return accessToken, nil
}

func (o *Oauth) buildAuthorizeURL(state, scope string) (*url.URL, error) {
	responseType := "code"

	callbackUrl := fmt.Sprintf("%s%s", o.redirectURI, "callback/")
	queryParams := url.Values{}
	queryParams.Add("response_type", responseType)
	queryParams.Add("client_id", o.clientID)
	queryParams.Add("redirect_uri", callbackUrl)
	queryParams.Add("scope", scope)
	queryParams.Add("state", state)

	url, err := url.Parse(fmt.Sprintf("%s?%s", o.authorizationURI, queryParams.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create the Authorization URL for %s provider: %q", o.provider, err)
	}

	return url, nil
}

func (o *Oauth) fecthAccessToken(code string) (string, error) {
	grantType := "authorization_code"

	values := url.Values{}
	values.Add("grant_type", grantType)
	values.Add("client_id", o.clientID)
	values.Add("client_secret", o.clientSecret)
	values.Add("redirect_uri", o.redirectURI)
	values.Add("code", code)

	req, err := http.NewRequest("POST", o.tokenURI, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(resp.Body)
	var accessTokenResponse accessTokenResponse
	err = decoder.Decode(&accessTokenResponse)
	if err != nil {
		return "", fmt.Errorf("failed to decode the JSON response: %q", err)
	}
	return accessTokenResponse.AccessToken, nil

}
