package logger

import (
	"bingo-example/pkg/config"
	"go.uber.org/zap/zapcore"
	"time"
)

type LogConfig struct {
}

func (c *LogConfig) Encoder() (zapcore.EncoderConfig, bool) {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,   // 每行日志的结尾添加 "\n"
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 日志级别名称大写，如 ERROR、INFO
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
		}, // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}, true

}

func (c *LogConfig) Writer() (filename, logType string, maxSize, maxBackup, maxAge int, compress, debug bool) {
	return config.GetString("log.filename"),
		config.GetString("log.type"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		true
}

func (c *LogConfig) Level() string {
	return config.GetString("log.level")
}
