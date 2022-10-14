package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"final-project/helpers"
	"final-project/models"
	"final-project/structs"
)

func Register(ctx *gin.Context) {
	// Initiate variable
	var user models.User

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&user)

	// Validate model
	_, err := user.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error creating user data: %s", err.Error()),
		})
		return
	}

	user.Password, err = helpers.Hash(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error creating user data: %s", err.Error()),
		})
		return
	}
	// Creating to database
	err = db.Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error creating user data: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, user)
}

func Login(ctx *gin.Context) {
	// Initiate variable
	var login structs.RequestLogin
	var user models.User

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&login)

	// Validate model
	_, err := login.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error creating user data: %s", err.Error()),
		})
		return
	}

	// Email Checking
	err = db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "F", "message": "User not found", "data": nil})
		return
	}

	// Verify password
	err = helpers.VerifyPassword(user.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "F", "message": "User not valid", "data": nil})
		return
	}

	// Generate JWT
	token, err := helpers.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		token: token,
	})
}

func UpdateUser(ctx *gin.Context) {
	// Initiate variable
	var update structs.RequestUpdateUser
	var user models.User

	userId := ctx.Param("userId")

	claims, _ := helpers.GetData(strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1])

	if claims.Id != userId {
		ctx.JSON(http.StatusForbidden, gin.H{
			"result": "Error updating user data: Forbidden",
		})
		return
	}

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&update)

	// Validate model
	_, err := update.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error creating user data: %s", err.Error()),
		})
		return
	}

	// Find user
	err = db.Model(models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "F", "message": "user not found", "data": nil})
		return
	}

	err = db.Model(&user).Updates(models.User{
		Username: update.Username,
		Email:    update.Email,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	// Setup database
	db := ctx.MustGet("db").(*gorm.DB)

	userId := ctx.Param("userId")

	user := models.User{}
	err := db.Model(models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "F", "message": "user not found", "data": nil})
		return
	}

	err = db.Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	//Response succes
	ctx.JSON(http.StatusOK, gin.H{"message": "Your Account Has been successfully deleted"})
}
