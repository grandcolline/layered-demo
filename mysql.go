package main

import (
	"fmt"

	"github.com/grandcolline/layered-demo/infrastructure/datastore/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mysqlConnect はMySQLに接続します
func mysqlConnect() (db *gorm.DB) {
	var err error

	// MySQLの設定の読み込み
	var conf mysqlConf
	conf.init()

	// DB接続
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	if db, err = gorm.Open("mysql", url); err != nil {
		panic(err)
	}

	// GORMのログ設定
	db.LogMode(conf.LogMode)

	// マイグレーション
	if conf.AutoMigrate {
		db.AutoMigrate(&model.ItemMdl{})
	}

	return
}

// mysqlClose はMySQLとの接続を切断します
func mysqlClose() {
	db.Close()
}
