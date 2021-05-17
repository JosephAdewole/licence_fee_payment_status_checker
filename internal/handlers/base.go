package handlers

import (
	"fmt"
	"mawakif/config"

	"github.com/jinzhu/gorm"
)

//ConnectionString returns a string for database connection
func ConnectionString(cfg config.CONFIG) string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassWord, cfg.DBName)
}

//Connect returns a connection pointer to database
//returns *gorm.DB , a close function and error if any
func Connect(connectionString string) (*gorm.DB, func(), error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, nil, err
	}

	cls := func() { db.Close() }

	return db, cls, nil
}
