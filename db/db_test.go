package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDB(t *testing.T) {
	db := GetDB()
	if db == nil {
		t.Errorf("get a nil db")
	}
	assert.NotNil(t, db, "db is not be nil")
	clearTables()
}
