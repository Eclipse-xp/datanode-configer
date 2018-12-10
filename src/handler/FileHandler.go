//配置文件处理器
package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"../model"
	"encoding/json"
	"io/ioutil"
	. "../constants"
	"../safe"
)

var (
	cfgKnight = safe.CfgKnight{}
)

func ConfigList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	resp := model.DtoGenerator{}.SuccessWithData(CfgWhiteList)
	json.NewEncoder(w).Encode(resp)
}

func ConfigInfoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fileId := p.ByName("fileId")
	if filePath, ok := cfgKnight.CheckCfgWhiteList(fileId); ok {
		//读文件
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("读文件异常", err.Error())
			resp := model.DtoGenerator{}.Fail()
			json.NewEncoder(w).Encode(resp)
		}
		content := string(buf)
		//拼装返回结果
		resp := model.DtoGenerator{}.SuccessWithData(content)
		//序列化并返回
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := model.DtoGenerator{}.FailWithContent(RespCodeFail, "不支持操作该文件")
		json.NewEncoder(w).Encode(resp)
	}
}