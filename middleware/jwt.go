package middleware

import (
	"lil-helper-backend/jwt"

	"github.com/gin-gonic/gin"
)

func HelperJwt() gin.HandlerFunc {
	return jwt.HelperJwt.MiddlewareFunc()
}

func AdminJwt() gin.HandlerFunc {
	return jwt.AdminJwt.MiddlewareFunc()
}
