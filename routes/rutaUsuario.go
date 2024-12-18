package routes

import (
	controladores "Proyecto/INVENTARIO/controllers" // Asegúrate de que la ruta del controlador sea correcta

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.Engine) {
	routes := router.Group("api/v1/Usuarios")

	// Ruta para crear un usuario
	routes.POST("/usuarios", controladores.UsuarioCreate)

	// Ruta para obtener un usuario por ID
	routes.GET("/usuarios/:id", controladores.UsuarioGet)

	// Ruta para actualizar un usuario por ID
	routes.PUT("/usuarios/:id", controladores.UsuarioUpdate)

	// Ruta para eliminar un usuario por ID
	routes.DELETE("/usuarios/:id", controladores.UsuarioDelete)

	// Nueva ruta para validar el código de un usuario
	routes.POST("/validate", controladores.UsuarioValidate) // Validar código del usuario
}
