package router

import (
	"github.com/gin-gonic/gin"
	"go-test/router/middleware"
	"go-test/service"
	"net/http"
)

//设置路由

//初始化路由
func InitRouter(g *gin.Engine)  {
	//数组
	middlewares := []gin.HandlerFunc{}
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 路由请求
	router :=g.Group("/user")

	{
		router.POST("/addUser",service.AddUser)
		router.POST("/addUserJson",service.AddUserJson)
		router.GET("/selectUser",service.SelectUser)
	}
}