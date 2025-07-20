package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App_Name        string
	Env             string
	Debug           bool
	Port            int
	Timeout_seconds int
	Timeout         time.Duration
}

func LoadConfig() *Config {

	// Set the file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Let's explicitly set the port value
	viper.Set("port", 9090)

	// Set the default values for fallbacks
	viper.SetDefault("app_name", "MyApp")
	viper.SetDefault("env", "staging")
	viper.SetDefault("port", 8080)
	viper.SetDefault("debug", false)
	viper.SetDefault("timeout_seconds", 25)

	// Bind environment variables
	viper.BindEnv("app_name", "APP_NAME")
	viper.BindEnv("debug", "DEBUG")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("env", "ENV")
	viper.BindEnv("timeout_seconds", "TIMEOUT_SECONDS")

	// read the config file

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("[WARNING] Could not find the config file %v \n", err)
		fmt.Println("Using default values and/or environment variables")
	} else {
		fmt.Println("✅ Config file loaded from :", viper.ConfigFileUsed())
	}

	cfg := &Config{
		App_Name:        viper.GetString("app_name"),
		Env:             viper.GetString("env"),
		Port:            viper.GetInt("port"),
		Debug:           viper.GetBool("debug"),
		Timeout_seconds: viper.GetInt("timeout_seconds"),
	}
	cfg.Timeout = time.Duration(cfg.Timeout_seconds) * time.Second

	return cfg

}
