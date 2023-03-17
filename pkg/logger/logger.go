package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {

	// 获取日志存储格式
	enc := getEncoder(true)
	// 获取日志写入介质
	ws := getWriter(filename, logType, maxSize, maxBackup, maxAge, compress, true)
	// 获取日志等级
	enab := getLevel(level)

	// 初始化 core
	core := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(ws...), enab)

	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// getEncoder 日志存储格式
func getEncoder(debug bool) zapcore.Encoder {

	// 日志格式规则
	ec := zapcore.EncoderConfig{
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
	}

	if debug {
		// 终端输出的关键词高亮
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 控制台输出（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(ec)
	}

	return zapcore.NewJSONEncoder(ec)
}

// 日志记录介质。Bingo 中使用了两种介质，os.Stdout 和文件
func getWriter(filename, logType string, maxSize, maxBackup, maxAge int, compress, debug bool) []zapcore.WriteSyncer {
	if logType == "daily" {
		name := time.Now().Format("2006-01-02") + ".log"
		filename = strings.ReplaceAll(filename, "logs.log", name)
	}

	// 滚动日志
	logger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}

	// 记录到文件
	writeSyncer := []zapcore.WriteSyncer{zapcore.AddSync(logger)}
	// 调试打印到终端
	if debug {
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	return writeSyncer
}

// getLevel 获取日志级别
func getLevel(level string) zap.AtomicLevel {
	logLevel := zap.NewAtomicLevel()
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误")
	}

	return logLevel
}
