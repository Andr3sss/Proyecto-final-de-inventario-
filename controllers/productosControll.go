package controllers

import (
	"net/http"
	"strconv"

	"Proyecto/INVENTARIO/configs"
	"Proyecto/INVENTARIO/models"

	"github.com/gin-gonic/gin"
)

type ProductoRequestBody struct {
	Nombre      string  `json:"nombre"`
	Descripción string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Color       string  `json:"color"`
	CategoriaID uint    `json:"categoria_id"`
}

func ProductoCreate(c *gin.Context) {
	var requestBody ProductoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	// Validar que la categoría exista y que el ID sea válido
	var categoria models.Categoria
	if requestBody.CategoriaID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}
	if err := configs.DB.First(&categoria, requestBody.CategoriaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Categoría no válida"})
		return
	}

	newProducto := models.Producto{
		Nombre:      requestBody.Nombre,
		Descripción: requestBody.Descripción,
		Precio:      requestBody.Precio,
		Color:       requestBody.Color,
		CategoriaID: requestBody.CategoriaID,
	}

	if err := configs.DB.Create(&newProducto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el producto: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProducto)
}

func ProductoGet(c *gin.Context) {
	var productos []models.Producto
	if err := configs.DB.Preload("Categoria").Find(&productos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los productos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, productos)
}

func ProductoGetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	var producto models.Producto
	if err := configs.DB.Preload("Categoria").First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, producto)
}

func ProductoUpdate(c *gin.Context) {
	// Obtener el ID del producto de la URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Obtener los datos del producto del cuerpo de la solicitud
	var requestBody ProductoRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que la categoría exista y que el ID sea válido
	var categoria models.Categoria
	if requestBody.CategoriaID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}
	if err := configs.DB.First(&categoria, requestBody.CategoriaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Categoría no válida"})
		return
	}

	// Actualizar el producto en la base de datos
	updatedProducto := models.Producto{
		Nombre:      requestBody.Nombre,
		Descripción: requestBody.Descripción,
		Precio:      requestBody.Precio,
		Color:       requestBody.Color,
		CategoriaID: requestBody.CategoriaID,
	}

	if err := configs.DB.Model(&models.Producto{}).Where("id = ?", id).Updates(updatedProducto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el producto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProducto)
}

func ProductoDelete(c *gin.Context) {
	id := c.Param("id")

	if err := configs.DB.Delete(&models.Producto{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el producto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
func CategoriaCreate(c *gin.Context) {
	// Definir una variable para el cuerpo de la solicitud
	var categoria models.Categoria

	// Intentar parsear el cuerpo de la solicitud en formato JSON
	if err := c.BindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	// Intentar guardar la nueva categoría en la base de datos
	if err := configs.DB.Create(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la categoría"})
		return
	}

	// Responder con la categoría creada y un código de estado 201 Created
	c.JSON(http.StatusCreated, categoria)
}

func CategoriaGet(c *gin.Context) {
	var categorias []models.Categoria
	if err := configs.DB.Find(&categorias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las categorías"})
		return
	}

	c.JSON(http.StatusOK, categorias)
}
