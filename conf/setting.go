package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init()(err error)  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err = viper.ReadInConfig()
	if err != nil{
		fmt.Println("viper.ReadInConfig() failed")
		return
	}
	return
}