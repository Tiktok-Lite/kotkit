package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/loginservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/kitex/server"
	"net"
)

var (
	logger      = log.Logger()
	userConfig  = conf.LoadConfig(constant.DefaultLoginConfigName)
	serviceName = userConfig.GetString("server.name")
	serviceAddr = fmt.Sprintf("%s:%d", userConfig.GetString("server.host"), userConfig.GetInt("server.port"))
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Errorf("Error occurs when resolving login service address: %v", err)
		panic(err)
	}
	svr := loginservice.NewServer(new(LoginServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running login service server: %v", err)
		panic(err)
	}
	logger.Infof("Login service server start successfully at %s", serviceAddr)
}

