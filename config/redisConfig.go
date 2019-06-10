package config

import "time"

type RedisConfigType struct {
	Host        string        `json:"Host"`
	Port        int           `json:"Port"`
	Password    string        `json:"Password"`
	SessionKey  string        `json:"SessionKey"`
	MaxIdle     int           `json:"MaxIdle"`
	MaxActive   int           `json:"MaxActive"`
	IdleTimeout time.Duration `json:"time.Duration"`
}

//redis configuration
var RedisConfig RedisConfigType
