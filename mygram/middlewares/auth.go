package middlewares

import (
	"errors"
	"final-project/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(token, "Bearer")

		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Unauthorized: %s", errors.New("sign in to proceed")),
			})
			return
		}

		token = strings.Split(token, " ")[1]
		verifyToken, err := helpers.ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Unauthorized: %s", err.Error()),
			})
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
