package util

import (
	"github.com/spf13/viper"
)
type ConfigPwd struct{
	DriverName string `mapstructure:"DB_DRIVER"`
	DriverURL string `mapstructure:"DB_URL"`

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