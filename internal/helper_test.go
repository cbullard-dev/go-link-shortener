package helper

import (
	"regexp"
	"testing"
)

func TestUrlGenerationLength(t *testing.T) {
	want := 8
	code := GenerateUrlCode(want)
	if len(code) != want {
		t.Errorf(`Generated code expected length: %q, actual length %q`, want, len(code))
	}
}

func TestUrlGenerationFormat(t *testing.T) {
	code := GenerateUrlCode(8)
	want, _ := regexp.Match("[a-zA-Z0-9]{8}", []byte(code))
	if !want {
		t.Errorf(`Generated code expected to match regex "[a-zA-Z0-9]{8}", actual length result: %q`, code)
	}
}
