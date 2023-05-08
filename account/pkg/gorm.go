package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-ser/account/configs"
	"time"
)

var GlobalGorm *gorm.DB

func InitGorm(config configs.MySQLConfig) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("初始化gorm失败，原因：%s", err.Error()))
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("初始化sqldb失败，原因：%s", err.Error()))
	}
	sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	if err := sqlDb.Ping(); err != nil {
		panic(fmt.Sprintf("链接数据库失败，原因：%s", err.Error()))
	}

	GlobalGorm = db
}
