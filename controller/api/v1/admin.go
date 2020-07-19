package v1

import (
	"errors"
	"lil-helper-backend/hashids"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegistAdmin ...
// @Tags Admin
// @Summary user registration
// @Produce application/json
// @Param data body apiModel.UserRegistParam true "User registration parameters"
// @Success 200 {object} handler.Response{data=helperModel.PublicUser}
// @Router /admin/regist [post]
func RegistAdmin(c *gin.Context) { // done
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

// CreateMission ...
// @Tags Admin
// @Summary create mission
// @Produce application/json
// @Param data body apiModel.SetMissionParam true "set mission params"
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /admin/mission [post]
func CreateMission(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.SetMissionParam

	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	mission := helpermodel.Mission{
		Content: params.Content,
		Picture: params.Picture,
		// Weight:  []int{1, 2, 3, 4, 5},
		Score:  10,
		Active: params.Active,
	}
	app.Response(http.StatusOK, e.SUCCESS, mission)
}

// GetMissions ...
// @Tags Admin
// @Summary list missions
// @Produce application/json
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/missions [get]
func GetMissions(c *gin.Context) {
	app := handler.Gin{C: c}

	missions := []helpermodel.Mission{}
	mission := helpermodel.Mission{}
	missions = append(missions, mission)
	missions = append(missions, mission)
	jsonArray := apimodel.NewJsonObjectArray(missions)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}

// GetMissionsByDate ...
// @Tags Admin
// @Summary list missions by date
// @Produce application/json
// @Params date path string true "mission date"
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/missions/{date} [get]
func GetMissionsByDate(c *gin.Context) {
	app := handler.Gin{C: c}

	date := c.Param("date")

	layout := "2006-01-02T15:04:05.000Z"
	date = "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, date)

	if err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	missions := []helpermodel.Mission{}
	mission := helpermodel.Mission{
		Date: t,
	}
	missions = append(missions, mission)
	missions = append(missions, mission)
	jsonArray := apimodel.NewJsonObjectArray(missions)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}

// UpdateMission ...
// @Tags Admin
// @Summary update mission
// @Produce application/json
// @Param uid path string true "mission uid"
// @Param data body apiModel.SetMissionParam true "set mission params"
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /admin/missions/{uid} [post]
func UpdateMission(c *gin.Context) {
	app := handler.Gin{C: c}
	missionID := c.Param("uid")
	var params apimodel.SetMissionParam

	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	mission := helpermodel.Mission{
		UID:     missionID,
		Content: params.Content,
		Picture: params.Picture,
		// Weight:  []int{1, 2, 3, 4, 5},
		Score:  10,
		Active: params.Active,
	}
	app.Response(http.StatusOK, e.SUCCESS, mission)
}

// DeleteMission ...
// @Tags Admin
// @Summary delete mission
// @Produce application/json
// @Param uid path string true "mission uid"
// @Success 200 {object} handler.Response
// @Router /admin/missions/{uid} [delete]
func DeleteMission(c *gin.Context) {
	app := handler.Gin{C: c}
	missionID := c.Param("uid")
	app.Response(http.StatusOK, e.SUCCESS, missionID)
}

// GetScreenshots ...
// @Tags Admin
// @Summary list screenshots
// @Produce application/json
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/screenshots [get]
func GetAllScreenshots(c *gin.Context) {
	app := handler.Gin{C: c}

	var screenshots []helpermodel.Screenshot
	var screenshot helpermodel.Screenshot
	screenshots = append(screenshots, screenshot)
	screenshots = append(screenshots, screenshot)
	jsonArray := apimodel.NewJsonObjectArray(screenshots)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}

// SetScreenshotApprove ...
// @Tags Admin
// @Summary audit screenshot
// @Produce application/json
// @Param data body apimodel.AuditScreenshotParams true "audit screenshot params"
// @Success 200 {object} handler.Response{data=helpermodel.Screenshot}
// @Router /admin/setscreenshotapprove [post]
func SetScreenshotApprove(c *gin.Context) {
	app := handler.Gin{C: c}

	var params apimodel.AuditScreenshotParams
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	screenshot := helpermodel.Screenshot{
		UserID:    params.UserID,
		MissionID: params.MissionID,
		Approve:   params.Approve,
		Audit:     true,
		Picture:   "this/is/a/path/of/picture.jpg",
	}
	app.Response(http.StatusOK, e.SUCCESS, screenshot)
	//set score
}

// GetHelpers ...
// @Tags Admin
// @Summary list helpers
// @Produce application/json
// @Param active query bool false "flag to query active user only (default: true)"
// @Param all query bool false "flag to query all users (default: false)"
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/helpers [get]
func GetHelpers(c *gin.Context) {
	app := handler.Gin{C: c}

	var active, all = true, false
	if activestr, ok := c.GetQuery("active"); ok && activestr == "false" {
		active = false
	}
	if allstr, ok := c.GetQuery("all"); ok && allstr == "true" {
		all = true
	}
	if helpers, err := helpermodel.GetUsers(active, false, all, false); err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		var publicHelpers []helpermodel.PublicUser
		for _, u := range helpers {
			publicHelpers = append(publicHelpers, u.Public())
		}
		jsonArray := apimodel.NewJsonObjectArray(publicHelpers)
		app.Response(http.StatusOK, e.SUCCESS, jsonArray)
	}
}

// BanHelper ...
// @Tags Admin
// @Summary ban helper
// @Produce application/json
// @Param uid path string true "User uid"
// @Success 200 {object} handler.Response
// @Router /admin/helpers/{uid} [delete]
func BanHelper(c *gin.Context) {
	app := handler.Gin{C: c}

	userID := c.Param("uid")
	if id, err := hashids.DecodeUserUID(userID); err != nil {
		app.Response(http.StatusBadRequest, e.ERR_INVALID_USER_UID, nil)
	} else if err := helpermodel.BanUser(id); err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	app.Response(http.StatusOK, e.SUCCESS, userID)
}
