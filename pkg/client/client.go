package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiVersion struct {
	Version string `json:"version"`
	Date    string `json:"version_date"`
	Number  string `json:"api_version"`
}

type TpmClient struct {
	Client   *http.Client
	Server   string
	Username string
	Password string
}

func New(server string, username string, password string) TpmClient {
	client := TpmClient{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		Server:   server,
		Username: username,
		Password: password,
	}

	return client
}

func (client *TpmClient) get(url string, v interface{}) error {
	req, reqError := http.NewRequest(http.MethodGet, client.Server+url, nil)
	if reqError != nil {
		return reqError
	}

	req.SetBasicAuth(client.Username, client.Password)
	req.Header.Add("Content-Type", `application/json; charset=utf-8`)
	req.Header.Set("User-Agent", "nrocco/tpm")

	res, resError := client.Client.Do(req)
	if resError != nil {
		return resError
	}

	body, bodyError := ioutil.ReadAll(res.Body)
	if bodyError != nil {
		return bodyError
	}

	jsonError := json.Unmarshal(body, v)
	if jsonError != nil {
		return jsonError
	}

	return nil
}

func (client *TpmClient) Version() (*ApiVersion, error) {
	apiVersion := &ApiVersion{}

	err := client.get("/api/v4/version.json", apiVersion)
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}
