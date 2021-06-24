package modules

import (
	Conf "cmdb-go/config"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
)

func ApolloConfig(Key string) interface{} {
	c := &config.AppConfig{
		AppID:          Conf.ApolloAppId,
		Cluster:        "dev",
		IP:             Conf.ApolloUrl,
		NamespaceName:  "opauto",
		IsBackupConfig: true,
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	cache := client.GetConfigCache(c.NamespaceName)
	v, _ := cache.Get(Key)
	return v
}
