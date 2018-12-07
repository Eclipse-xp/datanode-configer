//配置文件安全守护
package safe

import . "../constants"

type CfgKnight struct {

}

func (ck CfgKnight) CheckCfgWhiteList(fileId string) (string, bool) {
	filePath := CfgWhiteList[fileId]
	return filePath, filePath != ""
}