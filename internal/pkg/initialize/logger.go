package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/vagaryer/go-web/internal/pkg/config"
	"github.com/vagaryer/go-web/internal/pkg/logger"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

// 初始化日志，使用zap+lumberjack代替标准库的log
func InitLogger(conf *config.Configuration) *logger.GormZapLogger {

	// 自定义编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义时间格式
	encoderConfig.EncodeTime = ZapLogLocalTimeEncoder
	// 使用大写字母+颜色记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 日志文件名
	now := time.Now()
	fileName := fmt.Sprintf("%s/%04d-%02d-%02d", conf.Log.Path, now.Year(), now.Month(), now.Day())
	// 使用lumberjack进行日志配置
	lumberjackLog := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    conf.Log.MaxSize,
		MaxAge:     conf.Log.MaxAge,
		MaxBackups: conf.Log.MaxBackups,
		Compress:   conf.Log.Compress,
	}
	// 打印到控制台和日志文件
	writerSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLog), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, writerSyncer, conf.Log.Level)
	// 创建日志对象
	log := zap.New(core, zap.AddCaller())
	glog := logger.NewGormZapLogger(log, gormlogger.Config{
		Colorful: true,
	})
	return glog
}

// zap日志自定义时间格式
func ZapLogLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.MsecLocalTimeFormat))
}
