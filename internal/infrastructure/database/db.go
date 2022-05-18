package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DataSource struct {
	DB *gorm.DB
}

func NewDataSource() *DataSource {
	return &DataSource{getDb()}
}

func getDb() *gorm.DB {
	// https://github.com/go-sql-driver/mysql
	dsn := "root:@tcp(127.0.0.1:3306)/seezoon?charset=utf8mb4&parseTime=True&loc=Local&timeout=1s"
	if gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 驼峰表名，不加s
			SingularTable: true,
		},
	}); nil != err {
		panic(err)
	} else {
		if db, err := gdb.DB(); nil == err {
			db.SetMaxIdleConns(1)
			db.SetConnMaxLifetime(60 * 60 * time.Second)
			db.SetConnMaxIdleTime(30 * 60 * time.Second)
			db.SetMaxOpenConns(10)
		} else {
			panic(err)
		}
		return gdb
	}
}
