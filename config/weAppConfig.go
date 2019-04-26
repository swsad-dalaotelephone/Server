package config

type WeAppConfigType struct {
	AppID  string `json:"AppID"`
	Secret string `json:"Secret"`
}

//weApp configuration
var WeAppConfig WeAppConfigType
