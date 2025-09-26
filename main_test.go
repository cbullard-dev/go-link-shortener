package main

import (
	"regexp"
	"testing"
)

func TestUrlGenerationLength(t *testing.T) {
	want := UrlCodeLength
	code := generateUrlCode()
	if len(code) != want {
		t.Errorf(`Generated code expected length: %q, actual length %q`, want, len(code))
	}
}

func TestUrlGenerationFormat(t *testing.T) {
	code := generateUrlCode()
	want, _ := regexp.Match("[a-zA-Z0-9]{8}", []byte(code))
	if !want {
		t.Errorf(`Generated code expected to match regex "[a-zA-Z0-9]{8}", actual length result: %q`, code)
	}
}
