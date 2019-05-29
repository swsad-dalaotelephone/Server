package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

//encode struct to json
func StructToJson(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}

func JsonToMap(jsonData []byte) (mapData map[string]interface{}) {
	json.Unmarshal(jsonData, &mapData)
	return mapData
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
