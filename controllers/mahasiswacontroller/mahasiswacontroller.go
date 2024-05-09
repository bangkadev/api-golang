package mahasiswacontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tentangkode/api-golang/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var mahasiswa []models.Mahasiswa

	models.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})
}

func Show(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := models.DB.First(&mahasiswa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})
}

func Create(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})
}

func Update(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&models.Mahasiswa{}).Where("id = ?", id).Updates(&mahasiswa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).Delete(&mahasiswa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
