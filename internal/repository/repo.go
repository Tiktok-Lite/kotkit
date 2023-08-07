package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	v "github.com/Tiktok-Lite/kotkit/pkg/viper"
	z "github.com/Tiktok-Lite/kotkit/pkg/zap"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Repository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func init() {
	dbConf := v.LoadConfig(constant.DefaultDBConfigPath)
	DB = NewDB(dbConf)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:     db,
		logger: z.InitLogger(),
	}
}

func NewDB(config *viper.Viper) *gorm.DB {
	connInfo := config.GetString("data.mysql.user")
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
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
	if err = db.AutoMigrate(&model.User{}, &model.Login{}, &model.Comment{}, &model.Message{}, &model.Video{}); err != nil {
		panic(err)
	}

	return db
}
