package database

import (
	"go-admin/configs"
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
	dsn := configs.C.Db.Dsn
	if gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 驼峰表名，不加s
			SingularTable: true,
		},
	}); nil != err {
		panic(err)
	} else {
		if db, err := gdb.DB(); nil == err {
			db.SetMaxIdleConns(configs.C.Db.MaxIdleConns)
			db.SetConnMaxLifetime(configs.C.Db.ConnMaxLifetime * time.Second)
			db.SetConnMaxIdleTime(configs.C.Db.ConnMaxIdleTime * time.Second)
			db.SetMaxOpenConns(configs.C.Db.MaxOpenConns)
		} else {
			panic(err)
		}
		return gdb
	}
}
