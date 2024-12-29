package app

import (
	"net/http"
	"os"

	"github.com/Kaltner/oauth_test/app/handlers"
	"github.com/Kaltner/oauth_test/app/services/oauth"
)

func ServeMultiplexer() *http.ServeMux {
	clientID, _ := os.LookupEnv("GITHUB_CLIENT_ID")
	redirectURI, _ := os.LookupEnv("OAUTH_REDIRECT_URL")
	githubOauthService := oauth.NewGithub(clientID, redirectURI)

	oauthHandler := handlers.NewOauthHandler(githubOauthService)

	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/authorize", oauthHandler.Authorize)
	mux.HandleFunc("/oauth/callback", oauthHandler.Callback)

	return mux
}
