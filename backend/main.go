package main

import (
	"PROYECTO/controller"
	"PROYECTO/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	SearchController := controller.NewSearchController()
	AuthController := controller.NewJwtController()
	DBService := service.NewSaveSongsService()

	fmt.Println("Servidor en linea ....")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//Crear la tabla si no existe
	errDB := DBService.InitiateDB()
	if errDB != nil {
		fmt.Println("Error1 al iniciar la base de datos: ", errDB.Error())
	}

	//Endpoint de prueba
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Endpoint de busqueda
	r.GET("/search", SearchController.FindSongs)

	//Endpoint de autenticacion
	r.GET("/get-token", AuthController.GetToken)

	r.Run(":8000")
}
