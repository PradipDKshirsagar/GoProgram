package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct { 
	Chance int
}

func LoadConfiguration() (Config, error) {
	var conf = Config{}
	viper.SetConfigName("configuration") 
 	viper.AddConfigPath("./")   
	viper.AddConfigPath("./../")
	viper.SetConfigType("yaml")  
	err := viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf.Chance = viper.GetInt("Chance")
	return  conf, err
}