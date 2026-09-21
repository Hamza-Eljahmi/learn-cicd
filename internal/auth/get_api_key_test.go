package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey example-key")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "example-key" {
		t.Errorf("expected key %q, got %q", "example-key", key)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected missing-header error, got %v", err)
	}
}

func TestGetAPIKeyWrongPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer example-key")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error for the wrong authorization prefix")
	}
}
