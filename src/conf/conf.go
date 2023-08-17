package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitSystemConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./src/conf/")
	err := viper.ReadInConfig()
	if err != nil {
		// TODO:记录日志
		panic(fmt.Sprintf("配置文件读取错误,请检查配置：%s", err.Error()))
	}
	fmt.Println(viper.GetString("server.port"))
	fmt.Println("读取")
}
