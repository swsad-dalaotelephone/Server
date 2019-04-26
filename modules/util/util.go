package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

//encode struct to json
func StructToJson(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}

//string md5
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
