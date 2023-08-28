package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation/relationservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/jwt"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/kitex/server"
	"net"
)

var (
	logger = log.Logger()
	//这里不是userconfig 去看下最新的user下main的代码
	relationConfig = conf.LoadConfig(constant.DefaultRelationConfigName)
	serviceName    = relationConfig.GetString("server.name")
	serviceAddr    = fmt.Sprintf("%s:%d", relationConfig.GetString("server.host"), relationConfig.GetInt("server.port"))
	jwtConfig      = conf.LoadConfig(constant.DefaultLoginConfigName)
	signingKey     = jwtConfig.GetString("JWT.signingKey") // zhelibudui...你去看下最新的代码，不是从usserconfig读取签jwt的密钥
	Jwt            *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT([]byte(signingKey))
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Errorf("Error occurs when resolving login service address: %v", err)
		panic(err)
	}
	svr := relationservice.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running login service server: %v", err)
		panic(err)
	}
	logger.Infof("Relation service server start successfully at %s", serviceAddr)
}
