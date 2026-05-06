package main

import (
	"testing"
)

func TestGetenv(t *testing.T) {
	val := getenv("NON_EXISTENT_VAR_123", "fallback")
	if val != "fallback" {
		t.Errorf("expected fallback")
	}
}
