package config

import (
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var APPConfig *viper.Viper

func Setup() {
	if APPConfig != nil {
		return
	}

	appConfig := viper.New()

	// 配置文件名称
	appConfig.SetConfigName("app")

	// 配置文件路径
	configPath := getConfigPath()
	appConfig.AddConfigPath(configPath)

	// 配置文件类型
	appConfig.SetConfigType("toml")

	if err := appConfig.ReadInConfig(); err != nil {
		panic(err)
	}

	APPConfig = appConfig
}

func getConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	currentPath := path.Dir(filename)

	return currentPath + "/../../conf"
}
