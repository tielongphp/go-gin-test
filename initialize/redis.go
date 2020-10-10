package initialize

import (
	"fmt"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	"go-gin-test/global"
)

func RedisInit() {
	redisCfg, err := global.CTX_CONFIG.GetCfg().GetSection("redis")
	var client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Key("Addr").String(),
		Password: redisCfg.Key("Password").String(), // no password set
		//DB:       redisCfg.Key("DB").Int(),          // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Error(err)
	} else {
		fmt.Println(pong)
		log.Info("redis connect ping response:", pong)
		global.REDIS = client
	}
}
