package controllers

import (
	"go_restapi_gin/models"
	"net/http"

	"github.com/aryansh13/go-restapi-gin/database"
	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Lakukan validasi token JWT di sini (gunakan middleware)

	// Isi data pengguna dari token
	userID := 1 // Misalnya, asumsikan userID dari token adalah 1

	newPhoto := models.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   userID,
	}

	// Simpan newPhoto ke basis data menggunakan Gorm
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	result := db.Create(&newPhoto)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully"})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	result := db.Find(&photos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func UpdatePhotoByID(c *gin.Context) {
	photoID := c.Param("photoId")

	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var existingPhoto models.Photo
	result := db.First(&existingPhoto, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	existingPhoto.Title = updatedPhoto.Title
	existingPhoto.Caption = updatedPhoto.Caption
	existingPhoto.PhotoURL = updatedPhoto.PhotoURL

	updateResult := db.Save(&existingPhoto)
	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func DeletePhotoByID(c *gin.Context) {
	photoID := c.Param("photoId")

	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var existingPhoto models.Photo
	result := db.First(&existingPhoto, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	deleteResult := db.Delete(&existingPhoto)
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
