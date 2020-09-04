package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syahjamal/gin-full-api/models"
)

//Variable global buat baca db
var DB *gorm.DB

func InitDB() {
	var err error

	//koneksi db
	DB, err = gorm.Open("mysql", "root:@/go_gin_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrasi Model Article dari gorm
	DB.AutoMigrate(&models.Article{})
}
