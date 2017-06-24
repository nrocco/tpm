package client

type APIVersion struct {
	Version string `json:"version"`
	Date    string `json:"version_date"`
	Number  string `json:"api_version"`
}

func (client *TpmClient) Version() (*APIVersion, error) {
	apiVersion := &APIVersion{}

	err := client.get("/api/v4/version.json", apiVersion)
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}
