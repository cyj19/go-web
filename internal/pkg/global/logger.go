package global

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

/*
	自定义gorm logger，实现Interface接口
	参考gorm@v1.21.11源码logger.go
	基本思路：使用zap的logger替换标准库logger
*/

type GormZapLogger struct {
	log *zap.Logger
	logger.Config
	debugStr, infoStr, warnStr, errStr  string
	traceStr, traceErrStr, traceWarnStr string
}

func NewGormZapLogger(log *zap.Logger, config logger.Config) *GormZapLogger {
	var (
		debugStr     = "%s\n[debug]"
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &GormZapLogger{
		log:          log,
		Config:       config,
		debugStr:     debugStr,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// 实现Interface接口

func (l *GormZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l GormZapLogger) Debug(ctx context.Context, msg string, data ...interface{}) {
	if l.log.Core().Enabled(zapcore.DebugLevel) {
		fmt.Println("执行debug")
		l.log.Sugar().Debugf(l.debugStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l GormZapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	// zapcore的level等于大于info级别
	if l.log.Core().Enabled(zapcore.InfoLevel) {
		l.log.Sugar().Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l GormZapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.log.Core().Enabled(zapcore.WarnLevel) {
		l.log.Sugar().Warnf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l GormZapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.log.Core().Enabled(zap.ErrorLevel) {
		l.log.Sugar().Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 打印panic
	if !l.log.Core().Enabled(zapcore.DPanicLevel) || l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.log.Core().Enabled(zapcore.ErrorLevel) && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.log.Sugar().Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.log.Sugar().Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.log.Core().Enabled(zapcore.WarnLevel):
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.log.Sugar().Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.log.Sugar().Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.log.Core().Enabled(zapcore.InfoLevel):
		sql, rows := fc()
		if rows == -1 {
			l.log.Sugar().Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.log.Sugar().Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
