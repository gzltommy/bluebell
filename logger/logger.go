package logger

import (
	"bluebell/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	Log *zap.Logger
)

func Init(cfg *setting.LogConfig, mode string) error {
	l := new(zapcore.Level)
	err := l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}
	encoder := getEncoder()
	writeSyncer := getLogWriter(cfg)

	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee( // 多个输出
			zapcore.NewCore(encoder, writeSyncer, l),                                     // 往日志文件里面写
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 终端输出
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	Log = zap.New(core, zap.AddCaller())

	// 替换 zap 库中全局的 logger
	zap.ReplaceGlobals(Log)

	zap.L().Info("init logger success")
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(cfg *setting.LogConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		//Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
