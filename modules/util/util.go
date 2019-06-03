package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// encode struct to json(byte[])
func StructToJson(structModel interface{}) ([]byte, error) {
	json, err := json.Marshal(structModel)
	return json, err
}

// encode struct to json string
func StructToJsonStr(structModel interface{}) (string, error) {
	json, err := json.Marshal(structModel)
	return string(json), err
}

// decode json(byte[]) to map[string]interface{}
func JsonToMap(jsonData []byte) (mapData map[string]interface{}, err error) {
	err = json.Unmarshal(jsonData, &mapData)
	return mapData, err
}

//string md5
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

//get new uid
func GetNewUuid() string {
	//some version uuid.NewV4() return 2 results
	return uuid.Must(uuid.NewV4() /*, nil*/).String()
}
