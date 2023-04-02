package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"moul.io/zapgorm2"
)

type GormLogger struct {
	zapgorm2.Logger
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	log := GetLogger(ctx)
	sql, rows := fc()
	maxLengthSql := 1000
	if len(sql) > maxLengthSql {
		sql = fmt.Sprintf("%s ...... %s", sql[:maxLengthSql/2], sql[len(sql)-(maxLengthSql/2):])
	}

	message := fmt.Sprintf("%s\n[%.3fms] [rows:%v]\n%s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		log.Error(message, zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		log.Warn(message, zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	case l.LogLevel >= gormlogger.Info:
		log.Debug(message, zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	}
}
