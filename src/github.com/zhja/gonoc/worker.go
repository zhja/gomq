package gonoc

import (
	//"fmt"
	"strconv"
)

var MQW Worker

type Worker struct {
	Config map[string]map[string]string
}

func (this *Worker) Init() {
	//获取所有开启状态的队列配置信息，缓存
	Db.AndWhere("=", "status_id", strconv.Itoa(1))
	Db.AndWhere("!=", "parent_id", strconv.Itoa(0))
	Db.Field("id, name, parent_id, type, sleep, api, callback_api, status_id")
	Db.Select("mq_type")
	rows, _ := Db.Query()
	this.Config = make(map[string]map[string]string)
	for _, item := range rows {
		name := Db.GetValue(&item, "name")
		rs := make(map[string]string) 
		rs["id"] = Db.GetValue(&item, "id")
		rs["name"] = Db.GetValue(&item, "name")
		rs["parent_id"] = Db.GetValue(&item, "parent_id")
		rs["type"] = Db.GetValue(&item, "type")
		rs["sleep"] = Db.GetValue(&item, "sleep")
		rs["api"] = Db.GetValue(&item, "api")
		rs["callback_api"] = Db.GetValue(&item, "callback_api")
		rs["status_id"] = Db.GetValue(&item, "status_id")
		this.Config[name] = rs
		//创建队列监控数据
		MQ.Push(name + "Monitor", strconv.Itoa(MqUnix()))
	}
}