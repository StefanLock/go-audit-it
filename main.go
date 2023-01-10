package main

import (
	"go-audit-it/controllers"
	"go-audit-it/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	models.ConnectDatabase()

	router.POST("/systems", controllers.CreateSystem)
	router.GET("/systems", controllers.FindSystems)
	router.GET("/systems/:id", controllers.FindSystem)
	router.PATCH("/systems/:id", controllers.UpdateSystem)
	router.DELETE("/systems/:id", controllers.DeleteSystem)

	router.Run("localhost:5000")
}
