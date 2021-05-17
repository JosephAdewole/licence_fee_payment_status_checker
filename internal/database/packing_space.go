package database

import "github.com/jinzhu/gorm"

//PackingSpace represents a single packing spot
type PackingSpace struct {
	ID          uint   `json:"id" gorm:"uniqueIndex"`
	Designation string `json:"designation"`
}

//AddUpdate adds or updates a packing spot/slot in database tables
func (ps *PackingSpace) AddUpdate(db *gorm.DB) error {

	num := db.Model(&PackingSpace{}).Create(ps)

	if num.RowsAffected < int64(1) {
		e := db.Model(&Subscriber{}).Where("id = ?", ps.ID).UpdateColumn(ps)
		db.Model(&Subscriber{}).Where("id = ?", ps.ID).First(ps)
		if e.Error != nil {
			return e.Error
		}
	}

	db.Model(&Subscriber{}).Where("id = ?", ps.ID).First(ps)

	return nil
}

//GetAll returns a list of all the packing space/slots stored in database
func (ps PackingSpace) GetAll(db *gorm.DB) ([]PackingSpace, error) {
	var pss []PackingSpace
	er := db.Model(&ps).Find(&pss).Error
	return pss, er
}
