package database

import "github.com/jinzhu/gorm"

//PackingSpace represents a single packing spot
type PackingSpace struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Designation string `json:"designation"`
}

//AddUpdate adds or updates a packing spot/slot in database tables
func (ps PackingSpace) AddUpdate(db *gorm.DB) error {

	num := db.Model(&PackingSpace{}).Where("id=?", ps.ID).Update(&ps)
	if e := num.Error; e != nil {
		return e
	}

	if num.RowsAffected < int64(1) {
		num.Create(&ps)
		if e := num.Error; e != nil {
			return e
		}
	}

	return nil
}

//GetAll returns a list of all the packing space/slots stored in database
func (ps PackingSpace) GetAll(db *gorm.DB) ([]PackingSpace, error) {
	var pss []PackingSpace
	er := db.Model(&ps).Find(&pss).Error
	return pss, er
}
