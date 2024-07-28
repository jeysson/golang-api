package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string   `json:"tilte"`
	Edition uint     `json:"edition"`
	Authors []Author `gorm:"many2many:author_books"`
}
