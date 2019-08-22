package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// ServerConfig ServerConfig
type ServerConfig struct {
	APIAppendKey string
	APIMd5Key    string
	JWTSign      string
	JWTRealm     string
}

// GConfig GConfig
var GConfig ServerConfig

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

	GConfig.APIAppendKey = viper.GetString("server.salt")
	GConfig.APIMd5Key = viper.GetString("server.sign")
	GConfig.JWTSign = viper.GetString("server.jwt_sign")
	GConfig.JWTRealm = viper.GetString("server.jwt_realm")
}
