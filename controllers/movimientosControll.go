package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Proyecto/INVENTARIO/configs"
	"Proyecto/INVENTARIO/models"

	"github.com/gin-gonic/gin"
)

type MovimientoRequestBody struct {
	TipoMovimiento string `json:"tipo_movimiento"`
	Cantidad       int    `json:"cantidad"`
	Motivo         string `json:"motivo"`
}

func RegistrarMovimiento(c *gin.Context) {
	id := c.Param("id")

	// Validar ID del producto
	productoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Obtener el producto con sus movimientos asociados
	var producto models.Producto
	if err := configs.DB.Preload("Movimientos").First(&producto, productoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	// Leer cuerpo de la solicitud
	var requestBody MovimientoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	// Validar datos del movimiento
	if requestBody.TipoMovimiento != "entrada" && requestBody.TipoMovimiento != "salida" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de movimiento inválido"})
		return
	}
	if requestBody.Cantidad <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad debe ser mayor a 0"})
		return
	}

	// Crear y registrar el movimiento
	nuevoMovimiento := models.Movimiento{
		ProductoID:     uint(productoID),
		TipoMovimiento: requestBody.TipoMovimiento,
		Fecha:          time.Now(),
		Cantidad:       requestBody.Cantidad,
		Motivo:         requestBody.Motivo,
	}

	if err := configs.DB.Create(&nuevoMovimiento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el movimiento"})
		return
	}

	// Habilita la depuración de GORM para ver las consultas SQL en la consola
	configs.DB.Debug().Preload("Movimientos").First(&producto, productoID)

	fmt.Println("Ejecutando consulta para obtener el stock para el producto ID:", producto.ID)

	// Obtener el stock actual del producto (incluyendo el nuevo movimiento)
	var stockActual int
	// Ejecutamos la consulta
	if err := configs.DB.Model(&producto).Update("Stock", stockActual).Error; err != nil {
		fmt.Println("Error al actualizar el stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el stock del producto"})
		return
	}
	fmt.Println("Stock actualizado en la base de datos:", stockActual)

	// Responder con el movimiento y el stock actual
	c.JSON(http.StatusCreated, gin.H{
		"movimiento":   nuevoMovimiento,
		"stock_actual": stockActual,
		"producto":     producto,
	})

}

func ObtenerMovimientos(c *gin.Context) {
	id := c.Param("id")

	var movimientos []models.Movimiento
	if err := configs.DB.Where("producto_id = ?", id).Find(&movimientos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los movimientos"})
		return
	}

	c.JSON(http.StatusOK, movimientos)
}

func EliminarMovimiento(c *gin.Context) {
	movimientoID, err := strconv.Atoi(c.Param("movimiento_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de movimiento inválido"})
		return
	}

	if err := configs.DB.Delete(&models.Movimiento{}, movimientoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el movimiento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movimiento eliminado"})
}
