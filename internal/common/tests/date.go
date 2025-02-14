package tests

import (
	"testing"
	"time"
)

func Date(t *testing.T, input string) time.Time {
	t.Helper()

	date, err := time.Parse(time.DateOnly, input)
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	return date
}
