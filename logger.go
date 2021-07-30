package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option struct {
	Filename   string
	TimeFormat string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool

	ZapLevel zapcore.Level
}

type Logger struct {
	opt *Option
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Option(opt *Option) {
	l.opt = opt
}

func (l *Logger) console() zapcore.Core {
	consoleWrite := zapcore.AddSync(io.MultiWriter(os.Stdout))
	consoleConfig := zap.NewProductionEncoderConfig()
	consoleConfig.EncodeTime = zapcore.TimeEncoderOfLayout(l.opt.TimeFormat)
	// color.
	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleConfig),
		consoleWrite,
		zapcore.DebugLevel,
	)
}

func InitGlobalConsole() {
	l := &Logger{
		opt: &Option{
			TimeFormat: "2006-01-02 15:04:05.000",
		},
	}

	zap.ReplaceGlobals(zap.New(l.console(), zap.AddCaller()))
}
