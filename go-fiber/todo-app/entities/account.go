package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	ID       uint
	Username string
	Quota    uint
	Current  uint

	TodoList []Todo
}
