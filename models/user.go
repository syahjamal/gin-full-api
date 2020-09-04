package models

import "github.com/jinzhu/gorm"

//Struktur kolom
type User struct {
	gorm.Model
	Articles []Article //relasi
	Username string
	FullName string
	Email    string
	SocialId string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
