//命令行处理器
package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"../model"
	. "../constants"
	"../safe"
	"strings"
	"os/exec"
)

var cmdKnight = safe.CmdKnight{}

//docker-compose 通过系统命令行方式执行
func ComposeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp interface{}
	var cmdLine string
	fileId := p.ByName("fileId")
	if filePath, ok := cfgKnight.CheckCfgWhiteList(fileId); ok {
		var cmdPre string
		//选择命令行
		switch r.Method {
		case http.MethodGet:
			cmdPre, _ = cmdKnight.CheckCmdWhiteList("dc-pl")
			break
		case http.MethodPost:
			cmdPre, _ = cmdKnight.CheckCmdWhiteList("dc-u")
			break
		}
		//执行命令行，并获取返回结果
		cmdLine = strings.Replace(cmdPre, "?", filePath, 1)
		if cmdLine != "" {
			outputs, err := exec.Command(cmdLine).Output()
			resp = model.DtoGenerator{}.SuccessWithData(string(outputs))
			if err != nil {
				resp = model.DtoGenerator{}.FailWithContent(RespCodeFail, err.Error())
			}
		}
	} else {
		resp = model.DtoGenerator{}.FailWithContent(RespCodeFail, "不支持操作该文件")
	}
	json.NewEncoder(w).Encode(resp)
}
