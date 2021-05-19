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
	PlateNumber    string    `json:"plate_number"`
	PackingSpaceID uint      `json:"packing_space_id" gorm:"uniqueIndex"`
	CreatedAt      time.Time `json:"created_at"`
}

//AddUpdate adds a new check record to database
func (ck *Check) AddUpdate(db *gorm.DB) error {

	num := db.Model(&Check{}).Create(ck)

	if num.RowsAffected < int64(1) {
		e := db.Model(&Check{}).Where("packing_space_id = ?", ck.PackingSpaceID).UpdateColumn(ck)
		db.Model(&Check{}).Where("packing_space_id = ?", ck.PackingSpaceID).First(ck)
		if e.Error != nil {
			return e.Error
		}
	}
	db.Model(&Check{}).Where("packing_space_id = ?", ck.PackingSpaceID).First(ck)

	return nil
}

//GetAll returns all the checks done
func (ck Check) GetAll(db *gorm.DB) ([]Check, error) {
	var cks []Check
	er := db.Model(&ck).Find(&cks).Error
	return cks, er
}
