package database

import "github.com/jinzhu/gorm"

//Config is a record of a configuration
type Config struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"  gorm:"uniqueIndex"`
	Value string `json:"value"`
}

//AddUpdate adds or updates a config record in database
func (c Config) AddUpdate(db *gorm.DB) error {
	num := db.Model(&Config{}).Where("name=?", c.Name).Create(&c)
	if e := num.Error; e != nil {
		return e
	}

	if num.RowsAffected < int64(1) {
		num.Update(&c)
		if e := num.Error; e != nil {
			return e
		}
	}

	return nil
}
