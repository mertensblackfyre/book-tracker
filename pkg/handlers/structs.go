package handlers

type Users struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type Book struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Pages   string  `json:"pages"`
	Picture string  `json:"picture"`
	Prices  float64 `json:"prices"`
	Status  string  `json:"status"`
	UserID  string  `json:"user_id"`
}
