package database

import "github.com/jinzhu/gorm"

//Config is a record of a configuration
type Config struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"  gorm:"uniqueIndex"`
	Value string `json:"value"`
}

//AddUpdate adds or updates a config record in database
func (c *Config) AddUpdate(db *gorm.DB) error {

	num := db.Model(&Config{}).Create(c)

	if num.RowsAffected < int64(1) {
		e := db.Model(&Subscriber{}).Where("name = ?", c.Name).UpdateColumn(c)
		db.Model(&Subscriber{}).Where("name = ?", c.Name).First(c)
		if e.Error != nil {
			return e.Error
		}
	}

	db.Model(&Subscriber{}).Where("name = ?", c.Name).First(c)

	return nil
}

//Get gets the value from database
func (c *Config) Get(db *gorm.DB) error {
	return db.Model(&Config{}).Where("name=?", c.Name).First(c).Error
}
