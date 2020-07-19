package middleware

import (
	"lil-helper-backend/jwt"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return jwt.Jwt.MiddlewareFunc()
}
