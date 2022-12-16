package controllers

import (
	"go-audit-it/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSystemInput struct {
	Host   string `json:"host"`   // binding:"required"`
	Status string `json:"status"` // binding:"required"`
}

func CreateSystem(c *gin.Context) {
	var input CreateSystemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	system := models.System{Host: input.Host, Status: input.Status}
	models.DB.Create(&system)

	c.JSON(http.StatusOK, gin.H{"data": system})
}

func FindSystems(c *gin.Context) {
	var systems []models.System
	models.DB.Find(&systems)

	c.JSON(http.StatusOK, gin.H{"data": systems})
}

func FindSystem(c *gin.Context) {
	var system models.System

	if err := models.DB.Where("id = ?", c.Param("id")).First(&system).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": system})
}

type UpdateSystemInput struct {
	Host   string `json:"host"`
	Status string `json:"status"`
}

func UpdateSystem(c *gin.Context) {
	var system models.System
	if err := models.DB.Where("id = ?", c.Param("id")).First(&system).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdateSystemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSystem := models.System{Host: input.Host, Status: input.Status}

	models.DB.Model(&system).Updates(&updatedSystem)
	c.JSON(http.StatusOK, gin.H{"data": system})
}

func DeleteSystem(c *gin.Context) {
	var system models.System
	if err := models.DB.Where("id = ?", c.Param("id")).First(&system).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&system)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
