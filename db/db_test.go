package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	clearTables()
	db := GetDB()
	if db == nil {
		t.Errorf("get a nil db")
	}
	assert.NotNil(t, db, "db is not be nil")
}
