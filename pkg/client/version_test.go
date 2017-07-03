package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVersion(t *testing.T) {
	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"version": "7.50.100", "version_date": "2017-06-21", "api_version": "4"}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(versionHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	version, err := client.Version()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if version.Version != "7.50.100" {
		t.Errorf("Unexpected result: %v", version.Version)
	}

	if version.Date != "2017-06-21" {
		t.Errorf("Unexpected result: %v", version.Date)
	}

	if version.Number != "4" {
		t.Errorf("Unexpected result: %v", version.Number)
	}
}

func TestGetVersionUnavailable(t *testing.T) {
	generateHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ts := httptest.NewServer(http.HandlerFunc(generateHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	_, err := client.Version()
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}
