package v1

import (
	"lil-helper-backend/jwt"

	"github.com/gin-gonic/gin"
)

// Login ...
// @Tags Base
// @Summary User login
// @Produce application/json
// @Param data body apimodel.LoginParam true "user login parameters"
// @Success 200 {object} handler.Response{data=apimodel.LoginResData}
// @Router /base/login [post]
func Login(c *gin.Context) {
	jwt.Jwt.LoginHandler(c)
}

// RefreshToken ...
// @Tags Base
// @Summary User refresh token
// @Produce application/json
// @Success 200 {object} handler.Response{data=apimodel.LoginResData}
// @Router /base/refresh-token [get]
func RefreshToken(c *gin.Context) {
	jwt.Jwt.RefreshHandler(c)
}
