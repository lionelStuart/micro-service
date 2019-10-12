package config

type ConsulConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

func (c defaultConsulConfig) GetPort() int {
	return c.Port
}

func (c defaultConsulConfig) GetHost() string {
	return c.Host
}

func (c defaultConsulConfig) GetEnabled() bool {
	return c.Enabled
}
