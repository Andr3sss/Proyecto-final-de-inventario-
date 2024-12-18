package controllers

import (
	"Proyecto/INVENTARIO/configs"
	"Proyecto/INVENTARIO/models"

	"github.com/gin-gonic/gin"
)

type UsuarioRequestBody struct {
	Codigo int64 `json:"codigo"`
}

func UsuarioCreate(c *gin.Context) {
	body := UsuarioRequestBody{}

	// Parseamos el cuerpo de la solicitud
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid request body"})
		return
	}

	// Verificamos si el código ya existe en la base de datos
	var existingUsuario models.Usuarios
	result := configs.DB.Where("codigo = ?", body.Codigo).First(&existingUsuario)

	if result.RowsAffected > 0 { // Si existe un registro con el mismo código
		c.JSON(409, gin.H{"Error": "El código ya está registrado."})
		return
	}

	// Si no existe, creamos un nuevo usuario
	newUsuario := models.Usuarios{Codigo: body.Codigo}
	createResult := configs.DB.Create(&newUsuario)

	// Si ocurrió un error al insertar, respondemos con un error
	if createResult.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	// Si la creación fue exitosa, respondemos con el usuario creado
	c.JSON(200, &newUsuario)
}

func UsuarioValidate(c *gin.Context) {
	// Obtener el código enviado en el cuerpo de la solicitud
	var body struct {
		Codigo int64 `json:"codigo"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid request body"})
		return
	}

	// Verificar si el código existe en la base de datos
	var usuario models.Usuarios
	result := configs.DB.Where("codigo = ?", body.Codigo).First(&usuario)

	if result.Error != nil {
		// Si no se encuentra el código en la base de datos
		c.JSON(404, gin.H{"Error": "Código no encontrado"})
		return
	}

	// Si se encuentra el código
	c.JSON(200, gin.H{"Success": "Código válido"})
}
func UsuarioGet(c *gin.Context) {
	var usuarios []models.Usuarios

	// Obtenemos todos los usuarios desde la base de datos
	result := configs.DB.Find(&usuarios)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to retrieve records"})
		return
	}

	// Respondemos con la lista de usuarios
	c.JSON(200, &usuarios)
}

func UsuarioGetById(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuarios

	// Buscamos el usuario por su ID
	result := configs.DB.First(&usuario, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"Error": "Record not found"})
		return
	}

	// Respondemos con el usuario encontrado
	c.JSON(200, &usuario)
}

func UsuarioUpdate(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuarios

	// Buscamos el usuario por su ID
	result := configs.DB.First(&usuario, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"Error": "Record not found"})
		return
	}

	body := UsuarioRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid request body"})
		return
	}

	// Actualizamos el código del usuario
	updatedData := &models.Usuarios{Codigo: body.Codigo}
	updateResult := configs.DB.Model(&usuario).Updates(updatedData)
	if updateResult.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to update"})
		return
	}

	// Respondemos con el usuario actualizado
	c.JSON(200, &usuario)
}

func UsuarioDelete(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuarios

	// Eliminamos el usuario por su ID
	result := configs.DB.Delete(&usuario, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to delete"})
		return
	}

	// Respondemos indicando que el registro fue eliminado
	c.JSON(200, gin.H{"deleted": true})
}
