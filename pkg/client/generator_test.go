package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	generateHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"password": "5&Dk4[e9.1#=0!2@_h3\"b8;t6$Bo7?"}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(generateHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	password, err := client.GeneratePassword()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if password.Value != "5&Dk4[e9.1#=0!2@_h3\"b8;t6$Bo7?" {
		t.Errorf("Unexpected result: %v", password.Value)
	}
}

func TestGeneratePasswordUnavailable(t *testing.T) {
	generateHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ts := httptest.NewServer(http.HandlerFunc(generateHandler))
	defer ts.Close()

	client := New(ts.URL, "fuu", "bar")

	_, err := client.GeneratePassword()
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}
