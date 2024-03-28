package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ljinf/meet_server/cmd/web/account/wire"
	"github.com/ljinf/meet_server/internal/rpc"
	"github.com/ljinf/meet_server/pkg/config"
	"github.com/ljinf/meet_server/pkg/helper/port"
	"github.com/ljinf/meet_server/pkg/log"
	"github.com/ljinf/meet_server/pkg/name"
	"github.com/ljinf/meet_server/pkg/server/http"
	"go.uber.org/zap"
)

func main() {

	conf := config.NewConfig("config/config-dev.yaml")
	logger := log.NewLog(conf)

	serverPort := port.GenRandomPort()
	addr := fmt.Sprintf("%v:%v", conf.AccountWebConfig.Host, serverPort)
	logger.Info("handler start", zap.String("host", addr))

	app, cleanup, err := wire.NewApp(conf, logger,
		rpc.WithDiscoverAddr(fmt.Sprintf("%v:%v", conf.ConsulConfig.Host, conf.ConsulConfig.Port)),
		rpc.WithServiceName(conf.AccountWebConfig.DependOn))

	if err != nil {
		panic(err)
	}
	defer cleanup()

	consulRegister := name.NewConsulName(
		name.WithConsulAddr(fmt.Sprintf("%v:%v", conf.ConsulConfig.Host, conf.ConsulConfig.Port)),
		name.WithServiceAddr(addr),
		name.WithConsulID(uuid.New().String()),
		name.WithConsulName(conf.AccountWebConfig.ServiceName),
		name.WithHTTPHealthUrl(fmt.Sprintf("http://%v/health", addr)),
		name.WithConsulTags(conf.AccountWebConfig.Tags),
	)
	if err = consulRegister.Reg(); err != nil {
		logger.Error(fmt.Sprintf("consul 注册失败 %+v", err))
	}

	http.Run(app, fmt.Sprintf(":%d", serverPort))

}
