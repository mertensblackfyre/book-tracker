package handlers

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Locale  string `json:"locale"`
}
