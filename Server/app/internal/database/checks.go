package database

import (
	"database/sql"
	"time"
)

//Check is record of a check done by bot (rasp berry pi robot)
type Check struct {
	ID             int       `json:"id"`
	IsEmpty        bool      `json:"is_empty"`
	PlateNumber    string    `json:"plate_number"`
	PackingSpaceID int       `json:"packing_space_id"`
	CreatedAt      time.Time `json:"created_at"`
}

//Add adds a new check record to database
func (ck Check) Add(db *sql.DB) error {
	return nil
}

//GetAll returns all the checks done
func (ck Check) GetAll(db *sql.DB) ([]Check, error) {

	return nil, nil
}
