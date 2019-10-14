package config

import "strings"

type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}

type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}

type defaultRedisConfig struct {
	Enabled  bool          `json:"enabled"`
	Conn     string        `json:"conn"`
	Password string        `json:"password"`
	DBNum    int           `json:"dbNum"`
	Timeout  int           `json:"timeout"`
	sentinel redisSentinel `json:"sentinel"`
}

type redisSentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	Nodes   string `json:"nodes"`
	nodes   []string
}

func (r defaultRedisConfig) GetEnabled() bool {
	return r.Enabled
}

func (r defaultRedisConfig) GetConn() string {
	return r.Conn
}

func (r defaultRedisConfig) GetPassword() string {
	return r.Password
}

func (r defaultRedisConfig) GetDBNum() int {
	return r.DBNum
}

func (r defaultRedisConfig) GetSentinelConfig() RedisSentinelConfig {
	return r.sentinel
}

func (r redisSentinel) GetEnabled() bool {
	return r.Enabled
}

func (r redisSentinel) GetMaster() string {
	return r.Master
}

func (r redisSentinel) GetNodes() []string {
	if len(r.nodes) != 0 {
		for _, v := range strings.Split(r.Nodes, ",") {
			v = strings.TrimSpace(v)
			r.nodes = append(r.nodes, v)
		}
	}
	return r.nodes
}
