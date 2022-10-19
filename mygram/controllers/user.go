package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"final-project/helpers"
	"final-project/models"
	"final-project/structs"

	"github.com/gin-contrib/location"
)

// Register User godoc
// @Summary Registration User
// @Description Register New User with photo profile
// @Tags users
// @Accept  multipart/form-data
// @Produce  json
// @Param profile_image 	formData file 	true "Profile Image"
// @Param username 				formData string true "Username"
// @Param password 				formData string true "Password"
// @Param email 					formData string true "Email"
// @Param age 						formData int		true "Age"
// @Success 200 {object} structs.ResponseUserRegister
// @Router /users/register [post]
func Register(ctx *gin.Context) {
	// Initiate variable
	var user models.User

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Get Photo from Form Data
	file, err := ctx.FormFile("profile_image")

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

	// Parsing from request body to Order Model
	ctx.Bind(&user)

	current_time := time.Now()
	filename := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second()) + user.Username + "Profile" + filepath.Ext(file.Filename)

	baseUrl := location.Get(ctx)
	baseUrl.Path = ""
	baseUrl.RawQuery = ""
	baseUrl.Fragment = ""
	user.ProfileImageUrl = baseUrl.String() + "/content/" + filename

	// Validate model
	_, err = user.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Validation Error: %s", err.Error()),
		})
		return
	}

	user.Password, err = helpers.Hash(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing user password",
		})
		return
	}

	// Creating to database
	err = db.Create(&user).Error

	if err != nil {
		if strings.Contains(err.Error(), "already used") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user to database",
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
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

// Login User godoc
// @Summary Login User
// @Description Login User using email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body structs.RequestLogin true "Login user"
// @Success 200 {object} structs.ResponseUserLogin
// @Router /users/login [post]
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
			"message": fmt.Sprintf("Validation error: %s", err.Error()),
		})
		return
	}

	// Email Checking
	err = db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Verify password
	err = helpers.VerifyPassword(user.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong email or password"})
		return
	}

	// Generate JWT
	token, err := helpers.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error generate JWT"})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Update User godoc
// @Summary Update User
// @Description Update data User
// @Tags users
// @Accept  multipart/form-data
// @Produce  json
// @Security BearerAuth
// @Param userId 					path 		 int 		true "ID"
// @Param profile_image 	formData file 	true "Profile Image"
// @Param username 				formData string true "Username"
// @Param email 					formData string true "Email"
// @Success 200 {object} structs.ResponseUserUpdate
// @Router /users/{userId} [put]
func UpdateUser(ctx *gin.Context) {
	// Initiate variable
	var update structs.RequestUpdateUser
	var user models.User

	userId := ctx.Param("userId")

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.Id != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot update this user data",
		})
		return
	}

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	file, err := ctx.FormFile("profile_image")

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

	// Parsing from request body to Order Model
	ctx.Bind(&update)

	// Validate model
	_, err = update.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Validation error: %s", err.Error()),
		})
		return
	}

	// Find user
	err = db.Model(models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	current_time := time.Now()
	filename := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second()) + user.Username + "Profile" + filepath.Ext(file.Filename)

	baseUrl := location.Get(ctx)
	baseUrl.Path = ""
	baseUrl.RawQuery = ""
	baseUrl.Fragment = ""
	update.ProfileImageUrl = baseUrl.String() + "/content/" + filename

	err = db.Model(&user).Updates(models.User{
		Username: update.Username,
		Email:    update.Email,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating user to database",
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
	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

// Delete User godoc
// @Summary Delete user
// @Description Delete data user
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param userId 					path 			int 		true 	"ID"
// @Success 200 {object} structs.Response
// @Router /users/{userId} [delete]
func DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("userId")

	// Setup database
	db := ctx.MustGet("db").(*gorm.DB)

	// get userData
	claims := ctx.MustGet("userData").(*helpers.JWTClaim)

	if claims.Id != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're cannot delete this user data",
		})
		return
	}

	user := models.User{}
	err := db.Model(models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	err = db.Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting user",
		})
		return
	}

	//Response succes
	ctx.JSON(http.StatusOK, gin.H{"message": "Your Account Has been successfully deleted"})
}
