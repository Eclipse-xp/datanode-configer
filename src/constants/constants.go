package constants

const (
	DockerClientVersion = "1.39"
	RespCodeSuccess     = 0
	RespCodeFail        = 1
)

var (
	//配置文件白名单
	CfgWhiteList map[string]string
	//命令行白名单
	CmdWitheList map[string]string
)

func init() {
	CfgWhiteList = map[string]string{
		"appconfig":            "/data/datanode/appconfig.properties",
		"bundle-compose":       "/data/datanode/compose/bundle/docker-compose.yml",
		"bundle-compose-redis": "/data/datanode/compose/bundle/docker-compose-redis.yml",
	}
	CmdWitheList = map[string]string{
		"dc-u":  "docker-compose -f ? up -d",
		"dc-pl": "docker-compose -f ? pull",
	}
}
