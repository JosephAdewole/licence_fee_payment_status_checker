package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
func (ck Check) Add(db *gorm.DB) error {
	return db.Create(&ck).Error
}

//GetAll returns all the checks done
func (ck Check) GetAll(db *gorm.DB) ([]Check, error) {
	var cks []Check
	er := db.Model(&ck).Find(&cks).Error
	return cks, er
}
