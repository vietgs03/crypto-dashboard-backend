package config

import (
	"crypto-dashboard-backend/pkg/global"
	"crypto-dashboard-backend/pkg/utils"
	"fmt"
	"github.com/spf13/viper"
)

var config global.AppServer

func MustLoadConfig(configPath string) {
	var dirPath = utils.GetDirectoryPath(configPath)
	fileName, _ := utils.GetFileName(configPath)

	viper := viper.New()
	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&config.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
}
