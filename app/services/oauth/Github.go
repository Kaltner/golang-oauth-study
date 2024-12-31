package oauth

const githubAuthorizationURI = "https://github.com/login/oauth/authorize"
const githubTokenURI = "https://github.com/login/oauth/access_token"
const DefaultGithubScope = "read:user"

func NewGithub(clientID, clientSecret, redirectURI string) Oauth {
	return Oauth{
		clientID:         clientID,
		clientSecret:     clientSecret,
		redirectURI:      redirectURI,
		authorizationURI: githubAuthorizationURI,
		tokenURI:         githubTokenURI,
		provider:         GithubProvider,
	}
}
