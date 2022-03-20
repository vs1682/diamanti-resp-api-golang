package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"restapi/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isTokenValid bool
		var data interface{}

		isTokenValid = true
		token := c.Query("token")
		if token == "" {
			isTokenValid = false
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				isTokenValid = false
			}
		}

		if !isTokenValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "Unauthorized",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
