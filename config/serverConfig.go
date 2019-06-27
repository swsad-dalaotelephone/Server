package config

type ServerConfigType struct {
	Debug         bool   `json:"Debug"`
	FilePath      string `json:"FilePath"`
	ImgPath       string `json:"ImgPath"`
	Port          int    `json:"Port"`
	PageSize      int    `json:"PageSize"`
	MaxPageSize   int    `json:"MaxPageSize"`
	MinPageSize   int    `json:"MinPageSize"`
	MaxNameLen    int    `json:"MaxNameLen"`
	MaxRemarkLen  int    `json:"MaxRemarkLen"`
	MaxContentLen int    `json:"MaxContentLen"`
	Timeout       int    `json:"Timeout"`
}

// server configuration
var ServerConfig ServerConfigType
