package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
)

var err error
var db *gorm.DB

func InitDB() {
	db, err = gorm.Open(utils.Db, fmt.Sprint())
}
