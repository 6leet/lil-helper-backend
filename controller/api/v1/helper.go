package v1

import (
	"errors"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
func GetMission(c *gin.Context) { // need: userID
	app := handler.Gin{C: c}

	// var user *helpermodel.User
	// if user = app.GetUser()

	mission := helpermodel.Mission{}
	app.Response(http.StatusOK, e.SUCCESS, mission)
}

// GetScreenshot ...
// @Tags Helper
// @Summary get screenshot
// @Produce application/json
// @Success 200 {object} handler.Response{data=helpermodel.Screenshot}
// @Router /helper/screenshot [get]
func GetScreenshot(c *gin.Context) { // need: userID, with date (today)
	app := handler.Gin{C: c}

	var user *helpermodel.User
	if user = app.GetUser(); user == nil {
		return
	}

	screenshot := helpermodel.Screenshot{
		UserID: user.UID,
	}
	app.Response(http.StatusOK, e.SUCCESS, screenshot)
}

// GetScreenshots ...
// @Tags Helper
// @Summary list helper screenshots
// @Produce application/json
// @Success 200 {object} handler.Response{data=apimodel.JsonObjectArray}
// @Router /helper/screenshots [get]
func GetScreenshots(c *gin.Context) { // need: userID, dateFrom, dateTo
	app := handler.Gin{C: c}

	// var user *helpermodel.User
	// if user = app.GetUser()

	screenshots := []helpermodel.Screenshot{}
	screenshot := helpermodel.Screenshot{
		Picture: "this/is/a/path/of/picture.jpg",
	}
	screenshots = append(screenshots, screenshot)
	screenshots = append(screenshots, screenshot)
	jsonArray := apimodel.NewJsonObjectArray(screenshots)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}

// CreateScreenshot ...
// @Tags Helper
// @Summary create screenshot
// @Produce application/json
// @Param data body apiModel.SetScreenshotParams true "set screenshot params"
// @Success 200 {object} handler.Response{data=helpermodel.Screenshot}
// @Router /helper/screenshot [post]
func CreateScreenshot(c *gin.Context) { // need: userID, with date (today)
	app := handler.Gin{C: c}

	var params apimodel.SetScreenshotParams
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	// var user *helpermodel.User
	// if user = app.GetUser()

	screenshot := helpermodel.Screenshot{
		MissionID: params.MissionID,
		Picture:   params.Picture,
		Audit:     false,
		Approve:   false,
		Date:      time.Now(),
	}
	app.Response(http.StatusOK, e.SUCCESS, screenshot)
}

// UpdateScreenshot ...
// @Tags Helper
// @Summary update screenshot
// @Produce application/json
// @Param uid path string true "screenshot uid"
// @Param data body apimodel.SetScreenshotParams true "set screenshot params"
// @Success 200 {object} handler.Response{data=helpermodel.Screenshot}
// @Router /helper/screenshots/{uid} [post]
func UpdateScreenshot(c *gin.Context) { // need: userID
	app := handler.Gin{C: c}
	screenshotID := c.Param("uid")
	var params apimodel.SetScreenshotParams

	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// var user *helpermodel.User
	// if user = app.GetUser()

	screenshot := helpermodel.Screenshot{
		UID:       screenshotID,
		MissionID: params.MissionID,
		Picture:   params.Picture,
		Audit:     false,
		Approve:   false,
		Date:      time.Now(),
	}
	app.Response(http.StatusOK, e.SUCCESS, screenshot)
}

// DeleteScreenshot ...
// @Tags Helper
// @Summary delete screenshot
// @Produce application/json
// @Param uid path string true "screenshot uid"
// @Success 200 {object} handler.Response
// @Router /helper/screenshots/{uid} [delete]
func DeleteScreenshot(c *gin.Context) {
	app := handler.Gin{C: c}
	screenshotID := c.Param("uid")
	app.Response(http.StatusOK, e.SUCCESS, screenshotID)
}
