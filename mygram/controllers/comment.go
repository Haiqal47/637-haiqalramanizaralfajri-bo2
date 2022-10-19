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

// Create Comment godoc
// @Summary Create Comment
// @Description Create data comment to database
// @Tags comments
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param comment	 				body			structs.RequestCreateComment	true "Create Comment"
// @Success 200 {object} structs.ResponseCreateComment
// @Router /comments [post]
func CreateComment(ctx *gin.Context) {
	// Initialize variable
	var comment models.Comment

	// get db
	db := ctx.MustGet("db").(*gorm.DB)

	// get body
	err := json.NewDecoder(ctx.Request.Body).Decode(&comment)

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Error decode request body",
		})
		return
	}

	comment.UserId = claims.ID

	var photoExist models.Photo
	err = db.Model(models.Photo{}).Where("id = ?", comment.PhotoId).First(&photoExist).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo with this id Not Found",
		})
		return
	}

	err = db.Create(&comment).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating comment to database",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoId,
		"user_id":    comment.UserId,
		"created_at": comment.CreatedAt,
	})
}

// Get Comments godoc
// @Summary Get All Comment
// @Description Get All Comment with user data and photo data
// @Tags comments
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} []structs.ResponseComments
// @Router /comments [get]
func GetComments(ctx *gin.Context) {
	var comments []models.Comment

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// get data list
	err := db.Preload("User").Preload("Photo").Find(&comments).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting comment",
		})
		return
	}

	var response []structs.ResponseComments
	m := dto.Mapper{}
	m.Map(&response, comments)

	ctx.JSON(http.StatusOK, response)
}

// Update Comment godoc
// @Summary Update Comment
// @Description Update data Comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param commentId 			path 		 	int 													true "ID"
// @Param comment	 				body			structs.RequestUpdateComment	true "Update Comment"
// @Success 200 {object} structs.ResponseUpdateComment
// @Router /photos/{commentId} [put]
func UpdateComment(ctx *gin.Context) {
	// Initiate variable
	var update models.Comment
	var comment models.Comment

	commentId := ctx.Param("commentId")

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

	// Find comment
	err = db.Model(models.Comment{}).Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Comment Not Found",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.ID != comment.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this comment data",
		})
		return
	}

	var photoExist models.Photo
	err = db.Model(models.Photo{}).Where("id = ?", update.PhotoId).First(&photoExist).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo with this id Not Found",
		})
		return
	}

	err = db.Model(&comment).Updates(update).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating comment to database",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"user_id":    comment.UserId,
		"photo_id":   comment.PhotoId,
		"message":    comment.Message,
		"updated_at": comment.UpdatedAt,
	})
}

// Delete Comment godoc
// @Summary Delete Comment
// @Description Delete data comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param commentId 				path 			int 		true 	"ID"
// @Success 200 {object} structs.Response
// @Router /photos/{commentId} [delete]
func DeleteComment(ctx *gin.Context) {

	commentId := ctx.Param("commentId")

	// Setup database
	db := ctx.MustGet("db").(*gorm.DB)

	comment := models.Comment{}
	err := db.Model(models.Comment{}).Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Comment Not Found",
		})
		return
	}

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.ID != comment.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this comment data",
		})
		return
	}

	err = db.Delete(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting comment",
		})
		return
	}

	//Response succes
	ctx.JSON(http.StatusOK, gin.H{"message": "Your comment Has been successfully deleted"})
}
