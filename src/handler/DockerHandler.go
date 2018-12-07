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
)

//docker处理器，用于管理容器，直接源码调用docker

func Containers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//获取运行中的容器(docker ps)
	cli, err := client.NewClientWithOpts(client.WithVersion(model.DOCKER_CLIENT_VERSION))
	if err != nil {
		resp := model.DtoGenerator{}.FailWithContent(model.RESP_CODE_FAIL, err.Error())
		json.NewEncoder(w).Encode(resp)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		resp := model.DtoGenerator{}.FailWithContent(model.RESP_CODE_FAIL, err.Error())
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
