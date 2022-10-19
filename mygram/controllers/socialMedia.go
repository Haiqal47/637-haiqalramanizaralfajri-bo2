package controllers

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/models"
	"final-project/structs"
	"fmt"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create Social Media godoc
// @Summary Create Social Media
// @Description Create data Social Media to database
// @Tags socialMedias
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param socialMedia			body			structs.RequestCreateSocialMedia	true "Create Social Media"
// @Success 200 {object} structs.ResponseCreateSocialMedia
// @Router /socialmedias [post]
func CreateSocialMedia(ctx *gin.Context) {
	// Initialize variable
	var socialMedia models.SocialMedia

	// get db
	db := ctx.MustGet("db").(*gorm.DB)

	// get body
	err := json.NewDecoder(ctx.Request.Body).Decode(&socialMedia)

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Error decode request body",
		})
		return
	}

	socialMedia.UserId = claims.ID

	err = db.Create(&socialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating social media to database",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserId,
		"created_at":       socialMedia.CreatedAt,
	})
}

// Get Social Media godoc
// @Summary Get All Social Media
// @Description Get All Social Media with user data
// @Tags socialMedias
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} []structs.ResponseSocialMedia
// @Router /socialmedias [get]
func GetSocialMedias(ctx *gin.Context) {
	var socialMedias []models.SocialMedia

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// get data list
	err := db.Preload("User").Find(&socialMedias).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting social media",
		})
		return
	}

	var response []structs.ResponseSocialMedia
	m := dto.Mapper{}
	m.Map(&response, socialMedias)

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": response,
	})
}

// Update Social media godoc
// @Summary Update Social media
// @Description Update data Social media
// @Tags socialMedias
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param socialMediaId		path 		 	int 															true "ID"
// @Param socialMedia			body			structs.RequestUpdateSocialMedia	true "Update Social Media"
// @Success 200 {object} structs.ResponseUpdateSocialMedia
// @Router /socialmedias/{socialMediaId} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	// Initiate variable
	var update models.SocialMedia
	var socialMedia models.SocialMedia

	socialMediaId := ctx.Param("socialMediaId")

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&update)

	// Validate model
	_, err := update.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Validation error: %s", err.Error()),
		})
		return
	}

	// Find socialMedia
	err = db.Model(models.SocialMedia{}).Where("id = ?", socialMediaId).First(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social Media Not Found",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.ID != socialMedia.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this social media data",
		})
		return
	}

	err = db.Model(&socialMedia).Updates(update).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating social media to database",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserId,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

// Delete Social Media godoc
// @Summary Delete Social Media
// @Description Delete data Social Media
// @Tags socialMedias
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param socialMediaId			path 			int 		true 	"ID"
// @Success 200 {object} structs.Response
// @Router /socialmedias/{socialMediaId} [delete]
func DeleteSocialMedia(ctx *gin.Context) {

	socialMediaId := ctx.Param("socialMediaId")

	// Setup database
	db := ctx.MustGet("db").(*gorm.DB)

	socialMedia := models.SocialMedia{}
	err := db.Model(models.SocialMedia{}).Where("id = ?", socialMediaId).First(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Comment Not Found",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.ID != socialMedia.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this comment data",
		})
		return
	}

	err = db.Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting comment",
		})
		return
	}

	//Response succes
	ctx.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted"})
}
