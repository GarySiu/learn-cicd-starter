package auth

import (
	"net/http"
	"testing"
)

func TestNoAuthHeader(t *testing.T) {
	var h http.Header
	_, err := GetAPIKey(h)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("incorrect error: %v", err)
	}
}

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey abc123")
	got, _ := GetAPIKey(h)
	if got != "abc123" {
		t.Errorf("unexpected api value: %v", got)
	}
}
