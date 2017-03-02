package controller

import (
	"fmt"
	"github.com/zhja/gonoc"
	"encoding/json"
)

type LogbaseController struct {
	//gonoc.Controller
}

func (this *LogbaseController) Post() {
	//panic("test..........................")
	var type_name, data string
	type_name = gonoc.Requests.Get("type")
	data = gonoc.Requests.Get("data")
	returnData := gonoc.MQP.Post(type_name, data)
	//返回值转json
	rss_json, _ := json.Marshal(returnData)
	fmt.Fprintf(gonoc.Requests.W, string(rss_json))
}

func (this *LogbaseController) AllFailPost() {
	var type_name, status_id string
	status_id = gonoc.Requests.Get("status_id")
	type_name = gonoc.Requests.Get("type_name")
	returnData := gonoc.MQP.AllFailPost(status_id, type_name)
	//返回值转json
	rss_json, _ := json.Marshal(returnData)
	fmt.Fprintf(gonoc.Requests.W, string(rss_json))
}

func (this *LogbaseController) OneMqPush() {
	var id string
	id = gonoc.Requests.Get("id")
	returnData := gonoc.MQP.OnePush(id)
	//返回值转json
	rss_json, _ := json.Marshal(returnData)
	fmt.Fprintf(gonoc.Requests.W, string(rss_json))
}

func (this *LogbaseController) Jump() {
	var id string
	id = gonoc.Requests.Get("id")
	returnData := gonoc.MQP.Jump(id)
	//返回值转json
	rss_json, _ := json.Marshal(returnData)
	fmt.Fprintf(gonoc.Requests.W, string(rss_json))
}