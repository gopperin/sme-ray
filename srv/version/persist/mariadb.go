package persist

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gopperin/sme-ray/srv/version/bean"
	"github.com/gopperin/sme-ray/srv/version/config"
)

// GMariadb GMariadb
var GMariadb Mariadb

// Mariadb Mariadb
type Mariadb struct {
	db *gorm.DB
}

// Init Init
func (maria *Mariadb) Init() error {
	db, err := gorm.Open(config.MariaDB.Dialect, config.MariaDB.URL)
	if err != nil {
		return err
	}

	db.LogMode(false)
	db.DB().SetMaxIdleConns(config.MariaDB.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MariaDB.MaxOpenConns)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	maria.db = db

	if !db.HasTable(&bean.Version{}) {
		db.CreateTable(&bean.Version{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Version{})
	}

	return nil
}
