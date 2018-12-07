//执行命令安全守护者
package safe

type CmdKnight struct {}

func (ck CmdKnight) CheckCmdWhiteList() bool {
	return false
}