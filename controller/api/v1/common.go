package v1

import (
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTopScoreUsers ...
// @Tags Common
// @Summary list top score users
// @Produce application/json
// @Success 200 {object} handler.Response{data=apimodel.JsonObjectArray}
// @Router /helpers [get]
func GetTopScoreHelpers(c *gin.Context) {
	app := handler.Gin{C: c}

	users := []helpermodel.User{}
	user := helpermodel.User{}
	users = append(users, user)
	users = append(users, user)
	jsonArray := apimodel.NewJsonObjectArray(users)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}

// GetTopScoreUsersLimit ...
// @Tags Common
// @Summary list top score users with limit
// @Produce application/json
// @Param limit path string true "top helpers limit"
// @Success 200 {object} handler.Response{data=apimodel.JsonObjectArray}
// @Router /helpers/{limit} [get]
func GetTopScoreHelpersLimit(c *gin.Context) {
	app := handler.Gin{C: c}

	// var limit int = 10

	users := []helpermodel.User{}
	user := helpermodel.User{}
	users = append(users, user)
	users = append(users, user)
	jsonArray := apimodel.NewJsonObjectArray(users)
	app.Response(http.StatusOK, e.SUCCESS, jsonArray)
}
