package persist

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"github.com/gopperin/sme-ray/srv/version/config"
)

// ZapLogger ZapLogger
var ZapLogger *zap.Logger

func init() {

	ZapLogger = InitLogger(config.Logger.File, config.Logger.Level)

	return
}

// InitLogger InitLogger
func InitLogger(logpath string, loglevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 0,       // 日志文件最多保存多少个备份
		MaxAge:     0,       // 文件最多保存多少天
		Compress:   true,    // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		// NameKey:        "logger",
		CallerKey:  "linenum",
		MessageKey: "msg",
		// StacktraceKey:  "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder, // 小写编码器
		// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeTime:     ISO8601TimeEncoder,             // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	_atomicLevel := zap.NewAtomicLevel()
	var _level zapcore.Level
	switch loglevel {
	case "debug":
		_level = zap.DebugLevel
	case "info":
		_level = zap.InfoLevel
	case "error":
		_level = zap.ErrorLevel
	default:
		_level = zap.InfoLevel
	}
	_atomicLevel.SetLevel(_level)

	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewConsoleEncoder(encoderConfig),                                        // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		_atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	// caller := zap.AddCaller()
	// 开启文件及行号
	// development := zap.Development()
	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	// logger := zap.New(core, caller, development, filed)
	// logger := zap.New(core, caller, development)
	logger := zap.New(core)

	logger.Info("DefaultLogger init success")

	return logger
}

// ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
// with millisecond precision.
func ISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("20060102 15:04:05"))
}
