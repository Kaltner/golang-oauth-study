package oauth

const twitchAuthorizationURI = "https://id.twitch.tv/oauth2/authorize"
const twitchTokenURI = "https://id.twitch.tv/oauth2/token"
const DefaultTwitchScope = "user:bot"

func NewTwitch(clientID, clientSecret, redirectURI string) Oauth {
	return Oauth{
		clientID:         clientID,
		clientSecret:     clientSecret,
		redirectURI:      redirectURI,
		authorizationURI: twitchAuthorizationURI,
		tokenURI:         twitchTokenURI,
		provider:         TwitchProvider,
	}
}
