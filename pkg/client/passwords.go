package client

type CustomField struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Data  string `json:"data"`
}

type Password struct {
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

type Passwords []Password
