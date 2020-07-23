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
	"github.com/jmcvetta/randutil"
)

// GetMission ...
// @Tags Helper
// @Summary get mission
// @Produce application/json
// @Success 200 {object} handler.Response{data=helpermodel.Mission}
// @Router /helper/mission [get]
func GetMission(c *gin.Context) { // need: userID
	app := handler.Gin{C: c}

	var user *helpermodel.User
	if user = app.GetUser(); user == nil {
		return
	}
	userID, err := hashids.DecodeUserUID(user.UID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_INVALID_USER_UID, nil)
		return
	}
	if assignment, err := helpermodel.GetAssignment(userID); assignment != nil {
		mission, err := helpermodel.GetMission(assignment.MissionID)
		if err != nil {
			app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_MISSION, nil)
			return
		}
		app.Response(http.StatusOK, e.SUCCESS, mission.Public())
		return
	} else if err == e.ErrAssignmentNotExist {
		choices, err := helpermodel.GetMissionsWeight(user.Level)
		if err != nil {
			app.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}
		choice, err := randutil.WeightedChoice(choices)
		if err != nil {
			app.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}
		mission, _ := choice.Item.(helpermodel.Mission)
		missionID, err := hashids.DecodeMissionUID(mission.UID)
		if err := helpermodel.CreateAssignment(userID, missionID); err != nil {
			app.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}
		app.Response(http.StatusOK, e.SUCCESS, mission.Public())
	}

	// mission := helpermodel.Mission{}
}

// GetScreenshots ...
// @Tags Helper
// @Summary list screenshots by date(optional), else list today's screenshots
// @Produce application/json
// @Param datefrom query string false "screenshot date from"
// @Param dateto query string false "screenshot date to"
// @Success 200 {object} handler.Response{data=apiModel.JsonObjectArray}
// @Router /helper/screenshots [get]
func GetScreenshots(c *gin.Context) {
	app := handler.Gin{C: c}

	var user *helpermodel.User
	if user = app.GetUser(); user == nil {
		return
	}
	userID, err := hashids.DecodeUserUID(user.UID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_INVALID_USER_UID, nil)
		return
	}

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

	screenshots, err := helpermodel.GetScreenshotsByDate(userID, dateFrom, dateTo, false, false)
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
	missionID, err := hashids.DecodeMissionUID(params.MissionUID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_MISSION, nil)
		return
	}
	screenshot, err := helpermodel.CreateScreenshot(userID, missionID, params.Picture)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, screenshot.Public())
	}
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
	screenshotUID := c.Param("uid")
	screenshotID, err := hashids.DecodeScreenshotUID(screenshotUID)
	if err != nil {
		app.Response(http.StatusBadRequest, e.ERR_NO_SUCH_SCREENSHOT, nil)
		return
	}
	err = helpermodel.DeleteScreenshot(screenshotID)
	if errors.Unwrap(err) != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else if err != nil {
		app.MsgResponse(http.StatusBadRequest, e.INVALID_PARAMS, err.Error(), nil)
	} else {
		app.Response(http.StatusOK, e.SUCCESS, nil)
	}
}
