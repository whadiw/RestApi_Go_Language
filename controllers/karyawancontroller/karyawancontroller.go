package karyawancontroller

import (
	"encoding/json"
	"go-crud-restapi/database"
	"go-crud-restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	//find all karyawan yang masih active
	var karyawan []models.Karyawan
	inActive := "N"

	database.DB.Where("in_active = ?", inActive).Find(&karyawan)
	c.JSON(http.StatusOK, gin.H{"karyawan": karyawan})
}

func GetById(c *gin.Context) {
	var karyawan []models.Karyawan
	id := c.Param("id")

	if err := database.DB.First(&karyawan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"karyawan": karyawan})
}

func Create(c *gin.Context) {
	var karyawan models.Karyawan

	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"Karyawan": karyawan})
}

func Update(c *gin.Context) {
	var karyawan models.Karyawan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&karyawan).Where("id = ?", id).Updates(&karyawan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengubah data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbaharui"})
}

func Delete(c *gin.Context) {
	var karyawan models.Karyawan

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if database.DB.Delete(&karyawan, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
