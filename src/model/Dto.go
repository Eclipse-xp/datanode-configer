package model

import (
	. "../constants"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type BaseDto struct {
	Code    int
	Message string
}

type RespDto struct {
	BaseDto
	Data interface{}
}

//启动容器dto
type ContainerRunReqDto struct {
	ContainerName string
	HostConfig container.HostConfig
	NetWorkingConfig network.NetworkingConfig
}

type DtoGenerator struct {}

func (dg DtoGenerator) Success() RespDto {
	res := RespDto{}
	res.Code = RespCodeSuccess
	res.Message = "Success"
	return res
}

func (dg DtoGenerator) SuccessWithData(data interface{}) RespDto {
	res := RespDto{Data: data}
	res.Code = RespCodeSuccess
	res.Message = "Success"
	return res
}

func (dg DtoGenerator) Fail() RespDto {
	res := RespDto{}
	res.Code = RespCodeFail
	res.Message = "Failed"
	return res
}

func (dg DtoGenerator) FailWithContent(code int, message string) RespDto {
	res := RespDto{}
	res.Code = RespCodeFail
	res.Message = message
	return res
}
