package name

import (
	"errors"
	"strconv"
	"strings"
)

type Name interface {
	// 服务注册
	Reg() error
	// 服务列表
	ServiceList() (interface{}, error)
	// 获取某服务
	ServicesWithFilter(filter string) (interface{}, error)
}

// 返回host , port
func parseAddr(addr string) (string, int, error) {
	if len(strings.Split(addr, ":")) < 2 {
		return "", 0, errors.New("addr 无效")
	}

	split := strings.Split(addr, ":")
	portStr, err := strconv.Atoi(split[1])
	return split[0], portStr, err
}
