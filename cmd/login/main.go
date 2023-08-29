package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/loginservice"
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
	logger      = log.Logger()
	loginConfig = conf.LoadConfig(constant.DefaultLoginConfigName)
	serviceName = loginConfig.GetString("server.name")
	signingKey  = loginConfig.GetString("JWT.signingKey")
	serviceAddr = fmt.Sprintf("%s:%d", loginConfig.GetString("server.host"), loginConfig.GetInt("server.port"))
	expiryTime  = loginConfig.GetDuration("JWT.expiryTime")
	Jwt         *jwt.JWT
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
		logger.Errorf("Error occurs when resolving login service address: %v", err)
		panic(err)
	}
	svr := loginservice.NewServer(new(LoginServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running login service server: %v", err)
		panic(err)
	}
	logger.Infof("Login service server start successfully at %s", serviceAddr)
}
