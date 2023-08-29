package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation/relationservice"
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
	logger         = log.Logger()
	relationConfig = conf.LoadConfig(constant.DefaultRelationConfigName)
	serviceName    = relationConfig.GetString("server.name")
	serviceAddr    = fmt.Sprintf("%s:%d", relationConfig.GetString("server.host"), relationConfig.GetInt("server.port"))
	jwtConfig      = conf.LoadConfig(constant.DefaultLoginConfigName)
	signingKey     = jwtConfig.GetString("JWT.signingKey") //
	Jwt            *jwt.JWT
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
	svr := relationservice.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running login service server: %v", err)
		panic(err)
	}
	logger.Infof("Relation service server start successfully at %s", serviceAddr)
}
