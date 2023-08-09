package log

import (
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logConfig = conf.LoadConfig(constant.DefaultLogConfigName)
	infoPath  = logConfig.GetString("path.info")  // INFO/DEBUG/WARN 级别日志的路径
	errorPath = logConfig.GetString("path.error") // ERROR/FATAL 级别日志的路径
)

func InitLogger() *zap.SugaredLogger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zap.ErrorLevel && lvl >= zap.DebugLevel
	})

	encoder := getEncoder()

	infoSyncer := getLogWriter(infoPath)
	infoCore := zapcore.NewCore(encoder, infoSyncer, lowPriority)

	errorSyncer := getErrorLogWriter(errorPath)
	errorCore := zapcore.NewCore(encoder, errorSyncer, highPriority)

	core := zapcore.NewTee(infoCore, errorCore)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getErrorLogWriter(errorPath string) zapcore.WriteSyncer {
	return getLogWriter(errorPath)
}
