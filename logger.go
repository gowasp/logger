package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool

	ZapLevel zapcore.Level
}

var (
	timeFormat = "2006-01-02 15:04:05.000"
)

func console() zapcore.Core {
	consoleWrite := zapcore.AddSync(io.MultiWriter(os.Stdout))
	consoleConfig := zap.NewProductionEncoderConfig()
	consoleConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	// color.
	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleConfig),
		consoleWrite,
		zapcore.DebugLevel,
	)
}

func file(opt *Option) zapcore.Core {
	hook := &lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize,
		MaxBackups: opt.MaxBackups,
		MaxAge:     opt.MaxAge,
		Compress:   opt.Compress,
	}

	fileWrite := zapcore.AddSync(io.MultiWriter(hook))
	fileConfig := zap.NewProductionEncoderConfig()
	fileConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(fileConfig),
		fileWrite,
		opt.ZapLevel,
	)
	return fileCore
}

func SimpleGlobalConsole() {
	zap.ReplaceGlobals(zap.New(console(), zap.AddCaller()))
}

func SimpleGlobalFile(opt *Option) {
	core := zapcore.NewTee(console(), file(opt))
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
}
