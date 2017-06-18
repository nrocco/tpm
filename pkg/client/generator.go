package client

type GeneratedPassword struct {
	Value string `json:"password"`
}

func (client *TpmClient) GeneratePassword() (*GeneratedPassword, error) {
	password := &GeneratedPassword{}

	err := client.get("/api/v4/generate_password.json", password)
	if err != nil {
		return nil, err
	}

	return password, nil
}
