package v1

import (
	"errors"
	"lil-helper-backend/jwt"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Regist ...
// @Tags Base
// @Summary user registration
// @Produce application/json
// @Param data body apimodel.UserRegistParam true "User registration parameters"
// @Success 200 {object} handler.Response{data=helpermodel.PublicUser}
// @Router /base/regist [post]
func Regist(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.UserRegistParam
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var admin bool = false
	user, err := helpermodel.RegistUser(params.Username, params.Password, admin)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.Response(http.StatusConflict, e.ERR_USER_EXIST, nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, user.Public())
	}
}

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
