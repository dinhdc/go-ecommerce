package initialize

import (
	"fmt"

	"github.com/dinhdc/go-ecommerce/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read file configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// unmarshal
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}
