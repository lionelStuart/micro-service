package redis

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("[Init] already init redis ..")
		return
	}

	redisConfig := config.GetRedisConfig()

	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Log("init redis ..")

		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Log("init redis: sentinel model ...")
			initSentinel(redisConfig)
		} else {
			log.Log("init redis: normal model ...")
			initSingle(redisConfig)
		}

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatalf("redis fail ping:%s", err.Error())
		}
		log.Log("init redis ,ping ..")
		log.Logf("init redis ,ping .. %s", pong)
	}
}

func GetRedis() *redis.Client {
	return client
}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(),
		DB:       redisConfig.GetDBNum(),
	})
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		Password:      redisConfig.GetPassword(),
		DB:            redisConfig.GetDBNum(),
	})
}
