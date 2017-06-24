package client

// APIVersion represents the version of the team password manager server
type APIVersion struct {
	Version string `json:"version"`
	Date    string `json:"version_date"`
	Number  string `json:"api_version"`
}

// Version fetches the team password manager server verison.
// It returns a struct APIVersion and any error encountered
func (client *TpmClient) Version() (*APIVersion, error) {
	apiVersion := &APIVersion{}

	err := client.get("/api/v4/version.json", apiVersion)
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}
