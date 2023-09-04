package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/google" || strings.Contains(r.URL.Path, "/auth/callback") || strings.Contains(r.URL.Path, "https://www.googleapis.com/oauth2/v2/userinfo?access_token=") {
			next.ServeHTTP(w, r)
		} else {
			msg := Manager.Get(r.Context(), "name")
			if msg == nil {
				fmt.Fprintln(w, "Unauthrized")
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
