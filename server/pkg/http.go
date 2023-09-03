package pkg

import (
	"fmt"
	"net/http"
)

func Handler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			fmt.Fprintf(w, "Error", err)
			return
		}
		return
	}
}

func Han(v any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
