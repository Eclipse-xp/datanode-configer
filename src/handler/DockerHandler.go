//docker处理器，用于管理容器，直接源码调用docker
package handler

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"../model"
	. "../constants"
	"github.com/docker/docker/api/types/container"
	"io/ioutil"
)

//docker ps
func Containers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取运行中的容器
	cli := getDockerClient(w)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
		json.NewEncoder(w).Encode(resp)
	}
	//拼装返回结果
	resp := model.DtoGenerator{}.SuccessWithData(containers)
	//序列化并返回
	json.NewEncoder(w).Encode(resp)

	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.Names, container.ID[:10], container.Image)
	}
}

//docker run
func RunContainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cli := getDockerClient(w)
	ctx := context.Background()
	//TODO 接收数据改为用对象接收
	var containerConfig model.ContainerRunReqDto
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &containerConfig)
	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: DockerRepo + p.ByName("name") + ":" + p.ByName("tag"),
	}, &containerConfig.HostConfig, &containerConfig.NetWorkingConfig, containerConfig.ContainerName)
	if err != nil {
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
		json.NewEncoder(w).Encode(resp)
	}
	if err := cli.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
		json.NewEncoder(w).Encode(resp)
	}
	resp := model.DtoGenerator{}.SuccessWithData(createResp.ID)
	json.NewEncoder(w).Encode(resp)
	fmt.Println(createResp.ID)
}

//docker stop
func StopContainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cli := getDockerClient(w)
	err := cli.ContainerStop(context.Background(), p.ByName("containerId"), nil)
	if err != nil {
		fmt.Println(err)
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
		json.NewEncoder(w).Encode(resp)
	}
	resp := model.DtoGenerator{}.Success()
	json.NewEncoder(w).Encode(resp)
}

func getDockerClient(w http.ResponseWriter) (*client.Client) {
	cli, err := client.NewClientWithOpts(client.WithVersion(DockerClientVersion))
	if err != nil {
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
		json.NewEncoder(w).Encode(resp)
	}
	return cli
}
