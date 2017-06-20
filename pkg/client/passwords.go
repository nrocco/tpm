package client

type CustomField struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Data  string `json:"data"`
}

type Password struct {
	ID              int     `json:"id"`
	Project         Project `json:"project"`
	AccessInfo      string  `json:"access_info"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	Notes           string  `json:"notes"`
	Password        string  `json:"password"`
	Tags            string  `json:"tags"`
	CreatedOn       string  `json:"created_on"`
	UpdatedOn       string  `json:"updated_on"`
	Username        string  `json:"username"`
	ExpiryDate      string  `json:"expiry_date"`
	ExpiryStatus    int     `json:"expiry_status"`
	Archived        bool    `json:"archived"`
	Favorite        bool    `json:"favorite"`
	Locked          bool    `json:"locked"`
	ExternalSharing bool    `json:"external_sharing"`
	ExternalURL     string  `json:"external_url"`
	ManagedBy       User    `json:"managed_by"`
	CreatedBy       User    `json:"created_by"`
	UpdatedBy       User    `json:"updated_by"`

	CustomField1 CustomField `json:"custom_field1"`
	CustomField2 CustomField `json:"custom_field2"`
	CustomField3 CustomField `json:"custom_field3"`
	CustomField4 CustomField `json:"custom_field4"`
	CustomField5 CustomField `json:"custom_field5"`
	CustomField6 CustomField `json:"custom_field6"`
	CustomField7 CustomField `json:"custom_field7"`
	CustomField8 CustomField `json:"custom_field8"`
	CustomField9 CustomField `json:"custom_field9"`

	// NumFiles        int     `json:"num_files"`
	// groups_permissions
	// parents
	// user_permission
	// users_permissions
}

type Passwords []Password

func (client *TpmClient) PasswordSearch(search string) (Passwords, error) {
	passwords := Passwords{}

	err := client.get("/api/v4/passwords/search/"+search+"/page/1.json", &passwords)
	if err != nil {
		return nil, err
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
