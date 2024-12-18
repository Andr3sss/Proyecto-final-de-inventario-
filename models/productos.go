package models

import (
	"gorm.io/gorm"
)

type Producto struct {
	gorm.Model
	Nombre      string `gorm:"not null"`
	Descripción string
	Precio      float64 `gorm:"not null"`
	CategoriaID uint
	Categoria   Categoria
	Color       string       `gorm:"not null"` // Código de color del producto
	Movimientos []Movimiento `gorm:"foreignKey:ProductoID"`
	Stock       int
}

type Categoria struct {
	gorm.Model
	Nombre    string `gorm:"not null"`
	Productos []Producto
}
