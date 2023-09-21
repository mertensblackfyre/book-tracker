package pkg

type Users struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Created_at    string `json:"created_at"`
}

type Book struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Status     string  `json:"status"`
	Pages      int     `json:"pages"`
	Picture    string  `json:"picture"`
	Prices     float64 `json:"prices"`
	UserID     string  `json:"user_id"`
	Created_at string  `json:"created_at"`
}
