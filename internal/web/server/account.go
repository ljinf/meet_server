package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ljinf/meet_server/api/v1"
	"github.com/ljinf/meet_server/internal/middleware"
	"github.com/ljinf/meet_server/internal/web/handler"
	"github.com/ljinf/meet_server/pkg/jwt"
	"github.com/ljinf/meet_server/pkg/log"
)

func NewAccountHTTPServer(
	jwt *jwt.JWT,
	logger *log.Logger,
	accountHandler handler.AccountHandler,
) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		v1.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Meet Server! This Is Account HTTP Server",
		})
	})
	r.GET("/health", func(ctx *gin.Context) {
		v1.HandleSuccess(ctx, "ok")
	})

	r.POST("register", accountHandler.CreateAccount)
	r.POST("login", accountHandler.Login)

	accountGroup := r.Group("/account", middleware.StrictAuth(jwt, logger))
	{
		accountGroup.POST("/update", accountHandler.UpdateAccountInfo)
		accountGroup.GET("/userInfo", accountHandler.GetUserInfo)
		accountGroup.POST("/userInfo/update", accountHandler.UpdateUserInfo)
	}

	return r
}
