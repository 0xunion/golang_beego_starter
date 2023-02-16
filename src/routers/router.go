package routers

import (
	"time"

	conf "github.com/0xunion/exercise_back/src/const/conf"
	"github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
	auth "github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/decorator"
	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
)

func init() {
	// insert router chain
	beego.InsertFilter("/api/*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", conf.CorsOrigin())
		ctx.Output.Header("Access-Control-Allow-Headers", conf.CorsHeaders())
		ctx.Output.Header("Access-Control-Allow-Methods", conf.CorsMethods())
		ctx.Output.Header("Access-Control-Allow-Credentials", conf.CorsCredentials())
		ctx.Output.Header("Access-Control-Max-Age", conf.CorsAge())

		// if request is OPTIONS, just return 200 to ensure CORS
		if ctx.Request.Method == "OPTIONS" {
			ctx.Output.Status = 200
			ctx.Output.Body([]byte{})
			return
		}

		// check header for Authorization
		// if found, just get user info from cache
		token := ctx.Input.Header("Authorization")
		if token != "" {
			login_token := auth.NewAuthTokenWithToken(token)
			// check if token is valid
			if login_token.AnalysisToken() {
				if login_token.Session.Login_time+3600*24*7 < uint32(time.Now().Unix()) {
					ctx.Output.Status = 401
					ctx.Output.Body([]byte("login token expired"))
					return
				}
				uid := login_token.GetUid()
				user, err := decorator.WithMemCacheGet(func(key string) (types.User, error) {
					v, err := model.ModelGet[types.User](model.NewMongoFilter(model.IdFilter(uid)))
					if err != nil {
						return types.User{}, err
					}
					return *v, nil
				}, uid.Hex())
				if err == nil {
					ctx.Input.SetData("user", user)
				}
			}
		}
	})
	// check if controller from production environment
	registerCURDApi()
	registerCommon()
	registerProd()
}
