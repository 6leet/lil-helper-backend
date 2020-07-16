package jwt

import (
	"errors"
	"lil-helper-backend/hashids"
	apimodel "lil-helper-backend/model/apiModel"
	helpermodel "lil-helper-backend/model/helperModel"
	"lil-helper-backend/pkg/handler"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id" // to be set

var Jwt jwt.GinJWTMiddleware

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "lil-helper",                    // to be set
		Key:         []byte("lil-helper secret key"), // to be set
		Timeout:     time.Hour,                       // to be set
		MaxRefresh:  time.Hour,                       // to be set
		IdentityKey: identityKey,

		// done
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*helpermodel.User); ok {
				return jwt.MapClaims{
					identityKey: v.UID,
					"username":  v.Username,
					"admin":     v.Admin,
					"nonce":     v.UpdatedAt.Unix(),
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
			} else if !user.Active || user.UpdatedAt.Unix() > int64(nonce) {
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
				return user, nil
			} else if errors.Unwrap(err) != nil {
				return nil, jwt.ErrFailedTokenCreation
			}

			return nil, jwt.ErrFailedAuthentication
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	} else {
		Jwt = *authMiddleware
	}
}
