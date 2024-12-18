package routes

import (
	controladores "Proyecto/INVENTARIO/controllers"

	"github.com/gin-gonic/gin"
)

func ProductoRouter(router *gin.Engine) {
	routes := router.Group("api/v1/productos")

	// Ruta para crear un producto
	routes.POST("/", controladores.ProductoCreate)

	// Ruta para obtener todos los productos
	routes.GET("/", controladores.ProductoGet)

	// Ruta para obtener un producto por ID
	routes.GET("/:id", controladores.ProductoGetById)

	// Ruta para actualizar un producto por ID
	routes.PUT("/:id", controladores.ProductoUpdate)

	// Ruta para eliminar un producto por ID
	routes.DELETE("/:id", controladores.ProductoDelete)

	categorias := router.Group("api/v1/categorias")

	// Ruta para crear una categoria
	categorias.POST("/", controladores.CategoriaCreate)

	// Ruta para obtener todas las categorias
	categorias.GET("/", controladores.CategoriaGet)
}
