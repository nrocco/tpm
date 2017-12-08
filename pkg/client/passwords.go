package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type CustomField struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Data  string `json:"data"`
}

type Password struct {
	ID              string  `json:"id"`
	AccessInfo      string  `json:"access_info"`
	Archived        bool    `json:"archived"`
	CreatedBy       User    `json:"created_by"`
	CreatedOn       string  `json:"created_on"`
	Email           string  `json:"email"`
	ExpiryDate      string  `json:"expiry_date"`
	ExpiryStatus    int     `json:"expiry_status"`
	ExternalSharing bool    `json:"external_sharing"`
	ExternalURL     string  `json:"external_url"`
	Favorite        bool    `json:"favorite"`
	Locked          bool    `json:"locked"`
	ManagedBy       User    `json:"managed_by"`
	Name            string  `json:"name"`
	Notes           string  `json:"notes"`
	NumFiles        string  `json:"num_files"`
	Password        string  `json:"password"`
	Project         Project `json:"project"`
	Tags            string  `json:"tags"`
	UpdatedBy       User    `json:"updated_by"`
	UpdatedOn       string  `json:"updated_on"`
	Username        string  `json:"username"`

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

// UnmarshalJSON handles id's from team password manager as int and string
func (v *Password) UnmarshalJSON(data []byte) error {
	type Password2 Password
	x := struct {
		Password2
		ID       json.Number `json:"id"`
		NumFiles json.Number `json:"num_files"`
	}{Password2: Password2(*v)}

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	*v = Password(x.Password2)
	v.ID = x.ID.String()
	v.NumFiles = x.NumFiles.String()
	return nil
}

type Passwords []Password

func (v *Passwords) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, r := range raw {
		var password Password
		err := json.Unmarshal(r, &password)
		if err != nil {
			return err
		}
		*v = append(*v, password)
	}

	return nil
}

func (client *TpmClient) PasswordList(search string) (Passwords, error) {
	var baseURL string
	var err error

	if search != "" {
		baseURL = fmt.Sprintf("/api/v4/passwords/search/%s", url.PathEscape(search))
	} else {
		baseURL = "/api/v4/passwords"
	}

	count := Count{}
	err = client.get(baseURL+"/count.json", &count)
	if err != nil {
		return nil, err
	}

	page := 1
	passwords := Passwords{}

	for page <= count.Pages {
		err = client.get(fmt.Sprintf("%s/page/%d.json", baseURL, page), &passwords)
		if err != nil {
			return nil, err
		}

		page++
	}

	return passwords, nil
}

func (client *TpmClient) PasswordGet(id string) (*Password, error) {
	password := &Password{}

	err := client.get("/api/v4/passwords/"+id+".json", password)
	if err != nil {
		return nil, err
	}

	return password, nil
}
