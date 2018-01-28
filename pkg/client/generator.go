package client

// GeneratedPassword represents a API response with the generated password
type GeneratedPassword struct {
	Value string `json:"password"`
}

// GeneratePassword uses the team password manager API to generate a random
// password
func (client *TpmClient) GeneratePassword() (*GeneratedPassword, error) {
	password := &GeneratedPassword{}

	err := client.get("/api/v4/generate_password.json", password)
	if err != nil {
		return nil, err
	}

	return password, nil
}
