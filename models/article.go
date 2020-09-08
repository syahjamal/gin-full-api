package models

import "github.com/jinzhu/gorm"

//Models
type Article struct {
	gorm.Model //Wajib untuk db apapun
	Title      string
	Slug       string `gorm:"unique_index"` //Kolom data ramah url dan urlnya cth: "judul-pertama"
	Desc       string `sql:"type:text;"`
	Tag        string
	UserID     uint //relasi ke user.go
}
