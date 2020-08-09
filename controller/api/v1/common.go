package v1

import (
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTopScoreUsers ...
// @Tags Common
// @Summary list top score users
// @Produce application/json
// @Param limit query int false "top helpers limit"
// @Success 200 {object} handler.Response{data=apimodel.JsonObjectArray}
// @Router /helpers [get]
func GetTopScoreHelpers(c *gin.Context) {
	app := handler.Gin{C: c}

	if helpers, err := helpermodel.GetUsers(true, false, false, true, "%%"); err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		var publicHelpers []helpermodel.PublicUser
		if limitstr, ok := c.GetQuery("limit"); ok {
			limit, _ := strconv.Atoi(limitstr)
			for i, u := range helpers {
				if i >= limit {
					break
				}
				publicHelpers = append(publicHelpers, u.Public())
			}
		} else {
			for _, u := range helpers {
				publicHelpers = append(publicHelpers, u.Public())
			}
		}
		jsonArray := apimodel.NewJsonObjectArray(publicHelpers)
		app.Response(http.StatusOK, e.SUCCESS, jsonArray)
	}
}

// GetUser ...
// @Tags Common
// @Summary get user profile
// @Produce application/json
// @Success 200 {object} handler.Response{data=helpermodel.User}
// @Router /profile [get]
func GetUser(c *gin.Context) {
	app := handler.Gin{C: c}

	var user *helpermodel.User
	if user = app.GetUser(); user == nil {
		return
	}
	app.Response(http.StatusOK, e.SUCCESS, user.Public())
}
