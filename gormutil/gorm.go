package gormutil

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var defaultDB *gorm.DB

func Init(conf Config, debug bool) {
	defaultDB = New(conf, debug)
}

func Close() {
	defaultDB.Close()
}

func SetupPrefix(prefix string) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}
}

func AutoMigrate(models []interface{}) {
	defaultDB.AutoMigrate(models...)
}

func DB() *gorm.DB {
	return defaultDB
}

type Config struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

func New(conf Config, debug bool) *gorm.DB {
	db, err := gorm.Open(conf.Driver, conf.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	db.SingularTable(true)
	if debug {
		db.Debug()
	}

	return db
}
