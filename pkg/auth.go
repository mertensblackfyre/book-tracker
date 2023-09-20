package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "time"
 "github.com/golang-jwt/jwt"
)


const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func JWT(data string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  22,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(GetEnv("SECRET"))

	if err != nil {
		log.Println(err)
	}

    fmt.Println(tokenString)
	return tokenString

}

func Login(w http.ResponseWriter, r *http.Request) {

    str := JWT("ss")
    fmt.Println(str)
	cookie := http.Cookie{
		Name:     "Token",
		Value:   str,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Domain:   "localhost",
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
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
	fmt.Println(data)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	NewDB(db)

	//q.AddUser(string(data))

	// Set a cookie

	http.Redirect(w, r, "/", 200)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Success")
}
