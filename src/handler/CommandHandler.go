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
	var cmdLine *exec.Cmd
	fileId := p.ByName("fileId")
	if filePath, ok := cfgKnight.CheckCfgWhiteList(fileId); ok {
		var cmdMix string
		//选择命令行
		switch r.Method {
		case http.MethodPut:
			cmdMix, _ = cmdKnight.CheckCmdWhiteList("dc-pl")
			cmds := strings.Split(cmdMix, " ")
			cmdLine = exec.Command(cmds[0], cmds[1], filePath, cmds[3])
			break
		case http.MethodPost:
			cmdMix, _ = cmdKnight.CheckCmdWhiteList("dc-u")
			cmds := strings.Split(cmdMix, " ")
			cmdLine = exec.Command(cmds[0], cmds[1], filePath, cmds[3], cmds[4])
			break
		case http.MethodDelete:
			cmdMix, _ = cmdKnight.CheckCmdWhiteList("dc-d")
			cmds := strings.Split(cmdMix, " ")
			cmdLine = exec.Command(cmds[0], cmds[1], filePath, cmds[3])
		}
		//执行命令行，并获取返回结果
		if cmdLine != nil {
			outputs, err := cmdLine.CombinedOutput()
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
