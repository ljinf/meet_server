package server

import (
	"github.com/ljinf/meet_server/internal/service/handler"
	pb "github.com/ljinf/meet_server/pkg/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewAccountRpcServer(handler *handler.AccountServerHandler) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, handler)
	//consul健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
