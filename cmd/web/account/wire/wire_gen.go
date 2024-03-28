// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ljinf/meet_server/internal/rpc"
	"github.com/ljinf/meet_server/internal/web/cache"
	"github.com/ljinf/meet_server/internal/web/handler"
	"github.com/ljinf/meet_server/internal/web/server"
	"github.com/ljinf/meet_server/internal/web/service"
	"github.com/ljinf/meet_server/pkg/config"
	"github.com/ljinf/meet_server/pkg/jwt"
	"github.com/ljinf/meet_server/pkg/log"
)

// Injectors from wire.go:

func NewApp(conf *config.AppConfig, logger *log.Logger, option ...rpc.ClientOption) (*gin.Engine, func(), error) {
	jwtJWT := jwt.NewJwt(conf)
	handlerHandler := handler.NewHandler(logger, jwtJWT)
	client := rpc.NewGrpcClient(option...)
	serviceService := service.NewService(logger, client)
	redisClient := cache.NewRedis(conf)
	cacheCache := cache.NewCache(redisClient)
	accountCache := cache.NewAccountCache(cacheCache)
	accountService := service.NewAccountService(serviceService, accountCache)
	accountHandler := handler.NewAccountHandler(handlerHandler, accountService)
	engine := server.NewAccountHTTPServer(jwtJWT, logger, accountHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewAccountHTTPServer)

var JwtSet = wire.NewSet(jwt.NewJwt)

var RedisSet = wire.NewSet(cache.NewRedis)

var CacheSet = wire.NewSet(cache.NewCache, cache.NewAccountCache)

var ServiceSet = wire.NewSet(service.NewService, service.NewAccountService)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewAccountHandler)

var RpcClientSet = wire.NewSet(rpc.NewGrpcClient)