package config

type JwtConfig interface {
	GetSecretKey() string
}

type defaultJwtConfig struct {
	SecretKey string `json:"secret_key"`
}

func (m defaultJwtConfig) GetSecretKey() string {
	return m.SecretKey
}
