package main

import (
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
)

func main() {
	config := conf.NewConfig()
	repository.NewDB(config)
}
