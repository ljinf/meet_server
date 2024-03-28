package name

import (
	"errors"
	"github.com/hashicorp/consul/api"
)

type ConsulName struct {
	ConsulAddr     string
	ServiceAddr    string
	HTTPHealthUrl  string
	GrpcHealthAddr string
	Name           string
	Id             string
	Tags           []string
}

type ConsulOption func(n *ConsulName)

func NewConsulName(opts ...ConsulOption) Name {
	cn := &ConsulName{}

	for _, opt := range opts {
		opt(cn)
	}

	return cn
}

func WithConsulAddr(addr string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.ConsulAddr = addr
	}
}

func WithServiceAddr(addr string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.ServiceAddr = addr
	}
}

func WithHTTPHealthUrl(url string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.HTTPHealthUrl = url
	}
}

func WithGrpcHealthAddr(addr string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.GrpcHealthAddr = addr
	}
}

func WithConsulName(name string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.Name = name
	}
}

func WithConsulID(id string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.Id = id
	}
}

func WithConsulTags(tag []string) func(n *ConsulName) {
	return func(n *ConsulName) {
		n.Tags = tag
	}
}

func (c *ConsulName) Reg() error {
	config := api.DefaultConfig()

	config.Address = c.ConsulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	host, port, err := parseAddr(c.ServiceAddr)
	if err != nil {
		return errors.New("service addr 无效")
	}

	agent := &api.AgentServiceRegistration{}
	agent.Address = host
	agent.Port = port
	agent.ID = c.Id
	agent.Name = c.Name
	agent.Tags = c.Tags

	check := api.AgentServiceCheck{
		Timeout:                        "5s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "10s",
	}

	if c.HTTPHealthUrl != "" {
		check.HTTP = c.HTTPHealthUrl
	}
	if c.GrpcHealthAddr != "" {
		check.GRPC = c.GrpcHealthAddr
	}

	agent.Check = &check

	return client.Agent().ServiceRegister(agent)
}

func (c *ConsulName) ServiceList() (interface{}, error) {
	config := api.DefaultConfig()
	config.Address = c.ConsulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	// map[string]*AgentService
	services, err := client.Agent().Services()
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (c *ConsulName) ServicesWithFilter(filter string) (interface{}, error) {
	config := api.DefaultConfig()
	config.Address = c.ConsulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	//filter:Service==account_web
	services, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil, err
	}
	return services, nil
}
