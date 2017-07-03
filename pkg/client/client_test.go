package client

import (
	"testing"
)

func TestNewTpmClient(t *testing.T) {
	Version = "test"

	client := New("http://fafa", "fuu", "bar")

	if client.ContentType != "application/json; charset=utf-8" {
		t.Error("Incorrect content type")
	}

	if client.UserAgent != "tpm/test (darwin/amd64)" {
		t.Errorf("Incorrect user agent: %s", client.UserAgent)
	}

	if client.Server != "http://fafa" {
		t.Errorf("Incorrect user agent: %s", client.Server)
	}

	if client.Username != "fuu" {
		t.Errorf("Incorrect user agent: %s", client.Username)
	}

	if client.Password != "bar" {
		t.Errorf("Incorrect user agent: %s", client.Password)
	}
}
