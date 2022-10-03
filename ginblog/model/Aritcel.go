package model

import "github.com/jinzhu/gorm"

// 文章作者
type Aritcel struct {
	gorm.Model
	Category Category
	Title    string `gorm:"type:varchar(40);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`
	Desc     string `gorm:"type:varchar(200)" json:"desc"`
	Content  string `gorm:"type:longtext;not null" json:"content"`
	Image    string `gorm:"type:varchar(100)" json:"image"`
}
