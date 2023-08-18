package pkg

type Users struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type Book struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Pages   int     `json:"pages"`
	Picture string  `json:"picture"`
	Prices  float64 `json:"prices"`
	Status  string  `json:"status"`
	UserID  int     `json:"user_id"`
}

// Error implements error.
func (Book) Error() string {
	panic("unimplemented")
}
