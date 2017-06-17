package client

import (
    "time"
    "io/ioutil"
    "log"
    "encoding/json"
    "net/http"
)

type SimplePassword struct {
    Id int `json:"id"`
    AccessInfo string`json:"access_info"`
    Email string`json:"updated_on"`
    Name string`json:"name"`
    UpdatedOn string`json:"updated_on"`
    Username string`json:"username"`
    Tags string`json:"tags"`
}

type CustomField struct {
    Type string `json:"type"`
    Label string `json:"label"`
    Data string `json:"data"`
}

type DetailedPassword struct {
    Id int `json:"id"`
    AccessInfo string`json:"access_info"`
    Email string`json:"updated_on"`
    Name string`json:"name"`
    Notes string`"notes"`
    Password string`"password"`
    Tags string`json:"tags"`
    UpdatedOn string`json:"updated_on"`
    Username string`json:"username"`

    CustomField1 CustomField`json:"custom_field1"`
    CustomField2 CustomField`json:"custom_field2"`
    CustomField3 CustomField`json:"custom_field3"`
    CustomField4 CustomField`json:"custom_field4"`
    CustomField5 CustomField`json:"custom_field5"`
    CustomField6 CustomField`json:"custom_field6"`
    CustomField7 CustomField`json:"custom_field7"`
    CustomField8 CustomField`json:"custom_field8"`
    CustomField9 CustomField`json:"custom_field9"`
}

type SimplePasswords []SimplePassword

type Client struct {
    Client *http.Client
    Server string
    Username string
    Password string
}

func New(server string, username string, password string) Client {
    client := Client{
        Client: &http.Client{
            Timeout: time.Second * 10,
        },
        Server: server,
        Username: username,
        Password: password,
    }

    return client
}

func (client *Client) List(search string) (SimplePasswords, error) {
    url := client.Server + "/api/v4/passwords/search/" + search + "/page/1.json"

    req, reqError := http.NewRequest(http.MethodGet, url, nil)
    if reqError != nil {
        log.Fatal(reqError)
        return nil, reqError
    }

    req.SetBasicAuth(client.Username, client.Password)
    req.Header.Add("Content-Type", `application/json; charset=utf-8`)
    req.Header.Set("User-Agent", "nrocco/tpm")

    res, resError := client.Client.Do(req)
    if resError != nil {
        log.Fatal(resError)
        return nil, resError
    }

    body, bodyError := ioutil.ReadAll(res.Body)
    if bodyError != nil {
        log.Fatal(bodyError)
        return nil, bodyError
    }

    passwords := SimplePasswords{}

    jsonError := json.Unmarshal(body, &passwords)
    if jsonError != nil {
        log.Fatal(jsonError)
        return nil, jsonError
    }

    return passwords, nil
}

func (client *Client) Get(id string) (*DetailedPassword, error) {
    url := client.Server + "/api/v4/passwords/" + id + ".json"

    req, reqError := http.NewRequest(http.MethodGet, url, nil)
    if reqError != nil {
        log.Fatal(reqError)
        return nil, reqError
    }

    req.SetBasicAuth(client.Username, client.Password)
    req.Header.Add("Content-Type", `application/json; charset=utf-8`)
    req.Header.Set("User-Agent", "nrocco/tpm")

    res, resError := client.Client.Do(req)
    if resError != nil {
        log.Fatal(resError)
        return nil, resError
    }

    body, bodyError := ioutil.ReadAll(res.Body)
    if bodyError != nil {
        log.Fatal(bodyError)
        return nil, bodyError
    }

    detailedPassword := &DetailedPassword{}

    jsonError := json.Unmarshal(body, detailedPassword)
    if jsonError != nil {
        log.Fatal(jsonError)
        return nil, jsonError
    }

    return detailedPassword, nil
}
