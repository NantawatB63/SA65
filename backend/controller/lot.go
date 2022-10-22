package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NantawatB63/sa-65-example/entity"
)

// POST /lots
func CreateLot(c *gin.Context) {
	var lot entity.Lot
	if err := c.ShouldBindJSON(&lot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&lot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": lot})
}

// GET /lot/:id
func GetLot(c *gin.Context) {
	var lot entity.Lot
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&lot); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lot not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lot})
}

// GET /lots
func ListLot(c *gin.Context) {
	var lot []entity.Lot
	if err := entity.DB().Raw("SELECT * FROM lot").Scan(&lot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lot})
}

// DELETE /lots/:id
func DeleteLot(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM lot WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lot not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /lots
func UpdateLot(c *gin.Context) {
	var lot entity.Lot
	if err := c.ShouldBindJSON(&lot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", lot.ID).First(&lot); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lot not found"})
		return
	}

	if err := entity.DB().Save(&lot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lot})
}