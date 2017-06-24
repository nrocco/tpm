package client

// User represents a team password manager user
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
	Role         string `json:"role"`
}
