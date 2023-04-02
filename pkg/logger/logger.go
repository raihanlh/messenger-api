package logger

import (
	"context"
	"log"

	"gitlab.com/raihanlh/messenger-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	gormlogger "gorm.io/gorm/logger"
)

var globalLogger *zap.Logger

func Setup(config *config.Config) *zap.Logger {
	if config.Debug {
		l, err := zap.NewDevelopment()
		if err != nil {
			log.Println(err.Error())
		}
		return l
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, getLogFile(), zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	globalLogger = logger.With(zap.String("app", config.AppName)).With(zap.String("environment", config.Env))
	return globalLogger
}

func GetLogger(ctx context.Context) *zap.Logger {
	if globalLogger == nil {
		l, _ := zap.NewDevelopment()
		return l
	}
	return globalLogger
}

func GetLoggerGorm() gormlogger.Interface {
	if globalLogger == nil {
		return gormlogger.Default.LogMode(gormlogger.Info)
	}
	l := GormLogger{}
	l.ZapLogger = globalLogger
	l.LogLevel = gormlogger.Info
	return l
}

func getLogFile() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/kumbangsoal/kumbangsoal-log.json",
		MaxSize:    100, // megabytes
		MaxAge:     3,   // days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	})
}
