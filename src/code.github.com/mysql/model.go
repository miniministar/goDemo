package mysql

import "github.com/jinzhu/gorm"

type user struct {
	Id   int
	Age  int
	Name string
}

type AccountInfo struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not null;"`
	Status   uint   `gorm:"default:0"`
}
