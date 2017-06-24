package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type TpmClient struct {
	Client      *http.Client
	Server      string
	Username    string
	Password    string
	ContentType string
	UserAgent   string
}

type Count struct {
	Items        int `json:"num_items"`
	Pages        int `json:"num_pages"`
	ItemsPerPage int `json:"num_items_per_page"`
}

func New(server string, username string, password string) TpmClient {
	client := TpmClient{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		Server:      server,
		Username:    username,
		Password:    password,
		ContentType: "application/json; charset=utf-8",
		UserAgent:   "nrocco/tpm",
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
