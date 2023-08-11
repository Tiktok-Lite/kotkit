package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	once   sync.Once
	dbConf = conf.LoadConfig(constant.DefaultDBConfigName)
	db     *gorm.DB
)

func newDB(config *viper.Viper) *gorm.DB {
	connInfo := config.GetString("data.mysql.user")
	var err error
	_db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connInfo,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 自动创建和修改表结构
	if err = _db.AutoMigrate(&model.User{}, &model.Login{}, &model.Comment{}, &model.Message{}, &model.Video{}); err != nil {
		panic(err)
	}

	return _db
}

// DB 数据库单例模式，每次需要数据库处理调用这个就完事
func DB() *gorm.DB {
	once.Do(func() {
		db = newDB(dbConf)
	})

	return db
}
