package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// serverConf サーバの設定
type serverConf struct {
	Port string `default:"8080"` // サーバ起動時に受け付けるポート
}

// init サーバ設定を環境変数から取得します
func (conf *serverConf) init() {
	err := envconfig.Process("server", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// mysqlConf MySQL設定
type mysqlConf struct {
	Host        string `required:"true"` // 接続先ホスト
	Port        string `default:"3306"`  // 接続先ポート
	User        string `required:"true"` // DB接続ユーザ
	Password    string `required:"true"` // DB接続パスワード
	Database    string `required:"true"` // データベース名
	LogMode     bool   `default:"false"` // SQLログを出力するか
	AutoMigrate bool   `default:"false"` // マイグレーションを自動でするかどうか
}

// init MySQL設定を環境変数から取得します
func (conf *mysqlConf) init() {
	err := envconfig.Process("mysql", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// mongoConf mongoDB設定
type mongoConf struct {
	Host     string `required:"true"` // 接続先ホスト
	Port     string `default:"27017"` // 接続先ポート
	User     string `required:"true"` // DB接続ユーザ
	Password string `required:"true"` // DB接続パスワード
	Database string `required:"true"` // データベース名
}

// init mongoDB設定を環境変数から取得します
func (conf *mongoConf) init() {
	err := envconfig.Process("mongo", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
