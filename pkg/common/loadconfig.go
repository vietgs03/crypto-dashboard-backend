package common

import (
	"fmt"

	"crypto-dashboard/pkg/response"

	"github.com/spf13/viper"
)

func LoadConfig[T any]() (*T, *response.AppError) {
	var config *T
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("env")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		return nil, response.ServerError(err.Error())
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// configure structur
	if err := viper.Unmarshal(config); err != nil {
		return nil, response.ServerError(err.Error())
	}
	return config, nil
}
