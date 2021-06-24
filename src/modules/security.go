package modules

import (
	"cmdb-go/kits"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := map[string]string{"Authorization": "None", "token": "None"}
		// 从apollo获取数据库配置信息
		session := sessions.Default(c)
		v := ApolloConfig("api_url_config")
		conf := kits.StringToMap(v.(string))
		url := conf["token_verify"].(string)
		v = ApolloConfig("public_token")
		public := kits.StringToMap(v.(string))
		token := c.Request.Header.Get("token")
		// 公共token不进行验证
		if token == public["token"].(string) {
			session.Set("headers", headers)
			c.Next()
			return
		} else {
			req := HttpRequest.NewRequest()
			req.SetHeaders(map[string]string{
				"token": token,
			})
			resp, err := req.Post(url)
			if err != nil {
				kits.Log(err.Error(), "Error")
			} else {
				if resp != nil {
					body, err := resp.Export()
					if err == nil {
						println(body)
						session.Set("headers", map[string]string{"token": token})
						c.Next()
						return
					}
				}
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false, "message": "Unauthorized", "data": map[string]string{}})
		c.Abort()
		return
	}
}
