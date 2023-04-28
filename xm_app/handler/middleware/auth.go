package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/handler/companies"
	"github.com/sbuttigieg/test-xm/xm_app/handler/helpers"
)

func Auth(cfg *app.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, companies.ErrMsg{
				Code:  http.StatusUnauthorized,
				Error: companies.JWTError,
			})

			return
		}

		token := strings.Split(bearerToken, "Bearer ")[1]

		err := helpers.ValidateToken(cfg, token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, companies.ErrMsg{
				Code:  http.StatusUnauthorized,
				Error: companies.JWTError,
			})

			return
		}

		c.Next()
	}
}
