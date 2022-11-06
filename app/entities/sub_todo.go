package entities

import (
	"time"

	"gorm.io/gorm"
)

type SubTodo struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"tittle" gorm:"size:100;not null"`
	Description string `json:"description" gorm:"size:1000;not null"`
	File        string `json:"file"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	TodoID      int
	Todo        Todo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
