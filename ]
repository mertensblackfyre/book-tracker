package pkg

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/auth/google" || strings.Contains(r.URL.Path, "/auth/callback") || strings.Contains(r.URL.Path, "https://www.googleapis.com/oauth2/v2/userinfo?access_token=") {

			next.ServeHTTP(w, r)

		} else {

			cookie, err := r.Cookie("Token")

			if err != nil {
				log.Println(err)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorized"))
				return
			}

			// Verify and parse token
			tknStr := cookie.Value
			token, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(GetEnv("SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "data", claims["id"])
				r = r.WithContext(ctx)

			} else {
				fmt.Println(err)
				//w.WriteHeader(401)
				//	w.Write([]byte("Unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		}
	})
}
