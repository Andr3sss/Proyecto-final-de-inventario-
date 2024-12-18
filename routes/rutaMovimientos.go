package routes

import (
	controladores "Proyecto/INVENTARIO/controllers"

	"github.com/gin-gonic/gin"
)

func MovimientoRouter(router *gin.Engine) {
	routes := router.Group("api/v1/movimientos")

	routes.GET("/", controladores.ObtenerMovimientos)

	// Ruta para registrar un movimiento de un producto (entrada o salida)
	routes.POST("/:id", controladores.RegistrarMovimiento)

	// Ruta para obtener todos los movimientos de un producto
	routes.GET("/:id", controladores.ObtenerMovimientos)

	// Ruta para eliminar un movimiento por su ID
	routes.DELETE("/:movimiento_id", controladores.EliminarMovimiento)
}
