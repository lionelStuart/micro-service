package jwt

import "micro-service/basic"

type Jwt struct {
	SecretKey string `json:"secret_key"`
}

func init() {
	basic.Register(initJwt)
}

func initJwt() {

}
