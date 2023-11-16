package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Database(dsn string, tablePrefix string) {

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("connect err:", err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)  //设置连接池，空闲
	sqlDb.SetMaxOpenConns(100) //设置打开最大连接
	DB = db
	migration()
}
