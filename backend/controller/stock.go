package controller

import (
	"net/http"

	"github.com/NantawatB63/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateStock(c *gin.Context) {

	var stock entity.Stock
	var shelf entity.Shelf
	var lot entity.Lot
	var product entity.Product

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", stock.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", stock.LotID).First(&lot); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id
	if tx := entity.DB().Where("id = ?", stock.ShelfID).First(&shelf); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}
	// 12: สร้าง WatchVideo
	st := entity.Stock{
		Lot:      lot,            // โยงความสัมพันธ์กับ Entity Resolution
		Product:  product,        // โยงความสัมพันธ์กับ Entity Video
		Shelf:    shelf,          // โยงความสัมพันธ์กับ Entity Playlist
		Quantity: stock.Quantity, // ตั้งค่าฟิลด์ watchedTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&st).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": st})
}

// GET /watchvideo/:id
func GetStock(c *gin.Context) {
	var stock entity.Stock
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&stock); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stock})
}

// GET /watch_videos
func ListStocks(c *gin.Context) {
	var stocks []entity.Stock
	if err := entity.DB().Preload("lot").Preload("shelf").Preload("product").Raw("SELECT * FROM stocks").Find(&stocks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stocks})
}

// DELETE /watch_videos/:id
func DeleteStock(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM stocks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateStock(c *gin.Context) {
	var stock entity.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", stock.ID).First(&stock); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})
		return
	}

	if err := entity.DB().Save(&stock).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stock})
}
