package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"./handler"
	_ "./constants"
)

func main() {
	router := httprouter.New()
	//获取配置文件列表（这里可考虑将用数据库存储列表）
	router.GET("/config", nil)
	//查看配置文件信息
	router.GET("/config/:fileId", handler.ConfigInfoHandler)
	//修改配置文件内容
	router.POST("/config/:fileId", nil)

	//更新镜像
	router.POST("/image", nil)
	//获取容器列表
	router.GET("/container", handler.Containers)
	//查看容器具体信息
	router.GET("/container/:containerId", nil)
	//启动容器
	router.POST("/container/:containerId", nil)
	//停止容器
	router.DELETE("/container/:containerId", handler.StopContainer)

	//通过docker-compose启动服务
	router.POST("/compose/:fileId", handler.StopContainer)

	http.ListenAndServe(":8080", router)
}
