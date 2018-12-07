package constants


const (
	DOCKER_CLIENT_VERSION = "1.39"
	RESP_CODE_SUCCESS = 0
	RESP_CODE_FAIL = 1
)

//配置文件白名单
var CfgWhiteList map[string]string
//TODO 命令行白名单

func init() {
	CfgWhiteList = map[string]string{
		"appconfig": "/data/datanode/appconfig.properties",
		"bundle-compose": "/data/datanode/compose/bundle/docker-compose.yml",
	}
}