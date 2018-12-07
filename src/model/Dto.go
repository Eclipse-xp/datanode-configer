package model

import . "../constants"

type BaseDto struct {
	Code    int
	Message string
}

type RespDto struct {
	BaseDto
	Data interface{}
}

type DtoGenerator struct {}

func (dg DtoGenerator) Success() RespDto {
	res := RespDto{}
	res.Code = RESP_CODE_SUCCESS
	res.Message = "Success"
	return res
}

func (dg DtoGenerator) SuccessWithData(data interface{}) RespDto {
	res := RespDto{Data: data}
	res.Code = RESP_CODE_SUCCESS
	res.Message = "Success"
	return res
}

func (dg DtoGenerator) Fail() RespDto {
	res := RespDto{}
	res.Code = RESP_CODE_FAIL
	res.Message = "Failed"
	return res
}

func (dg DtoGenerator) FailWithContent(code int, message string) RespDto {
	res := RespDto{}
	res.Code = RESP_CODE_FAIL
	res.Message = message
	return res
}
