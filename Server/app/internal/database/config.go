package database

import "database/sql"

//Config is a record of a configuration
type Config struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Add adds a new config record to database
func (c Config) Add(db *sql.DB) error {
	return nil
}

//Update updates the record of a configuration in a database
func (c Config) Update(db *sql.DB) error {
	return nil
}
