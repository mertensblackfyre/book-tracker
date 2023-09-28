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
	UserID     string  `json:"user_id"`
	Status     string  `json:"status"`
	Pages      int     `json:"pages"`
	Picture    string  `json:"picture"`
	Prices     float64 `json:"prices"`
	Started_at string  `json:"started_at"`
	Created_at string  `json:"created_at"`
}
