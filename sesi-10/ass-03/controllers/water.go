package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"ass-03/models"
)

func Shuffle(ctx *gin.Context) {
	// Initiate variable
	water := &models.TbWater{
		Water: rand.Intn(99) + 1,
		Wind:  rand.Intn(99) + 1,
	}

	b, err := json.Marshal(water)

	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Error shuffling data: %s", err.Error()),
			})
			return
		}
	}

	f, err := os.Create("data.json")

	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Error shuffling data: %s", err.Error()),
			})
			return
		}
	}

	defer f.Close()

	_, err = f.WriteString(string(b))

	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Error shuffling data: %s", err.Error()),
			})
			return
		}
	}

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"status": water,
	})
}
