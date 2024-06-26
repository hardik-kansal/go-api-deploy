package util

import (
	"time"

	"github.com/spf13/viper"
)
type ConfigPwd struct{
	DriverName string `mapstructure:"DB_DRIVER"`
	DriverURL string `mapstructure:"DB_URL"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

}
func LoadConfig(path string)(ConfigPwd,error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err:=viper.ReadInConfig()
	var config ConfigPwd
	if err!=nil{
		return ConfigPwd{},err
	}
	err=viper.Unmarshal(&config)
	return config,err
}