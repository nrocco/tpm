package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPasswordGet(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/passwords/1.json" {
			fmt.Fprint(w, `{
    "id": 1,
    "name": "CRM account",
    "project": {
        "id": 1,
        "name": "www.mynewsite.com"
    },
    "tags": "google",
    "access_info": "https:\/\/www.mynewsite.com\/crm",
    "username": "thisisme",
    "email": "thisisme@mynewsite.com",
    "password": "(Ip8=c1|9@{d5!2-0.u",
        "expiry_date": "2015-08-24",
        "expiry_status": 3,
        "notes": "Some notes for the password",
        "custom_field1": {
            "type": "Text",
            "label": "Custom 1",
            "data": "Sample value"
        },
        "custom_field2": {
            "type": "Password",
            "label": "Custom 2",
            "data": "6Zi3=!\/7.b9#1,2}|4;"
    },
    "custom_field3": null,
    "custom_field4": null,
    "custom_field5": null,
    "custom_field6": null,
    "custom_field7": null,
    "custom_field8": null,
    "custom_field9": null,
    "custom_field10": null,
    "users_permissions": [
        {
            "user": {
                "id": 1,
                "username": "clairewood",
                "email_address": "claire@teampasswordmanager.com",
                "name": "Claire Wood",
                "role": "Normal user"
            },
            "permission": {
                "id": 1,
                "label": "Edit data"
            }
        }
    ],
    "groups_permissions": [
        {
            "group": {
                "id": 1,
                "name": "SEO"
            }
        }
    ],
    "parents": [],
    "user_permission": {
        "id": 1,
        "label": "Manage"
    },
    "archived": false,
    "favorite": true,
    "num_files": 10,
    "locked": false,
    "external_sharing": true,
    "external_url": "http:\/\/localhost\/tpm2\/site_enc\/index.php\/pwde\/view\/15abe2439e994f91befe55075b9729e6df514e4669",
    "managed_by": {
        "id": 1,
        "username": "john",
        "email_address": "john@teampasswordmanager.com",
        "name": "John Boss",
        "role": "Admin"
    },
    "created_on": "2015-08-16 08:37:52",
    "created_by": {
        "id": 1,
        "username": "john",
        "email_address": "john@teampasswordmanager.com",
        "name": "John Boss",
        "role": "Admin"
    },
    "updated_on": "2015-08-16 10:57:06",
    "updated_by": {
        "id": 1,
        "username": "john",
        "email_address": "john@teampasswordmanager.com",
        "name": "John Boss",
        "role": "Admin"
    }
}`)
		} else {
			t.Errorf("Unexpected api call: %s", r.RequestURI)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	password, err := client.PasswordGet(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if password.ID != 1 {
		t.Errorf("Unexpected result: %v", password.ID)
	}
}

func TestPasswordGetUnavailable(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	_, err := client.PasswordGet(1)
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestPasswordList(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/passwords/count.json" {
			fmt.Fprint(w, `{"num_items": 4, "num_pages": 2, "num_items_per_page": 2}`)
		} else if r.RequestURI == "/api/v4/passwords/page/1.json" {
			fmt.Fprint(w, `[
    {
        "id": 1,
        "name": "Wordpress admin",
        "project": {
            "id": 1,
            "name": "www.fictionalgadgetsite.com"
        },
        "notes_snippet": "some notes\nother notes",
        "tags": "wordpress",
        "access_info": "http:\/\/www.fictionalgadgetsite.com\/wp-admin",
        "username": "admin_sg",
        "email": "",
        "expiry_date": "2013-10-20",
        "expiry_status": 2,
        "archived": false,
        "favorite": true,
        "num_files": 1,
        "locked": false,
        "external_sharing": true,
        "updated_on": "2014-07-23 18:18:28"
    },
    {
        "id": 2,
        "name": "Sample password",
        "project": {
            "id": 1,
            "name": "www.fictionalgadgetsite.com"
        },
        "archived": false,
        "favorite": true,
        "locked": true,
        "external_sharing": false,
        "updated_on": "2014-07-23 18:18:06"
    }
]`)
		} else if r.RequestURI == "/api/v4/passwords/page/2.json" {
			fmt.Fprint(w, `[
    {
        "id": 3,
        "name": "Nagios online",
        "project": {
            "id": 1,
            "name": "www.nagios.com"
        },
        "notes_snippet": "some notes\nother notes",
        "tags": "nagios",
        "access_info": "http:\/\/www.fooobar.com",
        "username": "yala yala",
        "email": "",
        "expiry_date": "2013-10-20",
        "expiry_status": 2,
        "archived": false,
        "favorite": true,
        "num_files": 1,
        "locked": false,
        "external_sharing": true,
        "updated_on": "2014-07-23 18:18:28"
    },
    {
        "id": 4,
        "name": "My KVM password",
        "project": {
            "id": 1,
            "name": "www.kvmonline.com"
        },
        "archived": false,
        "favorite": true,
        "locked": true,
        "external_sharing": false,
        "updated_on": "2014-07-23 18:18:06"
    }
]`)
		} else {
			t.Errorf("Unexpected api call: %s", r.RequestURI)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	passwords, err := client.PasswordList("")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(passwords) != 4 {
		t.Errorf("Unexpected result: %v", len(passwords))
	}
}

func TestPasswordSearch(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/passwords/search/fuu/count.json" {
			fmt.Fprint(w, `{"num_items": 1, "num_pages": 1, "num_items_per_page": 1}`)
		} else if r.RequestURI == "/api/v4/passwords/search/fuu/page/1.json" {
			fmt.Fprint(w, `[
    {
        "id": 2,
        "name": "Sample password Fuu",
        "project": {
            "id": 1,
            "name": "www.fictionalgadgetsite.com"
        },
        "archived": false,
        "favorite": true,
        "locked": true,
        "external_sharing": false,
        "updated_on": "2014-07-23 18:18:06"
    }
]`)
		} else {
			t.Errorf("Unexpected api call: %s", r.RequestURI)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	passwords, err := client.PasswordList("fuu")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(passwords) != 1 {
		t.Errorf("Unexpected result: %v", len(passwords))
	}
}
