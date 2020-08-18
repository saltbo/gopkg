package gormutil

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Config struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

var defaultDB *gorm.DB

func Init(conf Config, models ...interface{}) {
	db, err := gorm.Open(conf.Driver, conf.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(models...)
	defaultDB = db
}

func Debug() {
	defaultDB = defaultDB.Debug()
}

func Close() {
	defaultDB.Close()
}

func DB() *gorm.DB {
	return defaultDB
}
