package pkg

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/auth/google" || strings.Contains(r.URL.Path, "/auth/callback") || strings.Contains(r.URL.Path, "https://www.googleapis.com/oauth2/v2/userinfo?access_token=") {
			cookie, _ := r.Cookie("Token")

			if cookie != nil {
				log.Println("Its not nil")
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, r)

		} else {
			cookie, err := r.Cookie("Token")

			if err != nil {
				log.Println(err)
				return
			}
			// Verify and parse token
			tknStr := cookie.Value
			token, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				next.ServeHTTP(w, r)
				return []byte(GetEnv("SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["id"])
			} else {
				fmt.Println(err)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
