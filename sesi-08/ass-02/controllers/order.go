package controllers

import (
	"ass-02/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.Order true "Create order"
// @Success 200 {object} models.Order
// @Router /orders [post]
func CreateOrder(ctx *gin.Context) {
	// Initiate variable
	var order models.Order

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&order)

	// Creating to database
	err := db.Model(models.Order{}).Create(&order).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error creating order data: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, order)
}

// GetOrders godoc
// @Summary Get Details of all orders
// @Description Get Details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Order
// @Router /orders [get]
func GetOrders(ctx *gin.Context) {
	// Initiate variable
	var orders []models.Order

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Get from database
	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error getting orders data: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, orders)
}

// GetOrder godoc
// @Summary Get Details order by id
// @Description Get Details order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {object} models.Order
// @Router /orders/{orderId} [get]
func GetOrder(ctx *gin.Context) {
	// Initiate variable
	var order models.Order

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Get Params
	id := ctx.Param("orderId")

	// Get by id from database
	err := db.Preload("Items").Where("order_id = ?", id).First(&order).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error getting order data by id: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, order)
}

// UpdateOrder godoc
// @Summary Update data order where orderId
// @Description Update data order where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {object} models.Order
// @Router /orders/{orderId} [put]
func UpdateOrder(ctx *gin.Context) {
	// Initiate variable
	var order models.Order
	var newOrder models.Order

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Get Params
	id := ctx.Param("orderId")

	// Parsing from request body to Order Model
	json.NewDecoder(ctx.Request.Body).Decode(&newOrder)

	// Get by id from database
	err := db.Preload("Items").Where("order_id = ?", id).First(&order).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"result": "Order Not Found",
		})
		return
	}

	// Update data to database
	order.Items = newOrder.Items
	order.CustomerName = newOrder.CustomerName
	order.OrderedAt = newOrder.OrderedAt
	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error updating order data by id: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, order)
}

// DeleteOrder godoc
// @Summary Delete data order where orderId
// @Description Delete data order where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func DeleteOrder(ctx *gin.Context) {
	// Initiate variable
	var order models.Order

	// Setup Database
	db := ctx.MustGet("db").(*gorm.DB)

	// Get Params
	id := ctx.Param("orderId")
	ids, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "Id Not Number",
		})
		return
	}

	// Get by id from database
	err = db.First(&order, ids).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"result": "Order Not Found",
		})
		return
	}

	// Delete data to database
	err = db.Delete(&order).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error deleting order data by id: %s", err.Error()),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Order deleted",
	})
}
