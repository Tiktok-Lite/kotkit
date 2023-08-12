package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video/videoservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/kitex/server"
	"net"
)

var (
	userConfig  = conf.LoadConfig(constant.DefaultVideoConfigName)
	serviceName = userConfig.GetString("server.name")
	serviceAddr = fmt.Sprintf("%s:%d", userConfig.GetString("server.host"), userConfig.GetInt("server.port"))
)

func main() {
	logger := log.Logger()

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Errorf("Error occurs when resolving video service address: %v", err)
		panic(err)
	}
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running video service server: %v", err)
		panic(err)
	}
	logger.Infof("Video service server start successfully at %s", serviceAddr)
}
