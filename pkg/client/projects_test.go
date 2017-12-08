package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectGet(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/projects/1.json" {
			fmt.Fprint(w, `{
    "id": 20,
    "name": "www.mynewsite.com",
    "parent_id": 19,
    "tags": "dit,is,een,tags",
    "notes": "SEO for www.mynewsite.com.",
    "managed_by": {
        "id": 2,
        "username": "alan",
        "email_address": "alan@teampasswordmanager.com",
        "name": "Alan",
        "role": "Project manager"
    },
    "grant_all_permission": {
        "id": -1,
        "label": "(Do not set)"
    },
    "users_permissions": [
        {
            "user": {
                "id": 3,
                "username": "ann",
                "email_address": "ann@teampasswordmanager.com",
                "name": "Ann",
                "role": "Normal user"
            },
            "permission": {
                "id": 50,
                "label": "Read \/ Manage passwords"
            }
        }
    ],
    "groups_permissions": [
        {
            "group": {
                "id": 1,
                "name": "SEO"
            },
            "permission": {
                "id": 20,
                "label": "Read"
            }
        }
    ],
    "num_passwords": 1,
    "num_files": 1,
    "user_permission": {
        "id": 60,
        "label": "Manage"
    },
    "user_can_create_passwords": true,
    "is_leaf": true,
    "parents": [
        13,
        19
    ],
    "archived": false,
    "favorite": true,
    "created_on": "2015-08-16 02:16:49",
    "created_by": {
        "id": 1,
        "username": "john",
        "email_address": "john@teampasswordmanager.com",
        "name": "John Boss",
        "role": "Admin"
    },
    "updated_on": "2015-08-16 08:37:14",
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

	project, err := client.ProjectGet("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if project.ID != "20" {
		t.Errorf("Unexpected result: %v", project.ID)
	}
}

func TestProjectGetUnavailable(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	_, err := client.ProjectGet("1")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestProjectList(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/projects/count.json" {
			fmt.Fprint(w, `{"num_items": 4, "num_pages": 2, "num_items_per_page": 2}`)
		} else if r.RequestURI == "/api/v4/projects/page/1.json" {
			fmt.Fprint(w, `[
    {
        "id": 1,
        "name": "Project created with the API",
        "tags": "client,seo,website",
        "managed_by": {
            "id": 1,
            "name": "John Boss"
        },
        "archived": false,
        "favorite": true,
        "num_files": 0,
        "updated_on": "2014-07-04 14:02:14"
    },
    {
        "id": 2,
        "name": "www.fictionalgadgetsite.com",
        "tags": "client",
        "managed_by": {
            "id": 3,
            "username": "claire",
            "email_address": "claire@teampasswordmanager.com",
            "name": "Claire Wood",
            "role": "Project manager"
        },
        "archived": false,
        "favorite": false,
        "num_files": 0,
        "updated_on": "2014-07-03 19:13:37"
    }
]`)
		} else if r.RequestURI == "/api/v4/projects/page/2.json" {
			fmt.Fprint(w, `[
    {
        "id": 3,
        "name": "Some fancy stuff",
        "tags": "",
        "managed_by": {
            "id": 1,
            "name": "John Boss"
        },
        "archived": false,
        "favorite": true,
        "num_files": 0,
        "updated_on": "2014-07-04 14:02:14"
    },
    {
        "id": 4,
        "name": "www.foobarbla.com",
        "tags": "foo,bar,stakker,akker",
        "managed_by": {
            "id": 3,
            "username": "claire",
            "email_address": "claire@teampasswordmanager.com",
            "name": "Claire Wood",
            "role": "Project manager"
        },
        "archived": false,
        "favorite": false,
        "num_files": 0,
        "updated_on": "2014-07-03 19:13:37"
    }
]`)
		} else {
			t.Errorf("Unexpected api call: %s", r.RequestURI)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	projects, err := client.ProjectList("")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(projects) != 4 {
		t.Errorf("Unexpected result: %v", len(projects))
	}
}

func TestProjectSearch(t *testing.T) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/v4/projects/search/foo/count.json" {
			fmt.Fprint(w, `{"num_items": 1, "num_pages": 1, "num_items_per_page": 2}`)
		} else if r.RequestURI == "/api/v4/projects/search/foo/page/1.json" {
			fmt.Fprint(w, `[
    {
        "id": 4,
        "name": "www.foobarbla.com",
        "tags": "foo,bar,stakker,akker",
        "managed_by": {
            "id": 3,
            "username": "claire",
            "email_address": "claire@teampasswordmanager.com",
            "name": "Claire Wood",
            "role": "Project manager"
        },
        "archived": false,
        "favorite": false,
        "num_files": 0,
        "updated_on": "2014-07-03 19:13:37"
    }
]`)
		} else {
			t.Errorf("Unexpected api call: %s", r.RequestURI)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	projects, err := client.ProjectList("foo")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(projects) != 1 {
		t.Errorf("Unexpected result: %v", len(projects))
	}
}
