package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Subscriber represents a packing space user, in this case a customer
type Subscriber struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	PlateNumber string    `json:"plate_number"  gorm:"uniqueIndex"`
	StartTime   time.Time `json:"start_time"`
	Status      bool      `json:"status"`
}

//AddUpdate adds or update subcriber record to database
func (sub Subscriber) AddUpdate(db *gorm.DB) error {

	num := db.Model(&Subscriber{}).Create(&sub)

	if num.RowsAffected < int64(1) {
		num.UpdateColumn(&sub)
		num.Where("plate_number = ?", sub.PlateNumber).First(&sub)
		if e := num.Error; e != nil {
			return e
		}
	}

	return num.Error
}

//GetAll returns a list of all subcribers from the database
func (sub Subscriber) GetAll(db *gorm.DB) ([]Subscriber, error) {
	var subs []Subscriber
	er := db.Model(&sub).Find(&subs).Error

	return subs, er
}
