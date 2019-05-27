package config

import (
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/goinggo/mapstructure"
)

// all configuration
var jsonData map[string]interface{}

func readJSON() {
	GOPATH := os.Getenv("GOPATH")
	file := GOPATH + "/src/github.com/swsad-dalaotelephone/Server/config.json"
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.ErrorLog.Println(err.Error())
	}
	configStr := string(bytes[:])

	//delete all comment in config.json
	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initDB() {
	if err := mapstructure.Decode(jsonData["database"].(map[string]interface{}), &DBConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
	//set database url
	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
}

func initServer() {
	if err := mapstructure.Decode(jsonData["apiServer"].(map[string]interface{}), &ServerConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initWeApp() {
	if err := mapstructure.Decode(jsonData["weApp"].(map[string]interface{}), &WeAppConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initRedis() {
	if err := mapstructure.Decode(jsonData["redis"].(map[string]interface{}), &RedisConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initCookie() {
	if err := mapstructure.Decode(jsonData["cookie"].(map[string]interface{}), &CookieConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func init() {
	readJSON()
	initDB()
	initServer()
	initWeApp()
	initRedis()
	initCookie()
}
