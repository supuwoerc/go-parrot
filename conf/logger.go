package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 初始化日志组件
func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	fmt.Println(viper.GetBool("mode.dev"))
	if !viper.GetBool("mode.dev") {
		logMode = zapcore.InfoLevel
	}
	//输出目标(文件&控制台)
	writeSyncers := zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout))
	//配置encoder,输出目标,模式
	core := zapcore.NewCore(getEncoder(), writeSyncers, logMode)
	return zap.New(core).Sugar()
}

// 配置日志输出内容encoder
func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time"                          //时间key
	config.EncodeLevel = zapcore.CapitalLevelEncoder //level全部大写
	config.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(config)
}

func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	dir, _ := os.Getwd()
	logFile := strings.Join([]string{dir, "log", time.Now().Format("2006-01-02") + ".txt"}, separator)
	lumberJackLogger := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false,
	}
	return zapcore.AddSync(&lumberJackLogger)
}
