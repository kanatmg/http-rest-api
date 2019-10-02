package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TextMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "furdiegang:furdiegang@/godb"
	}

	os.Exit(m.Run())
}
