package entities

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uint
	Title     string `json:"title"`
	Author    string `json:"author"`
	AccountID uint
	Status    string `json:"status"`
}
