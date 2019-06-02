package taskController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/swsad-dalaotelephone/Server/models/task"
)

func TestPublishTask(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/task/publishTask",
		strings.NewReader("id=&publisher_id=576177e5-6496-4c87-9d5c-91ac04c23c79&name=测试"))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8; application/json")
	//cookie1 := &http.Cookie{Name: "baobaozhuan_cookie", Value: "MTU1OTQ0MTg4MnxOd3dBTkVNM1ExTkZRbFZIVGtoUlQwUk5OMVZCVUU5Tk5UZEdXVUpFUWtaV1VVWklWa2xCU1VkQlYxZEZObFEwV2tSRk5GQlFTVUU9fPiJeKNwWEJlxx14vtCmRizTQOVLCLtiove11GB1GPf1:"}
	//req.AddCookie(cookie1)
	resp, err := client.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}

	type Result struct {
		Msg  string         `json:"msg"`
		Task taskModel.Task `json:"task"`
	}

	result := &Result{}
	json.Unmarshal(body, result) //解析json字符串

	t.Log(result)
}
