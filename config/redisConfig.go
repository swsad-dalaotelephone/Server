package config

type RedisConfigType struct {
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	Password   string `json:"Password"`
	SessionKey string `json:"SessionKey"`
}

//redis configuration
var RedisConfig RedisConfigType
