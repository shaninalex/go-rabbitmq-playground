package db_test

import (
	"orders/internal/db"
	"testing"
)

func TestInitializeDatabase(t *testing.T) {
	_, err := db.InitializeDatabase(":memory:")
	if err != nil {
		t.Errorf("Cant create db instance: %v", err)
	}

}
