package controllers

import (
	"go-audit-it/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSystemInput struct {
	Host   string `json:"host" binding:"required"`
	Status bool   `json:"status" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreateSystemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	system := models.System{Host: input.Host, Status: input.Status}
	models.DB.Create(&system)

	c.JSON(http.StatusOK, gin.H{"data": system})
}
