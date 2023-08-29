package main

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment/commentservice"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/jwt"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/kitex/server"
	"net"
)

var (
	logger        = log.Logger()
	commentConfig = conf.LoadConfig(constant.DefaultCommentConfigName)
	serviceName   = commentConfig.GetString("server.name")
	serviceAddr   = fmt.Sprintf("%s:%d", commentConfig.GetString("server.host"), commentConfig.GetInt("server.port"))
	jwtConfig     = conf.LoadConfig(constant.DefaultLoginConfigName)
	signingKey    = jwtConfig.GetString("JWT.signingKey")
	Jwt           *jwt.JWT
)

func Init() {
	Jwt = jwt.NewJWT([]byte(signingKey))
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Errorf("Error occurs when resolving comment service address: %v", err)
		panic(err)
	}
	svr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServiceAddr(addr),
	)
	Init()
	err = svr.Run()

	if err != nil {
		logger.Errorf("Error occurs when running comment service server: %v", err)
		panic(err)
	}
	logger.Infof("Comment service server start successfully at %s", serviceAddr)
}
