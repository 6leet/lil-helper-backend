package jwt

import (
	"errors"
	"fmt"
	"lil-helper-backend/config"
	"lil-helper-backend/hashids"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/handler"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var AdminJwt jwt.GinJWTMiddleware

func init() {
	config := config.UserJwt
	var identityKey = config.IdentityKey
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       config.Realm,          // to be set
		Key:         []byte(config.Secret), // to be set
		Timeout:     config.Timeout,        // to be set
		MaxRefresh:  config.MaxRefresh,     // to be set
		IdentityKey: identityKey,

		// done
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("check")
			fmt.Println(identityKey)
			if u, ok := data.(*helpermodel.User); ok {
				return jwt.MapClaims{
					identityKey: u.UID,
					"username":  u.Username,
					"admin":     u.Admin,
					"nonce":     u.UpdatedAt.Unix(),
					"nickname":  u.Nickname,
				}
			}
			return jwt.MapClaims{}
		},

		// done
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if uid, ok := claims[identityKey].(string); !ok {
				return nil
			} else if userID, err := hashids.DecodeUserUID(uid); err != nil {
				return nil
			} else if user, err := helpermodel.GetUser(userID); err != nil {
				return nil
			} else if nonce, ok := claims["nonce"].(float64); !ok {
				return nil
			} else if !user.Admin {
				return nil
			} else if !user.Active || int64(user.UpdatedAt.Unix()) > int64(nonce) {
				return nil
			} else {
				return user
			}
		},

		// done
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals apimodel.LoginParam
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			// to-do: login function (database)
			if user, err := helpermodel.Login(username, password); err == nil {
				// app.SetCookie("userUID", user.UID, time.Now().Add(config.Timeout), config)
				// app.SetCookie("nickname", "__admin__", time.Now().Add(config.Timeout), config)
				return user, nil
			} else if errors.Unwrap(err) != nil {
				return nil, jwt.ErrFailedTokenCreation
			}

			return nil, jwt.ErrFailedAuthentication
		},

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			app := handler.Gin{C: c}
			data := apimodel.LoginResData{Token: token, Expire: expire.Format(time.RFC3339)}
			app.SetJwtCookie(token, expire, config)

			claims := jwt.ExtractClaims(c)
			app.SetCookie("userUID", claims[identityKey].(string), expire, config)
			app.SetCookie("nickname", "__admin__", expire, config)
			app.Response(http.StatusOK, e.SUCCESS, data)
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			app := handler.Gin{C: c}
			data := apimodel.LoginResData{Token: token, Expire: expire.Format(time.RFC3339)}
			app.SetJwtCookie(token, expire, config)
			app.Response(http.StatusOK, e.SUCCESS, data)
		},

		// done
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*helpermodel.User); ok {
				return true
			}

			return false
		},

		// done
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println("jwt.go")
			app := handler.Gin{C: c}
			app.MsgResponse(code, code, message, nil)
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: x-token",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		// TimeFunc: time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	} else {
		AdminJwt = *authMiddleware
	}
}
