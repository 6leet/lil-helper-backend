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

func HelloUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

// RegistHelper ...
// @Tags Helper
// @Summary user registration
// @Produce application/json
// @Param data body apiModel.UserRegistParam true "User registration parameters"
// @Success 200 {object} handler.Response{data=helperModel.PublicUser}
// @Router /helper/regist [post]
func RegistHelper(c *gin.Context) {
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

// GetMission ...
// @Tags Helper
// @Summary get mission
// @Produce application/json
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /helper/mission [get]
func GetMission(c *gin.Context) {
	app := handler.Gin{C: c}

	// var user *helpermodel.User
	// if user = app.GetUser()

	mission := helpermodel.Mission{}
	app.Response(http.StatusOK, e.SUCCESS, mission)
}
