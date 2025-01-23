package services_test

import (
	"testing"
	"time"
)

func date(t *testing.T, value string) time.Time {
	t.Helper()

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		t.Fatalf("failed to parse date, error: %v", err)
	}

	return date
}
