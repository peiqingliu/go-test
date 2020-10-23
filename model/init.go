package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var DB  *sql.DB
func Init() error {
	var err error
	DB,err = sql.Open("mysql",viper.GetString("mysql.source_name"))
	if nil != err {
		return err
	}
	//设置最大超时时间
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	//连接
	err = DB.Ping()
	if nil != err {
		return err
	}else {
		log.Println("mysql startup normal !")
	}
	return nil
}
