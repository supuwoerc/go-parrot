package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go-parrot/src/global"
)

func InitSystemConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./src/conf/")
	err := viper.ReadInConfig()
	if err != nil {
		errorInfo := fmt.Sprintf("配置文件读取错误,请检查配置：%s", err.Error())
		global.Logger.Error(errorInfo)
		panic(errorInfo)
	}
}
