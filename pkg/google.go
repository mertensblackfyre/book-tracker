package pkg

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleAuthConfig() *oauth2.Config {

	var url string

	if GetEnv("NODE_ENV") == "production" {
		url = "https://backend-t76n.onrender.com"
	}

	if GetEnv("NODE_ENV") == "development" {
		url = "http://localhost:5000"
	}
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf := &oauth2.Config{
		ClientID:     GetEnv("GOOGLE_CLIENT"),
		ClientSecret: GetEnv("GOOGLE_SECRET"),
		RedirectURL:  url + "/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}
