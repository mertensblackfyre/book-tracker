package pkg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func JWT(data string) string {
	secret := []byte(GetEnv("SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	if err != nil {
		log.Println(err)
	}

	return tokenString

}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	config := GoogleAuthConfig()
	url := config.AuthCodeURL("ran")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request) {

	state := r.URL.Query()["state"][0]

	if state != "ran" {
		fmt.Fprintln(w, "States dont match")
	}

	code := r.URL.Query()["code"][0]
	if len(code) == 0 {
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

	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	q := NewDB(db)
	q.AddUser(string(data))

	var user Users
	err = json.Unmarshal([]byte(data), &user)

	if err != nil {
		log.Println(err)
	}


	// Set token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   JWT(user.ID),
		Secure:  false,
		Path:    "/",
		Expires: time.Now().Add(30 * time.Minute),
	})

	http.Redirect(w, r, "/", 301)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "Token",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, &cookie)

	//http.Redirect(w, r, "/login", 301)
	fmt.Println(w, "Success")
}
