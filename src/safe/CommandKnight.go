//执行命令安全守护者
package safe

import . "../constants"

type CmdKnight struct{}

func (ck CmdKnight) CheckCmdWhiteList(cmdId string) (string, bool) {
	cmd := CmdWitheList[cmdId]
	return cmd, cmd != ""
}
