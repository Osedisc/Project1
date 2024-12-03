package infras

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	DB = newMariaConn()
}

type DBConfig struct {
	user     string
	passWord string
	host     string
	port     string
	dbName   string
}

func newMariaConn() *gorm.DB {
	dbConfig := DBConfig{
		user:     viper.GetString("db.user"),
		passWord: viper.GetString("db.passWord"),
		host:     viper.GetString("db.host"),
		port:     viper.GetString("db.port"),
		dbName:   viper.GetString("db.dbName"),
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.user, dbConfig.passWord, dbConfig.host, dbConfig.port, dbConfig.dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot open mysql connection:%s", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.MaxIdelConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.MaxOpenConn"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("db.MaxLifeTime"))
	return db
}
