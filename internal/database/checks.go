package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Check is record of a check done by bot (rasp berry pi robot)
type Check struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	IsEmpty        bool      `json:"is_empty"`
	PlateNumber    string    `json:"plate_number" gorm:"uniqueIndex"`
	PackingSpaceID uint      `json:"packing_space_id"`
	CreatedAt      time.Time `json:"created_at"`
}

//AddUpdate adds a new check record to database
func (ck Check) AddUpdate(db *gorm.DB) error {

	num := db.Model(&Check{}).Where("plate_number=?", ck.PlateNumber).Update(&ck)
	if e := num.Error; e != nil {
		return e
	}

	if num.RowsAffected < int64(1) {
		num.Create(&ck)
		if e := num.Error; e != nil {
			return e
		}
	}

	return nil
}

//GetAll returns all the checks done
func (ck Check) GetAll(db *gorm.DB) ([]Check, error) {
	var cks []Check
	er := db.Model(&ck).Find(&cks).Error
	return cks, er
}
