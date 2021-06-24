package config

const CryptKey = "4098879a2529ca11b8675505ahf88a2d"
const ApiAssetKey = "cmdb_api_asset_key"
const LogFile = "/opt/logs/cmdb.log"
const InfoFile = "/opt/logs/cmdb.info"
const ErrorFile = "/opt/logs/cmdb.error"
const DebugFile = "/opt/logs/cmdb.debug"
const ApolloUrl = "http://apollo.service.xxx:8080/"
const ApolloAppId = "xxx"

var AssetTypes = map[string]string{"physical": "物理机", "vmware": "VM虚机",
	"aliyun_ecs": "阿里云主机"}

var ValidAgents = [5]string{"5dd78d14f4f3a4f1609d5739"}

type Apollo struct {
	ConfServerUrl       string `yaml:"conf_server_url"`
	AppId               string `yaml:"app_id"`
	BackupRedisHost     string `yaml:"backup_redis_host"`
	BackupRedisPort     int    `yaml:"backup_redis_port"`
	BackupRedisPassword string `yaml:"backup_redis_password"`
}

type Log struct {
	InfoFile  string `yaml:"info_file"`
	ErrorFile string `yaml:"error_file"`
	DebugFile string `yaml:"debug_file"`
}
