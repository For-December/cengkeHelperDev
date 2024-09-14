package logger

import (
	"context"
	"fmt"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct{}

// LogMode 定义日志模式
func (l *GormLogger) LogMode(_ gormLog.LogLevel) gormLog.Interface {
	return l
}

// Info 输出信息日志
func (l *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	fmt.Printf("[INFO] %s\n", fmt.Sprintf(msg, data...))
}

// Warn 输出警告日志
func (l *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	fmt.Printf("[WARN] %s\n", fmt.Sprintf(msg, data...))
}

// Error 输出错误日志
func (l *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	fmt.Printf("[ERROR] %s\n", fmt.Sprintf(msg, data...))
}

// Trace 输出跟踪日志
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	info, _ := fc()

	if err != nil {
		TraceF("[%.3fms] [error] %s", float64(time.Since(begin).Microseconds())/1000, info)
	} else {
		TraceF("[%.3fms] %s", float64(time.Since(begin).Microseconds())/1000, info)
	}
}
