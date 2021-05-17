package database

import "github.com/jinzhu/gorm"

//PackingSpace represents a single packing spot
type PackingSpace struct {
	ID          int    `json:"id"`
	Designation string `json:"designation"`
}

//Add adds a packing spot/slot to database tables
func (ps PackingSpace) Add(db *gorm.DB) error {
	return db.Create(&ps).Error
}

//GetAll returns a list of all the packing space/slots stored in database
func (ps PackingSpace) GetAll(db *gorm.DB) ([]PackingSpace, error) {
	var pss []PackingSpace
	er := db.Model(&ps).Find(&pss).Error
	return pss, er
}
