package db

import (
	"testing"
)

func clearTables() {
	db.Exec("truncate DemoOrder")
}

func TestMain(m *testing.M) {
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("get", testGetDB)
}

func testGetDB(t *testing.T) {
	db := GetDB()
	if db == nil {
		t.Errorf("get a nil db")
	}
}
