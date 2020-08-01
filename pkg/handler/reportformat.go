package handler

import (
	"fmt"
	"lil-helper-backend/config"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"ok"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) MsgResponse(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) GetUser() (user *helpermodel.User) {
	if userInterface, ok := g.C.Get(config.UserJwt.IdentityKey); ok {
		if user, ok = userInterface.(*helpermodel.User); ok {
			return user
		}
	}
	fmt.Println("reportformat.getUser")
	g.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
	return nil
}

func (g *Gin) SetJwtCookie(token string, expire time.Time, config config.JwtConfig) {
	maxage := int(expire.Unix() - time.Now().Unix())
	if int(config.MaxRefresh.Seconds()) > maxage {
		maxage = int(config.MaxRefresh.Seconds())
	}
	g.C.SetCookie(
		config.CookieName,
		token,
		maxage,
		config.CookiePath,
		config.CookieDomain,
		config.SecureCookie,
		config.CookieHTTPOnly,
	)
}

func (g *Gin) SetCookie(cookieName string, cookie string, expire time.Time, config config.JwtConfig) {
	maxage := int(expire.Unix() - time.Now().Unix())
	if int(config.MaxRefresh.Seconds()) > maxage {
		maxage = int(config.MaxRefresh.Seconds())
	}
	g.C.SetCookie(
		cookieName,
		cookie,
		maxage,
		config.CookiePath,
		config.CookieDomain,
		config.SecureCookie,
		config.CookieHTTPOnly,
	)
}

func (g *Gin) ClearCookie(cookieName string, config config.JwtConfig) {
	g.C.SetCookie(
		cookieName,
		"",
		-1,
		config.CookiePath,
		config.CookieDomain,
		config.SecureCookie,
		config.CookieHTTPOnly,
	)
}

func (g *Gin) Redirect(path string) {
	g.C.Redirect(http.StatusMovedPermanently, path)
	return
}
