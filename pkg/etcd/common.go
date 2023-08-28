package etcd

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"sync"
)

var (
	once       sync.Once
	etcdConfig = conf.LoadConfig(constant.DefaultEtcdConfigName)
	etcdAddr   = fmt.Sprintf("%s:%d", etcdConfig.GetString("etcd.ip"), etcdConfig.GetInt("etcd.port"))
	logger     = log.Logger()
)
