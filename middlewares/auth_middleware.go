package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homeschooledu/consts"
)

var jwtSecret = []byte(consts.JWTSecretKey)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			consts.RespondWithError(c, -4, "Token not provided")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				consts.RespondWithError(c, -4, "Invalid token signature")
				c.Abort()
				return
			}

			consts.RespondWithError(c, -4, "Token parsing error")
			c.Abort()
			return
		}

		if !token.Valid {
			consts.RespondWithError(c, -4, "Invalid token")
			c.Abort()
			return
		}

		// 解析成功后继续处理请求
		c.Next()
	}
}
