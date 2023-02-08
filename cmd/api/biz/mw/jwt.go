package mw

import (
	"context"
	"time"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var secretKey = "Byte Dance"
var identityKey = "id"

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(secretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		IdentityKey:   identityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &shared.User{
				ID: int64(claims[identityKey].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		// to be continue
		//
		// Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
		// 	var err error
		// 	var req basic.UserLoginRequest
		// 	if err = c.BindAndValidate(&req); err != nil {
		// 		return "", jwt.ErrMissingLoginValues
		// 	}
		// 	if len(req.Username) == 0 || len(req.Password) == 0 {
		// 		return "", jwt.ErrMissingLoginValues
		// 	}
		// 	return rpc.CheckUser(context.Background(), &demouser.CheckUserRequest{
		// 		Username: req.Username,
		// 		Password: req.Password,
		// 	})
		// },
		// LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
		// 	c.JSON(http.StatusOK, utils.H{
		// 		"code":   errno.Success.ErrCode,
		// 		"token":  token,
		// 		"expire": expire.Format(time.RFC3339),
		// 	})
		// },
		// Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
		// 	c.JSON(http.StatusOK, utils.H{
		// 		"code":    errno.AuthorizationFailedErr.ErrCode,
		// 		"message": message,
		// 	})
		// },
		// HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
		// 	switch t := e.(type) {
		// 	case errno.ErrNo:
		// 		return t.ErrMsg
		// 	default:
		// 		return t.Error()
		// 	}
		// },
	})
}
