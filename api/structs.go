package api

// Response struct
type Response struct {
	Status     string
	StatusCode int
	Body       string
}

// Configuration struct
type Configuration struct {
	Host  string `json:"hostname"`
	Token string `json:"token"`
}

// Network strcut
type Network struct {
	ID string `json:"network_id"`
}

// Member struct
type Member struct {
	ID string `json:"member_id"`
}

// User struct
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Disabled bool   `json:"disabled"`
}
