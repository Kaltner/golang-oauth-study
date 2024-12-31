package app

import (
	"net/http"
	"os"

	"github.com/Kaltner/oauth_test/app/handlers"
	"github.com/Kaltner/oauth_test/app/interactors"
	"github.com/Kaltner/oauth_test/app/services/oauth"
)

func ServeMultiplexer() *http.ServeMux {
	githubClientID, _ := os.LookupEnv("GITHUB_CLIENT_ID")
	githubClientSecret, _ := os.LookupEnv("GITHUB_CLIENT_SECRET")
	twitchClientID, _ := os.LookupEnv("TWITCH_CLIENT_ID")
	twitchClientSecret, _ := os.LookupEnv("TWITCH_CLIENT_SECRET")
	redirectURI, _ := os.LookupEnv("OAUTH_REDIRECT_URL")
	githubOauthService := oauth.NewGithub(githubClientID, githubClientSecret, redirectURI)
	twitchOauthService := oauth.NewTwitch(twitchClientID, twitchClientSecret, redirectURI)

	oauthInteractor := interactors.NewOauthInteractor(githubOauthService, twitchOauthService)

	oauthHandler := handlers.NewOauthHandler(oauthInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/authorize/", oauthHandler.Authorize)
	mux.HandleFunc("/oauth/callback/", oauthHandler.Callback)

	return mux
}
