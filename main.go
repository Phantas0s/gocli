package main

import (
	"path/filepath"

	"github.com/Phantas0s/go-cli/cmd"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	cmd.Execute()
}

func initConfig() {
	viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "go-cli"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
