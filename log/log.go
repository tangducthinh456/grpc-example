package log

import (
	"log"
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerDebug        *zap.SugaredLogger //= new(zap.SugaredLogger)
	loggerWithoutDebug *zap.SugaredLogger
)

func init() {
	loggerDebug = ConfigZapWithDebug()
	loggerWithoutDebug = ConfigZapWithoutDebug()
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func ConfigZapWithDebug() *zap.SugaredLogger {
	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"server.log", "stderr"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeLevel:  CustomLevelEncoder,
			EncodeTime:   SyslogTimeEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}

func ConfigZapWithoutDebug() *zap.SugaredLogger {

	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"server.log", "stderr"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			LevelKey:    "level",
			CallerKey:   "caller",
			EncodeLevel: CustomLevelEncoder,
			EncodeTime:  SyslogTimeEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}

func Debug(str string) {
	loggerDebug.Debugf("%s", str)
}

func Info(str string) {
	loggerWithoutDebug.Infof("%s", str)
}

func Warning(str string) {
	loggerWithoutDebug.Warnf("%s", str)
}

func Error(str string) {
	loggerDebug.Errorf("%s", str)
}

func Fatal(str string) {
	loggerDebug.Fatalf("%s", str)
}
