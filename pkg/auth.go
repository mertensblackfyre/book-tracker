package pkg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func Login(w http.ResponseWriter, r *http.Request) {


	fmt.Println("SETT cookies")
	str := JWT("1")

    data := Users{
        Name: "Cool",
		Picture: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.jryuUgIHWL-1FVD2ww8oWgHaHa%26pid%3DApi&f=1&ipt=0fd7506ba6c5071ef907b0ea3537ca197dc8b919bd03cf1cb172b9ebd91e4768&ipo=images",
	}

	// fmt.Println(str)


	// Set token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   str,
		Secure:  false,
		Path:    "/",
		Expires: time.Now().Add(30 * time.Minute),
	})

    JSONWritter(w,200,data)
	http.Redirect(w, r, "/mybooks", 302)
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
		Secure:  true,
		Path:    "/",
		Expires: time.Now().Add(30 * time.Minute),
	})

	JSONWritter(w, 200, user)

	http.Redirect(w, r, "/mybooks", 302)

}

func Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   "",
		Secure:  true,
		Path:    "/",
		Expires: time.Unix(0, 0),
	})

	fmt.Println(w, "Success")
	http.Redirect(w, r, "/login", 302)
}
