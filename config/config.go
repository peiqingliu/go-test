package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// 初始化配置文件
func LogInfo()  {
	file := "./" + time.Now().Format("2020-10-23") + ".log"
	logFile,_ :=os.OpenFile(file,os.O_RDWR | os.O_CREATE| os.O_APPEND,0766)
	log.SetFlags(log.Ldate | log.Ltime| log.Lshortfile)
	log.SetOutput(logFile)
}

//Init 读取初始化配置文件
func Init() error {

	err :=Config()
	if err != nil{
		return err
	}

	//初始化日志包
	LogInfo()
	return nil
}

//Config 解析配置文件
func Config()  error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}