package db

import (
	"testing"
)

func Find_test(t *testing.T) {

	var dbClient DB
	dbClient = SqliteDB{}

	_, _, err := dbClient.Find("")
	if err != nil {
		t.Error(err)
	}
}
