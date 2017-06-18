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

func (client *TpmClient) Version() (*ApiVersion, error) {
	url := client.Server + "/api/v4/version.json"

	req, reqError := http.NewRequest(http.MethodGet, url, nil)
	if reqError != nil {
		return nil, reqError
	}

	req.SetBasicAuth(client.Username, client.Password)
	req.Header.Add("Content-Type", `application/json; charset=utf-8`)
	req.Header.Set("User-Agent", "nrocco/tpm")

	res, resError := client.Client.Do(req)
	if resError != nil {
		return nil, resError
	}

	body, bodyError := ioutil.ReadAll(res.Body)
	if bodyError != nil {
		return nil, bodyError
	}

	apiVersion := &ApiVersion{}

	jsonError := json.Unmarshal(body, apiVersion)
	if jsonError != nil {
		return nil, jsonError
	}

	return apiVersion, nil
}

func (client *TpmClient) List(search string) (Passwords, error) {
	url := client.Server + "/api/v4/passwords/search/" + search + "/page/1.json"

	req, reqError := http.NewRequest(http.MethodGet, url, nil)
	if reqError != nil {
		return nil, reqError
	}

	req.SetBasicAuth(client.Username, client.Password)
	req.Header.Add("Content-Type", `application/json; charset=utf-8`)
	req.Header.Set("User-Agent", "nrocco/tpm")

	res, resError := client.Client.Do(req)
	if resError != nil {
		return nil, resError
	}

	body, bodyError := ioutil.ReadAll(res.Body)
	if bodyError != nil {
		return nil, bodyError
	}

	passwords := Passwords{}

	jsonError := json.Unmarshal(body, &passwords)
	if jsonError != nil {
		return nil, jsonError
	}

	return passwords, nil
}

func (client *TpmClient) Get(id string) (*Password, error) {
	url := client.Server + "/api/v4/passwords/" + id + ".json"

	req, reqError := http.NewRequest(http.MethodGet, url, nil)
	if reqError != nil {
		return nil, reqError
	}

	req.SetBasicAuth(client.Username, client.Password)
	req.Header.Add("Content-Type", `application/json; charset=utf-8`)
	req.Header.Set("User-Agent", "nrocco/tpm")

	res, resError := client.Client.Do(req)
	if resError != nil {
		return nil, resError
	}

	body, bodyError := ioutil.ReadAll(res.Body)
	if bodyError != nil {
		return nil, bodyError
	}

	password := &Password{}

	jsonError := json.Unmarshal(body, password)
	if jsonError != nil {
		return nil, jsonError
	}

	return password, nil
}
