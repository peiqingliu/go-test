package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-test/config"
	"go-test/model"
	"go-test/router"
	"log"
)

func main()  {
	fmt.Println("hello,web")

	//1、初始化日志
	if err := config.Init();err != nil {
		//panic(err)
		log.Print(err)
	}

	//2、初始化数据库
	if err :=model.Init();err != nil {
		panic(err)
	}

	//3、设置请求模式
	gin.SetMode(viper.GetString("runmode"))
	g :=gin.New()
	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: %s\n", viper.GetString("addr"))
	if err := g.Run(viper.GetString("addr"));err != nil {log.Fatal("ListenAndServe:", err)}
}