package v1

import (
	"encoding/json"
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
	user, err := helpermodel.RegistUser(params.Username, params.Password, params.Email, admin)
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
// @Param data body apiModel.SetMissionParams true "set mission params"
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /admin/mission [post]
func CreateMission(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.SetMissionParams

	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var user *helpermodel.User
	if user = app.GetUser(); user == nil {
		return
	}
	userID, err := hashids.DecodeUserUID(user.UID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_INVALID_USER_UID, nil)
		return
	}

	// weightstr := apimodel.IntSliceToString(params.Weight)
	weightjson, _ := json.Marshal(params.Weight)
	// reverse operation json.Unmarshal(string(weightjson))
	// var weight []int
	// err := json.Unmarshal([]byte(string(weightjson)), &weight)
	// if err != nil {
	// 	fmt.Println("error on unmarshal")
	// }
	// fmt.Println(weight)
	mission, err := helpermodel.CreateMission(userID, params.Content, params.Picture, string(weightjson), params.Score, params.Activeat, params.Inactiveat)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, mission.Public())
	}
}

// GetMissions ...
// @Tags Admin
// @Summary list missions by date(optional), else list today's missions
// @Produce application/json
// @Param datefrom query string false "mission date from"
// @Param dateto query string false "mission date to"
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/missions [get]
func GetMissions(c *gin.Context) {
	app := handler.Gin{C: c}

	var dateFrom, dateTo string

	if dateFromq, ok := c.GetQuery("datefrom"); ok {
		dateFrom = dateFromq
	}
	if dateToq, ok := c.GetQuery("dateto"); ok {
		dateTo = dateToq
	}
	if dateFrom == "" || dateTo == "" {
		dateFrom = time.Now().String()[0:10]
		dateTo = time.Now().AddDate(0, 0, 1).String()[0:10]
	}

	missions, err := helpermodel.GetMissionsByDate(dateFrom, dateTo)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		var publicMissions []helpermodel.PublicMission
		for _, m := range missions {
			publicMissions = append(publicMissions, m.Public())
		}
		jsonArray := apimodel.NewJsonObjectArray(publicMissions)
		app.Response(http.StatusOK, e.SUCCESS, jsonArray)
	}
}

// UpdateMission ...
// @Tags Admin
// @Summary update mission
// @Produce application/json
// @Param uid path string true "mission uid"
// @Param data body apiModel.SetMissionParams true "set mission params"
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /admin/missions/{uid} [post]
func UpdateMission(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.SetMissionParams

	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	missionUID := c.Param("uid")
	missionID, err := hashids.DecodeMissionUID(missionUID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_MISSION, nil)
		return
	}

	weightjson, _ := json.Marshal(params.Weight)

	mission, err := helpermodel.UpdateMission(missionID, params.Content, params.Picture, string(weightjson), params.Score, params.Active, params.Activeat, params.Inactiveat)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, mission.Public())
	}
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
	missionUID := c.Param("uid")
	missionID, err := hashids.DecodeMissionUID(missionUID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_MISSION, nil)
		return
	}
	err = helpermodel.DeleteMission(missionID)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, nil)
	}
}

// GetScreenshots ...
// @Tags Admin
// @Summary list screenshots
// @Produce application/json
// @Param datefrom query string false "screenshot date from"
// @Param dateto query string false "screenshot date to"
// @Param audit query bool true "if screenshot auditted (default: false)"
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/screenshots [get]
func GetAllScreenshots(c *gin.Context) {
	app := handler.Gin{C: c}

	var dateFrom, dateTo string
	if dateFromq, ok := c.GetQuery("datefrom"); ok {
		dateFrom = dateFromq
	}
	if dateToq, ok := c.GetQuery("dateto"); ok {
		dateTo = dateToq
	}
	if dateFrom == "" || dateTo == "" {
		dateFrom = time.Now().String()[0:10]
		dateTo = time.Now().AddDate(0, 0, 1).String()[0:10]
	}

	var audit = false
	if auditstr, ok := c.GetQuery("audit"); ok && auditstr == "true" {
		audit = true
	}

	screenshots, err := helpermodel.GetScreenshotsByDate(0, dateFrom, dateTo, audit, true)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		var publicScreenshots []helpermodel.PublicScreenshot
		for _, s := range screenshots {
			publicScreenshots = append(publicScreenshots, s.Public())
		}
		jsonArray := apimodel.NewJsonObjectArray(publicScreenshots)
		app.Response(http.StatusOK, e.SUCCESS, jsonArray)
	}
}

// SetScreenshotApprove ...
// @Tags Admin
// @Summary audit screenshot
// @Produce application/json
// @Param uid path string true "screenshot uid"
// @Param data body apimodel.AuditScreenshotParams true "audit screenshot params"
// @Success 200 {object} handler.Response{data=helpermodel.Screenshot}
// @Router /admin/screenshots/{uid} [post]
func SetScreenshotApprove(c *gin.Context) {
	app := handler.Gin{C: c}
	var params apimodel.AuditScreenshotParams
	if err := c.BindJSON(&params); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	screenshotUID := c.Param("uid")
	screenshotID, err := hashids.DecodeScreenshotUID(screenshotUID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_SCREENSHOT, nil)
		return
	}
	screenshot, err := helpermodel.SetScreeshotApprove(screenshotID, params.Approve)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
		return
	}
	//set score

	app.Response(http.StatusOK, e.SUCCESS, screenshot.Public())

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

// ReorganizeMission ...
// @Tags Admin
// @Summary update mission table (active, inactive)
// @Produce application/json
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /admin/reorganize [get]
func ReorganizeMission(c *gin.Context) {
	app := handler.Gin{C: c}

	stat, err := helpermodel.ReorganizeMission()
	if err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	app.Response(http.StatusOK, e.SUCCESS, stat)
}
