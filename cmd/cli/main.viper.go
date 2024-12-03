package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
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
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
	fmt.Println("config port:::", config.Server.Port)
	// Print the database connection details
	for _, db := range config.Databases {
		// Use fmt.Printf to correctly format the string with variables
		fmt.Printf("Connections: %s:%s@%s/%s\n", db.User, db.Password, db.Host, db.DbName)
	}
}
