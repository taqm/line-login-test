package line

import (
	"os"

	"golang.org/x/oauth2"
)

var oauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("LINE_CLIENT_ID"),
	ClientSecret: os.Getenv("LINE_CLIENT_SECRET"),
	Scopes:       []string{"profile"},
	RedirectURL:  os.Getenv("LINE_REDIRECT_URL"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
		TokenURL: "https://api.line.me/oauth2/v2.1/token",
	},
}

func AuthCodeURL() string {
	return oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
}
