package main

import (
	login "github.com/Tiktok-Lite/kotkit/kitex_gen/login/userservice"
	"log"
)

func main() {
	svr := login.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
