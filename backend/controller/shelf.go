package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NantawatB63/sa-65-example/entity"
)

// POST /shelfs
func CreateShelf(c *gin.Context) {
	var shelf entity.Shelf
	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": shelf})
}

// GET /shelf/:id
func GetShelf(c *gin.Context) {
	var shelf entity.Shelf
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&shelf); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelf})
}

// GET /shelfs
func ListShelf(c *gin.Context) {
	var shelf []entity.Shelf
	if err := entity.DB().Raw("SELECT * FROM resolutions").Scan(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelf})
}

// DELETE /shelfs/:id
func DeleteShelf(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM resolutions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /shelfs
func UpdateShelf(c *gin.Context) {
	var shelf entity.Shelf
	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", shelf.ID).First(&shelf); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&shelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelf})
}