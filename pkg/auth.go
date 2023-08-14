package pkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	pgx "github.com/jackc/pgx/v5"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	config := GoogleAuthConfig()
	url := config.AuthCodeURL("ran")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request) {

	// var tx pgx.Tx
	conn := DBConfig()

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

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return AddUser(context.Background(), tx, data)
	})

	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	Manager.Destroy(r.Context())

	fmt.Println(w, "Success")
	return nil
}
