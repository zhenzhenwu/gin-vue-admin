package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})

	//client := redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs:    []string{"101.34.247.33:6379", "122.51.229.32:6379"},
	//	Password: redisCfg.Password,
	//})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
