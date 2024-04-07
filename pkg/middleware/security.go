package middleware

import (
	"errors"
	"muller-odontologia/pkg/web"
	"os"

	"github.com/gin-gonic/gin"
)

// Authentication provee seeguridad mediante la validacion de un token
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			c.Abort()
			return
		}
		c.Next()
	}
}
