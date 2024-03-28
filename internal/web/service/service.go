package service

import (
	"github.com/ljinf/meet_server/internal/rpc"
	"github.com/ljinf/meet_server/pkg/log"
)

type Service struct {
	logger    *log.Logger
	rpcClient *rpc.Client
}

func NewService(l *log.Logger, rpcClient *rpc.Client) *Service {
	return &Service{
		logger:    l,
		rpcClient: rpcClient,
	}
}
