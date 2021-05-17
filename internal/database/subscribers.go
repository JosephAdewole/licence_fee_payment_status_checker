package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Subscriber represents a packing space user, in this case a customer
type Subscriber struct {
	ID             int       `json:"id" gorm:"primary_key"`
	PlateNumber    string    `json:"plate_number"`
	PackingSpaceID int       `json:"packing_space_id"`
	StartTime      time.Time `json:"start_time"`
	Status         bool      `json:"status"`
}

//Add adds a new subcriber to database
func (sub Subscriber) Add(db *gorm.DB) error {
	return db.Model(&sub).Create(sub).Error
}

//GetAll returns a list of all subcribers from the database
func (sub Subscriber) GetAll(db *gorm.DB) ([]Subscriber, error) {
	var subs []Subscriber
	er := db.Model(&sub).Find(&subs).Error

	return subs, er
}
