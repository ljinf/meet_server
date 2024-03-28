package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ljinf/meet_server/cmd/service/account/wire"
	"github.com/ljinf/meet_server/pkg/config"
	"github.com/ljinf/meet_server/pkg/helper/port"
	"github.com/ljinf/meet_server/pkg/log"
	"github.com/ljinf/meet_server/pkg/name"
	"go.uber.org/zap"
	"net"
)

func main() {

	conf := config.NewConfig("config/config-dev.yaml")
	logger := log.NewLog(conf)

	addr := fmt.Sprintf("%v:%v", conf.AccountSrvConfig.Host, port.GenRandomPort())

	logger.Info("handler start", zap.String("host", addr))

	server, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("account srv 启动失败", err.Error())
		panic(err)
	}

	consulRegister := name.NewConsulName(
		name.WithConsulAddr(fmt.Sprintf("%v:%v", conf.ConsulConfig.Host, conf.ConsulConfig.Port)),
		name.WithServiceAddr(addr),
		name.WithConsulID(uuid.New().String()),
		name.WithConsulName(conf.AccountSrvConfig.ServiceName),
		name.WithGrpcHealthAddr(addr),
		name.WithConsulTags(conf.AccountSrvConfig.Tags),
	)
	if err = consulRegister.Reg(); err != nil {
		logger.Error(fmt.Sprintf("consul 注册失败 %+v", err))
	}

	if err = server.Serve(listen); err != nil {
		panic(err)
	}

}
