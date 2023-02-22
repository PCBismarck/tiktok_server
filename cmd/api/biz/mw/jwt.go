package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/basic"
	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/shared"
	"github.com/PCBismarck/tiktok_server/cmd/api/biz/rpc"
	"github.com/PCBismarck/tiktok_server/cmd/user/db_config/consts"
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
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req basic.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			resp, err := rpc.VerifyUser(context.Background(), basic.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			c.Set("userId", resp.UserId)
			return resp.UserId, err
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			uid, _ := c.Get("userId")
			c.JSON(http.StatusOK, basic.UserLoginResponse{
				StatusCode: consts.SuccessCode,
				StatusMsg:  "Login success",
				UserId:     uid.(int64),
				Token:      token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, basic.UserLoginResponse{
				StatusCode: consts.FailureCode,
				StatusMsg:  "Unauthorized",
			})
		},
	})
}
