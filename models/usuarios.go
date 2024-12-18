package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuarios struct {
	gorm.Model
	ID        uint
	Codigo    int64 `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Usuarios) TableName() string {
	return "Usuarios"
}
