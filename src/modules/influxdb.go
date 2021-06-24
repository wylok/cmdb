package modules

import (
	"cmdb-go/kits"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
)

func ConnInflux() client.Client {
	// 从apollo获取数据库配置信息
	v := ApolloConfig("influxdb_go_config")
	conf := kits.StringToMap(v.(string))
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     conf["Addr"].(string),
		Username: conf["Username"].(string),
		Password: conf["Password"].(string),
	})
	if err != nil {
		log.Fatal(err)
	} else {
		println("Influxdb client ok")
	}
	return cli
}
