package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"

)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/logout" || r.URL.Path == "/auth/google" || strings.Contains(r.URL.Path, "/auth/callback") || strings.Contains(r.URL.Path, "https://www.googleapis.com/oauth2/v2/userinfo?access_token=") {
			next.ServeHTTP(w, r)
		} else {
			cookie, err := r.Cookie("Authorization")
			if err != nil {
				fmt.Println(err)
			}

			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return GetEnv("SECRET"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["id"], claims["nbf"])
			} else {
				fmt.Println(err)
                return
			}
			next.ServeHTTP(w, r)
		}
	})
}
