package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message/messageservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/jwt"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
)

var (
	logger        = log.Logger()
	messageConfig = conf.LoadConfig(constant.DefaultMessageConfigName)
	serviceName   = messageConfig.GetString("server.name")
	serviceAddr   = fmt.Sprintf("%s:%d", messageConfig.GetString("server.host"), messageConfig.GetInt("server.port"))
	jwtConfig     = conf.LoadConfig(constant.DefaultLoginConfigName)
	signingKey    = jwtConfig.GetString("JWT.signingKey")
	Jwt           *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT([]byte(signingKey))
}

func main() {
	r, err := etcd.Registry()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd registry: %v", err)
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Errorf("Error occurs when resolving message service address: %v", err)
		panic(err)
	}
	svr := messageservice.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running message service server: %v", err)
		panic(err)
	}
	logger.Infof("Message service server start successfully at %s", serviceAddr)
}
