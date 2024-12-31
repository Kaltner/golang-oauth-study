# Golang Oauth Study Repo

This is a personal repo for me to study how to implement different Oauth2 flows using Golang.

## Requiremnts
- Golang

## Setup
1. Clone this repo
2. Copy the `.env.example` to a `.env` file and add the environment variables
3. Run `go install`
4. Start the application with the `go run .` command
5. Go to [http:localhost:8080/oauth/authorize](http:localhost:8080/oauth/authorize) endpoint in your browser.

## Objectives
- Implement a Github Oauth 2 flow using Golang Standard libraries
- Add unit tests for all implementation
- Implement the Twitch Oauth2 flow.

### Work in progress
- [x] Twitch Oauth2 Implementation
- [] GitHub unit tests
- [] Error handling from Oauth Servers responses
- [] Refresh token flow     

## Study material
[Oauth Webside](https://www.oauth.com/)
[Github Oauth Docs](https://docs.github.com/en/apps/oauth-apps/using-oauth-apps/installing-an-oauth-app-in-your-personal-account)

