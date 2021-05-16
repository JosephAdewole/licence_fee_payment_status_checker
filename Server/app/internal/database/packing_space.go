package database

import "database/sql"

//PackingSpace represents a single packing spot
type PackingSpace struct {
	ID          int    `json:"id"`
	Designation string `json:"designation"`
}

//Add adds a packing spot/slot to database tables
func (ps PackingSpace) Add(db *sql.DB) error {

	return nil
}

//GetAll returns a list of all the packing space/slots stored in database
func (ps PackingSpace) GetAll(db *sql.DB) ([]PackingSpace, error) {
	return nil, nil
}
