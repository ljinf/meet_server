//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ljinf/meet_server/internal/service/handler"
	"github.com/ljinf/meet_server/internal/service/repository"
	"github.com/ljinf/meet_server/internal/service/server"
	"github.com/ljinf/meet_server/pkg/config"
	"github.com/ljinf/meet_server/pkg/helper/sid"
	"github.com/ljinf/meet_server/pkg/log"
	"google.golang.org/grpc"
)

var SidSet = wire.NewSet(sid.NewSid)

var DBSet = wire.NewSet(repository.NewDB)

var RepositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewAccountRepository,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAccountServerHandler,
)

var ServerSet = wire.NewSet(server.NewAccountRpcServer)

func NewApp(conf *config.AppConfig, logger *log.Logger) (*grpc.Server, func(), error) {
	panic(wire.Build(
		SidSet,
		DBSet,
		RepositorySet,
		HandlerSet,
		ServerSet),
	)
}
