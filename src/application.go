package main

import (
	"github.com/julienschmidt/httprouter"
	"./handler"
	_ "./constants"
	"net/http"
)
//FIXME content_type 过滤？
func main() {
	router := httprouter.New()
	//获取配置文件列表（TODO 这里可考虑将用数据库存储列表）
	router.GET("/config", handler.ConfigList)
	//查看配置文件信息
	router.GET("/config/:fileId", handler.ConfigInfoHandler)
	//修改配置文件内容
	router.POST("/config/:fileId", handler.ConfigUpdate)

	//获取镜像列表
	router.GET("/image", handler.ListImages)
	//拉取镜像pull image
	router.PUT("/image/:name/:tag", handler.PullImage)
	//启动容器docker run
	router.POST("/image/:name/:tag", handler.RunContainer)
	//删除镜像rmi
	router.DELETE("/image/:imageId", handler.DeleteImage)
	//获取容器列表
	router.GET("/container", handler.Containers)
	//查看容器具体信息
	router.GET("/container/:containerId", nil)
	//停止容器stop
	router.POST("/container/:containerId", handler.StopContainer)
	//删除容器rm
	router.DELETE("/container/:containerId", nil)

	//通过docker-compose up启动服务
	router.POST("/compose/:fileId", handler.ComposeHandler)
	//通过docker-compose pull更新镜像
	router.PUT("/compose/:fileId", handler.ComposeHandler)
	//通过docker-compose down停止服务
	router.DELETE("/compose/:fileId", handler.ComposeHandler)

	http.ListenAndServe(":8080", router)
}
