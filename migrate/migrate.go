package main

import (
	"Proyecto/INVENTARIO/configs"
	"Proyecto/INVENTARIO/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Usuarios{})
	configs.DB.AutoMigrate(&models.Producto{})
	configs.DB.AutoMigrate(&models.Movimiento{})

}
