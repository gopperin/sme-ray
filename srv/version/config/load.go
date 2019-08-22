package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/gopperin/sme-ray/srv/version/bean"
)

// MariaDB 数据库相关配置
var MariaDB bean.DBConfig

// Logger Logger
var Logger bean.LoggerConfig

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

	MariaDB.Dialect = viper.GetString("database.dialect")
	MariaDB.Database = viper.GetString("database.database")
	MariaDB.User = viper.GetString("database.user")
	MariaDB.Password = viper.GetString("database.password")
	MariaDB.Host = viper.GetString("database.host")
	MariaDB.Port = viper.GetInt("database.port")
	MariaDB.Charset = viper.GetString("database.charset")
	MariaDB.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	MariaDB.MaxOpenConns = viper.GetInt("database.maxOpenConns")
	MariaDB.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		MariaDB.User, MariaDB.Password, MariaDB.Host, MariaDB.Port, MariaDB.Database, MariaDB.Charset)

	Logger.File = viper.GetString("logger.file")
	Logger.Level = viper.GetString("logger.level")
}
