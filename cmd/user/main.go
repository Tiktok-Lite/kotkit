package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user/userservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	userConfig  = conf.LoadConfig(constant.DefaultUserConfigName)
	serviceName = userConfig.GetString("server.name")
	serviceAddr = fmt.Sprintf("%s:%d", userConfig.GetString("server.host"), userConfig.GetInt("server.port"))
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		// TODO(century): 记录到日志中
	}
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
