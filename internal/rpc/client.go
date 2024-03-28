package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

type Client struct {
	DiscoverAddr string //获取服务信息的地址
	ServiceName  string
}

type ClientOption func(s *Client)

func NewGrpcClient(opts ...ClientOption) *Client {
	s := &Client{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithDiscoverAddr(addr string) ClientOption {
	return func(s *Client) {
		s.DiscoverAddr = addr
	}
}

func WithServiceName(serviceName string) ClientOption {
	return func(s *Client) {
		s.ServiceName = serviceName
	}
}

func (c *Client) Dial() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		c.DiscoverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	return conn, err
}

func (c *Client) DialWithConsul() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%v/%v?wait=14s", c.DiscoverAddr, c.ServiceName), //consul
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	return conn, err
}
