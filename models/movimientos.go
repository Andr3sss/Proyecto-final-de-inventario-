package models

import (
	"time"

	"gorm.io/gorm"
)

type Movimiento struct {
	gorm.Model
	ProductoID     uint      `gorm:"not null"`
	TipoMovimiento string    `gorm:"not null"`
	Fecha          time.Time `gorm:"not null"`
	Cantidad       int       `gorm:"not null"`
	Motivo         string
}
