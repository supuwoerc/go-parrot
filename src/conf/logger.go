package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(config)
}

func getWriterSyncer() zapcore.WriteSyncer {
	projectDir, _ := os.Getwd()
	separator := string(filepath.Separator)
	logFileNameWithoutSuffix := strings.Join([]string{projectDir, "log", time.Now().Format(time.DateOnly)}, separator)
	logFileName := strings.Join([]string{logFileNameWithoutSuffix, "txt"}, ".")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger), zapcore.AddSync(os.Stdout))
}

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.dev") {
		logMode = zapcore.ErrorLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriterSyncer(), logMode)
	return zap.New(core).Sugar()
}
