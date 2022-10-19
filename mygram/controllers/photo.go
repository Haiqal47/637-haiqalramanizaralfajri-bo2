package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/structs"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dranikpg/dto-mapper"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create Photo godoc
// @Summary Create Photo
// @Description Create data photo to database
// @Tags photos
// @Accept  multipart/form-data
// @Produce  json
// @Security BearerAuth
// @Param photo	 					formData file 	true "Photo"
// @Param title		 				formData string true "Title"
// @Param caption					formData string true "Caption"
// @Success 200 {object} structs.ResponseCreatePhoto
// @Router /photos [post]
func CreatePhoto(ctx *gin.Context) {

	var photo models.Photo

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	file, err := ctx.FormFile("photo")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
		return
	}

	fileHeader := make([]byte, 512)
	fileOpen, err := file.Open()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
		return
	}

	defer fileOpen.Close()

	if _, err = fileOpen.Read(fileHeader); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
		return
	}

	filetype := http.DetectContentType(fileHeader)

	if !strings.Contains(filetype, "image/") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File not image",
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Image not found",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	current_time := time.Now()
	filename := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second()) + claims.Username + filepath.Ext(file.Filename)

	// Parsing from request body to Order Model
	ctx.Bind(&photo)

	baseUrl := location.Get(ctx)
	baseUrl.Path = ""
	baseUrl.RawQuery = ""
	baseUrl.Fragment = ""
	photo.PhotoUrl = baseUrl.String() + "/content/" + filename
	photo.UserId = claims.ID

	_, err = photo.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Validation error: %s", err.Error()),
		})
		return
	}

	err = db.Create(&photo).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating photo to database",
		})
		return
	}

	if _, err := os.Stat("./uploaded"); os.IsNotExist(err) {
		err := os.Mkdir("./uploaded", os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error sending file to server",
			})
			return
		}
	}

	if err := ctx.SaveUploadedFile(file, "./uploaded/"+filename); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending file to server",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserId,
		"created_at": photo.CreatedAt,
	})
}

// Get Photos godoc
// @Summary Get All Photo
// @Description Get All Photo with user data
// @Tags photos
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} []structs.ResponsePhotos
// @Router /photos/ [get]
func GetPhotos(ctx *gin.Context) {
	var photos []models.Photo

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// get data list
	err := db.Preload("User").Find(&photos).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting photo from database",
		})
		return
	}

	var response []structs.ResponsePhotos
	m := dto.Mapper{}
	m.Map(&response, photos)

	ctx.JSON(http.StatusOK, response)
}

// Update Photo godoc
// @Summary Update Photo
// @Description Update data Photo
// @Tags photos
// @Accept  multipart/form-data
// @Produce  json
// @Security BearerAuth
// @Param photoId 				path 		 int 		true "ID"
// @Param photo	 					formData file 	true "Photo"
// @Param title		 				formData string true "Title"
// @Param caption					formData string true "Caption"
// @Success 200 {object} structs.ResponseUpdateComment
// @Router /photos/{photoId} [put]
func UpdatePhoto(ctx *gin.Context) {
	// Initiate variable
	var update models.Photo
	var photo models.Photo

	photoId := ctx.Param("photoId")

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	file, err := ctx.FormFile("photo")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
	}

	fileHeader := make([]byte, 512)
	fileOpen, err := file.Open()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
		return
	}

	defer fileOpen.Close()

	if _, err = fileOpen.Read(fileHeader); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Image not found",
		})
		return
	}

	filetype := http.DetectContentType(fileHeader)

	if !strings.Contains(filetype, "image/") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File not image",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	current_time := time.Now()
	filename := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second()) + claims.Username + filepath.Ext(file.Filename)

	// Parsing from request body to Order Model
	ctx.Bind(&update)

	baseUrl := location.Get(ctx)
	baseUrl.Path = ""
	baseUrl.RawQuery = ""
	baseUrl.Fragment = ""
	update.PhotoUrl = baseUrl.String() + "/content/" + filename

	// Validate model
	_, err = update.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Validation Error: %s", err.Error()),
		})
		return
	}

	// Find photo
	err = db.Model(models.Photo{}).Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
		return
	}

	if claims.ID != photo.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this photo data",
		})
		return
	}

	err = db.Model(&photo).Updates(update).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Error updating photo to database"})
		return
	}

	if _, err := os.Stat("./uploaded"); os.IsNotExist(err) {
		err := os.Mkdir("./uploaded", os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error creating photo data: %s", err.Error()),
			})
			return
		}
	}

	if err := ctx.SaveUploadedFile(file, "./uploaded/"+filename); err != nil {
		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserId,
		"updated_at": photo.UpdatedAt,
	})
}

// Delete Photo godoc
// @Summary Delete photo
// @Description Delete data photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param photoId 					path 			int 		true 	"ID"
// @Success 200 {object} structs.Response
// @Router /photos/{photoId} [delete]
func DeletePhoto(ctx *gin.Context) {

	photoId := ctx.Param("photoId")

	// Setup database
	db := ctx.MustGet("db").(*gorm.DB)

	photo := models.Photo{}
	err := db.Model(models.Photo{}).Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.ID != photo.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot delete this photo data",
		})
		return
	}

	err = db.Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Error deleting photo"})
		return
	}

	//Response succes
	ctx.JSON(http.StatusOK, gin.H{"message": "Your photo Has been successfully deleted"})
}
