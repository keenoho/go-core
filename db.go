package core

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func LoadDb() {
	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", conf.DbUsername, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbDatabase, "Asia%2fShanghai")

	mysqlDb := mysql.Open(dsn)
	linkDb, err := gorm.Open(mysqlDb, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := linkDb.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(1)  // 空闲链接数
	sqlDB.SetMaxOpenConns(10) // 最大链接数
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = linkDb
	testDbConnect()
}

func testDbConnect() {
	var sum int
	Db.Raw("SELECT 1+1").Scan(&sum)
	if sum != 2 {
		log.Fatalln(sum)
	}
	log.Println("db test: 1+1 =", sum, "ok")
}
