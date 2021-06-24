package modules

import (
	"cmdb-go/kits"
	"encoding/json"
	"github.com/go-redis/redis"
)

var RC *redis.ClusterClient

func init() {
	// 从apollo获取数据库配置信息
	var Addrs []string
	v := ApolloConfig("redis_cluster_config")
	conf := kits.StringToMap(v.(string))
	c, _ := json.Marshal(conf["Addrs"].([]interface{}))
	_ = json.Unmarshal(c, &Addrs)
	// 初始化redis连接
	RC = redis.NewClusterClient(&redis.ClusterOptions{Addrs: Addrs, Password: ""})
	err := RC.Ping().Err()
	if err == nil {
		println("Redis cluster ok!")
	} else {
		println(err.Error())
	}
	// 连接rabbitmq
	queueExchange := &RabbitMQExchange{
		"queue_go",
		"msg",
		"exchange",
		"topic",
	}
	mq := RabbitMQNew(queueExchange)
	mq.Start("xxx")
	ConnInflux()
}
