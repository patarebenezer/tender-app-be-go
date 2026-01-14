package db

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(dbPath string) (*gorm.DB, error) {
	_ = os.MkdirAll(filepath.Dir(dbPath), 0755)
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}
