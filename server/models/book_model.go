package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	//ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}
