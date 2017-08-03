package Helper

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
)

func Connection(param string) *gorm.DB {
	db, err := gorm.Open(param, MySQL())
		if(err != nil) {
			log.Println("Connection ORM Error")
		}
	return db
}