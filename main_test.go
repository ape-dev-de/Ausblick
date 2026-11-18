package main

import (
	"testing"
)

func TestGetGreeting(t *testing.T) {
	expected := "Ausblick v1.0.0"
	actual := getGreeting()

	if actual != expected {
		t.Errorf("getGreeting() = %q; erwartet %q", actual, expected)
	}
}
