package entities

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int    `gorm:"primary_key"`
	Title       string `gorm:"size:100;not null"`
	Description string `gorm:"size:1000";not null`
	File        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	SubTodo     []SubTodo `gorm:"ForeignKey:TodoID"`
}
