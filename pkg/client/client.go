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

type SimplePassword struct {
	Id         int    `json:"id"`
	AccessInfo string `json:"access_info"`
	Email      string `json:"updated_on"`
	Name       string `json:"name"`
	UpdatedOn  string `json:"updated_on"`
	Username   string `json:"username"`
	Tags       string `json:"tags"`
}

type CustomField struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Data  string `json:"data"`
}

type DetailedPassword struct {
	Id         int    `json:"id"`
	AccessInfo string `json:"access_info"`
	Email      string `json:"updated_on"`
	Name       string `json:"name"`
	Notes      string `"notes"`
	Password   string `"password"`
	Tags       string `json:"tags"`
	UpdatedOn  string `json:"updated_on"`
	Username   string `json:"username"`

	CustomField1 CustomField `json:"custom_field1"`
	CustomField2 CustomField `json:"custom_field2"`
	CustomField3 CustomField `json:"custom_field3"`
	CustomField4 CustomField `json:"custom_field4"`
	CustomField5 CustomField `json:"custom_field5"`
	CustomField6 CustomField `json:"custom_field6"`
	CustomField7 CustomField `json:"custom_field7"`
	CustomField8 CustomField `json:"custom_field8"`
	CustomField9 CustomField `json:"custom_field9"`
}

type SimplePasswords []SimplePassword

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

func (client *TpmClient) List(search string) (SimplePasswords, error) {
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

	passwords := SimplePasswords{}

	jsonError := json.Unmarshal(body, &passwords)
	if jsonError != nil {
		return nil, jsonError
	}

	return passwords, nil
}

func (client *TpmClient) Get(id string) (*DetailedPassword, error) {
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

	detailedPassword := &DetailedPassword{}

	jsonError := json.Unmarshal(body, detailedPassword)
	if jsonError != nil {
		return nil, jsonError
	}

	return detailedPassword, nil
}
