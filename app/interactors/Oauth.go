package interactors

import (
	"errors"

	"github.com/Kaltner/oauth_test/app/services/oauth"
)

type OauthInteractor struct {
	GithubService oauth.Oauth
	TwitchService oauth.Oauth
	StateManager  *oauth.StateManager
}

func NewOauthInteractor(githubService, twitchService oauth.Oauth) OauthInteractor {
	return OauthInteractor{
		GithubService: githubService,
		TwitchService: twitchService,
		StateManager:  oauth.NewStateManager(),
	}
}

func (o OauthInteractor) Authorize(provider string) (string, error) {
	p := oauth.OauthProvider(provider)
	if p == "" {
		return "", errors.New("provider is not supported")
	}

	state := o.StateManager.GenerateState(p)
	switch p {
	case oauth.GithubProvider:
		return o.GithubService.Authorize(state, oauth.DefaultGithubScope)
	case oauth.TwitchProvider:
		return o.TwitchService.Authorize(state, oauth.DefaultTwitchScope)
	}

	return "", errors.New("unexpected error when getting the provider")
}

func (o OauthInteractor) Callback(code, state string) (string, error) {
	provider, ok := o.StateManager.FindState(state)
	if !ok {
		return "", errors.New("state could not be found")
	}
	switch provider {
	case oauth.GithubProvider:
		return o.GithubService.Callback(code)
	case oauth.TwitchProvider:
		return o.TwitchService.Callback(code)
	}

	return "", errors.New("provider is not supported")
}
