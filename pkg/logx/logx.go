package logx

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger zap日志简单封装 默认开发模式 打印控制台
func NewLogger(devMode bool, name string) *zap.Logger {
	// 日志轮转
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "../../tmp/" + name + ".log",
		MaxSize:    100, // megabytes
		MaxBackups: 60,
		MaxAge:     1, // days
		LocalTime:  true,
		Compress:   false,
	})
	// 编码设置
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "name"
	encoderConfig.CallerKey = "caller"

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core
	if devMode {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	} else {
		core = zapcore.NewCore(encoder, w, zap.WarnLevel)
	}
	// zapTee := zapcore.NewTee(
	// 	zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	// 	zapcore.NewCore(encoder, w, zap.DebugLevel),
	// )
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
