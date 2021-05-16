package handlers

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

//Connect returns a connection pointer to database
func Connect(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open("mysql", db)
}
