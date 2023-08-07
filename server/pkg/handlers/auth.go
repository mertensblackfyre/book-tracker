package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	pgx "github.com/jackc/pgx/v5"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	config := GoogleAuthConfig()
	url := config.AuthCodeURL("ran")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request) {

	var tx pgx.Tx

	state := r.URL.Query()["state"][0]

	if state != "ran" {
		fmt.Fprintln(w, "States dont match")
		return
	}

	code := r.URL.Query()["code"][0]
	token, err := GoogleAuthConfig().Exchange(context.Background(), code)

	if err != nil {
		fmt.Fprintln(w, "Code-Token exchane failed")
	}

	response, err := http.Get(OauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		fmt.Fprintln(w, "Code-Token exchane failed")
	}

	data, err := ioutil.ReadAll(response.Body)

	fmt.Println(w, string(data))

	AddUser(context.Background(), tx, data)

}
