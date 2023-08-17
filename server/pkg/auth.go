package pkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	config := GoogleAuthConfig()
	url := config.AuthCodeURL("ran")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request)[]byte {


	state := r.URL.Query()["state"][0]

	if state != "ran" {
		fmt.Fprintln(w, "States dont match")
		return nil
	}

	code := r.URL.Query()["code"][0]
    if len(code) == 0{
        log.Fatalln("Code is 0")
    }

	token, err := GoogleAuthConfig().Exchange(context.Background(), code)

	if err != nil {
		fmt.Fprintln(w, "Code-Token exchane failed")
	}

	response, err := http.Get(OauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		fmt.Fprintln(w, "Code-Token exchane failed")
	}

	data, err := ioutil.ReadAll(response.Body)


	if err != nil {
		fmt.Fprintln(w, err)
	}

    http.Redirect(w,r,"/",200)
    return data
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	Manager.Destroy(r.Context())

	fmt.Println(w, "Success")
	return nil
}
