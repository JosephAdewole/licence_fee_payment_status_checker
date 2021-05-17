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

//Add adds a new subcriber to database
func (sub Subscriber) Add(db *gorm.DB) error {
	return db.Model(&sub).Where("plate_number=?", sub.PlateNumber).Create(sub).Error
}

//GetAll returns a list of all subcribers from the database
func (sub Subscriber) GetAll(db *gorm.DB) ([]Subscriber, error) {
	var subs []Subscriber
	er := db.Model(&sub).Find(&subs).Error

	return subs, er
}
