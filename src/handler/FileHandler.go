package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"fmt"
	"bufio"
	"model"
	"encoding/json"
)

//配置文件处理器

var fileMap map[string]string

func ConfigInfoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fileId := p.ByName("fileId")
	if filePath := fileMap[fileId]; filePath != "" {
		//读文件
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("读文件异常", err.Error())
			resp := model.DtoGenerator{}.Fail()
			json.NewEncoder(w).Encode(resp)
		}
		defer file.Close()
		fileReader := bufio.NewReader(file)
		content := ""
		for {
			line ,_ := fileReader.ReadString('\n')
			content += line
		}
		//拼装返回结果
		resp := model.DtoGenerator{}.SuccessWithData(content)
		//序列化并返回
		json.NewEncoder(w).Encode(resp)
	}
}

func init() {
	fileMap = map[string]string{
		"appconfig": "/data/datanode/appconfig.properties",
		"bundle-compose": "/data/datanode/compose/bundle/docker-compose.yml",
	}
}