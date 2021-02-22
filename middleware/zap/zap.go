package zap

import (
	"os"

	// grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

// ZapInterceptor lumberjack to log files
func ZapInterceptor() *zap.Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeDuration = zapcore.StringDurationEncoder
	config.EncodeCaller = zapcore.FullCallerEncoder

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	// errorLevel  输出文件
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), zapcore.AddSync(os.Stdout), infoLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(), errorLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	// grpc_zap.ReplaceGrpcLogger(logger)
	Log = logger.Sugar()
	return logger
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "api.log", // 日志文件位置
		MaxSize:    10,        // 日志文件最大大小(MB)
		MaxBackups: 5,         // 保留旧文件最大数量
		MaxAge:     30,        // 保留旧文件最长天数
		Compress:   true,      // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 编码器
func getEncoder() zapcore.Encoder {
	// 使用默认的JSON编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
