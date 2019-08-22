package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// JWTConfig JWTConfig
type JWTConfig struct {
	Expire int
	Sign   string
	Realm  string
}

// GConfig GConfig
var GConfig JWTConfig

const cmdRoot = "core"

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	GConfig.Sign = viper.GetString("jwt.sign")
	GConfig.Realm = viper.GetString("jwt.realm")
	GConfig.Expire = viper.GetInt("jwt.expire")
}
