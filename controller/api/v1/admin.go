package v1

import (
	"errors"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloAdmin ...
// @Tags Admin
// @Summary Get hello admin
// @Produce application/json
// @Success 200 string hihihi
// @Router /admin/helloadmin [get]
func HelloAdmin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello admin",
	})
}

// RegistAdmin ...
// @Tags Admin
// @Summary User registration
// @Produce application/json
// @Param data body apiModel.UserRegistParam true "User registration parameters"
// @Success 200 {object} handler.Response{data=helperModel.PublicUser}
// @Router /admin/regist [post]
func RegistAdmin(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.UserRegistParam
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var admin bool = true
	user, err := helpermodel.RegistUser(params.Username, params.Password, admin)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.Response(http.StatusConflict, e.ERR_USER_EXIST, nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, user.Public())
	}
}

// LoginAdmin ...
// @Tags Admin
// @Summary User login test
// @Produce application/json
// @Param data body apiModel.UserRegistParam true "User login parameters"
// @Success 200 {object} handler.Response{data=helperModel.PublicUser}
// @Router /admin/login [post]
func LoginAdmin(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.LoginParam
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	user, err := helpermodel.Login(params.Username, params.Password)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, user.Public())
	}
}
