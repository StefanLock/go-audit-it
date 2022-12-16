package main

import (
	"go-audit-it/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase() // new!

	// ...

	router.Run("localhost:5000")
}
