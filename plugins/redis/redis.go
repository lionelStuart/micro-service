package redis

import (
	r "github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic"
	"micro-service/basic/config"
	"strings"
	"sync"
)

var (
	client *r.Client
	m      sync.RWMutex
	inited bool
)

type redis struct {
	Enabled  bool           `json:"enabled"`
	Conn     string         `json:"conn"`
	Password string         `json:"password"`
	DBNum    int            `json:"dbNum"`
	Timeout  string         `json:"timeout"`
	Sentinel *RedisSentinel `json:"sentinel"`
}

type RedisSentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	XNodes  string `json:"nodes"`
	nodes   []string
}

func init() {
	basic.Register(initRedis)
}

func initRedis() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("[Init] already init redis ..")
		return
	}

	c := config.C()
	cfg := &redis{}
	err := c.App("redis", cfg)
	if err != nil {
		log.Logf("[InitRedis] %s", err)
	}

	if !cfg.Enabled {
		log.Logf("[InitRedis] not Enabled")
		return
	}

	if cfg.Sentinel != nil && cfg.Sentinel.Enabled {
		log.Log("[initRedis] redis start with sentinel")
		initSentinel(cfg)
	} else {
		log.Log("[initRefid] redis start normal")
		initSingle(cfg)
	}

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis fail ping:%s", err.Error())
	}
	log.Log("init redis ,ping ..")
	log.Logf("init redis ,ping .. %s", pong)

}

func GetRedis() *r.Client {
	return client
}

func initSingle(redisConfig *redis) {

	client = r.NewClient(&r.Options{
		Addr:     redisConfig.Conn,
		Password: redisConfig.Password,
		DB:       redisConfig.DBNum,
	})
}

func initSentinel(redisConfig *redis) {
	client = r.NewFailoverClient(&r.FailoverOptions{
		MasterName:    redisConfig.Sentinel.Master,
		SentinelAddrs: redisConfig.Sentinel.GetNodes(),
		Password:      redisConfig.Password,
		DB:            redisConfig.DBNum,
	})
}

func (s *RedisSentinel) GetNodes() []string {
	if len(s.XNodes) != 0 {
		for _, v := range strings.Split(s.XNodes, ",") {
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes, v)
		}
	}
	return s.nodes
}
