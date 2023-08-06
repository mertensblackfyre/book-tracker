package handlers

import (
	"context"
	"fmt"
	"net/http"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	config := GoogleAuthConfig()
	url := config.AuthCodeURL("ran")

	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallBack(res http.ResponseWriter, req *http.Request) {

	state := req.URL.Query()["state"][0]

	if state != "ran" {
		fmt.Fprintln(res, "States dont match")
		return
	}

	code := req.URL.Query()["code"][0]
	token, err := GoogleAuthConfig().Exchange(context.Background(), code)

	if err != nil {
		fmt.Fprintln(res, "Code-Token exchane failed")
	}

	
}
