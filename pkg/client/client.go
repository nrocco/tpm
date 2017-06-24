package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

var (
	// Version holds the version number of the tpm rest api client library
	Version string
)

// TpmClient represents the client that handles all interaction with the server
type TpmClient struct {
	Client      *http.Client
	Server      string
	Username    string
	Password    string
	ContentType string
	UserAgent   string
}

// Count describes how many items and pages exist in a collection such as projects or passwords.
type Count struct {
	Items        int `json:"num_items"`
	Pages        int `json:"num_pages"`
	ItemsPerPage int `json:"num_items_per_page"`
}

// New creates a new instance of TpmClient which can then be used to make api calls
// It configures the underlying net/http.client with a sensible timeout.
func New(server string, username string, password string) TpmClient {
	client := TpmClient{
		Client: &http.Client{
			Timeout: time.Second * 3,
		},
		Server:      server,
		Username:    username,
		Password:    password,
		ContentType: "application/json; charset=utf-8",
		UserAgent:   fmt.Sprintf("tpm/%s (%s/%s)", Version, runtime.GOOS, runtime.GOARCH),
	}

	return client
}

func (client *TpmClient) request(method string, url string, v interface{}) error {
	req, reqError := http.NewRequest(method, client.Server+url, nil)
	if reqError != nil {
		return reqError
	}

	req.SetBasicAuth(client.Username, client.Password)
	req.Header.Add("Content-Type", client.ContentType)
	req.Header.Set("User-Agent", client.UserAgent)

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

func (client *TpmClient) put(url string) error {
	return client.request(http.MethodPut, url, nil)
}

func (client *TpmClient) get(url string, v interface{}) error {
	return client.request(http.MethodGet, url, v)
}
