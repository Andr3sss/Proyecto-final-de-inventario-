package main

import (
	"html/template"
	"net/http"

	"Proyecto/INVENTARIO/configs"
	routes "Proyecto/INVENTARIO/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")

	// Configurar el HTMLRender
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Redirigir la ruta raíz a /login
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	// Rutas para las diferentes interfaces
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", nil)
	})

	r.GET("/productos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "productos.html", nil)
	})

	r.GET("/actualizacion", func(c *gin.Context) {
		c.HTML(http.StatusOK, "actualizacion.html", nil)
	})

	r.GET("/movimientos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "movimientos.html", nil)
	})

	r.GET("/inventario", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inventario.html", nil)
	})

	// Registrar las rutas
	routes.UsuarioRouter(r)
	routes.ProductoRouter(r)
	routes.MovimientoRouter(r)

	// Iniciar el servidor en el puerto 8080
	log.Fatal(r.Run(":8080"))
}
