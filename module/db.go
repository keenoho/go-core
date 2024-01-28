package module

import (
	"fmt"
	"log"
	"time"

	"gitee.com/keenoho/go-core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var storeDb *gorm.DB

type DbModule struct {
	core.Module
}

func (m *DbModule) Init(app *core.App) {
	m.initDb()
}

func (m *DbModule) initDb() {
	if storeDb != nil {
		return
	}
	env := core.ConfigGet("ENV")
	database := core.ConfigGet("DB_DATABASE")
	username := core.ConfigGet("DB_USERNAME")
	password := core.ConfigGet("DB_PASSWORD")
	host := core.ConfigGet("DB_HOST")
	port := core.ConfigGet("DB_PORT")
	dsnAppend := core.ConfigGet("DB_DSN_APPEND")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", username, password, host, port, database, dsnAppend)
	mysqlDb := mysql.Open(dsn)
	ormConfig := gorm.Config{}
	if env == "production" {
		ormConfig.Logger = logger.Default.LogMode(logger.Error)
	} else {
		ormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	linkDb, err := gorm.Open(mysqlDb, &ormConfig)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := linkDb.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(2)   // idle connect num
	sqlDB.SetMaxOpenConns(300) // max connect num
	sqlDB.SetConnMaxLifetime(time.Hour)
	storeDb = linkDb

	m.testConnect()
}

func (m *DbModule) Db() *gorm.DB {
	return storeDb
}

func (m *DbModule) testConnect() {
	var sum int
	storeDb.Raw("SELECT 1+1").Scan(&sum)
	if sum != 2 {
		log.Fatalln(sum)
	}
	log.Println("db test: 1+1 =", sum, "ok")
}
