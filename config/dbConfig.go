package config

type DBConfigType struct {
	Dialect      string `json:"Dialect"`
	Database     string `json:"Database"`
	User         string `json:"Username"`
	Password     string `json:"Password"`
	Host         string `json:"Host"`
	Port         int    `json:"Port"`
	Charset      string `json:"Charset"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
	URL          string
}

// database configuration
var DBConfig DBConfigType
