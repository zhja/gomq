package gonoc

import (
	//"fmt"
	"strconv"
	"time"
)

var MQP Producer

type Producer struct {
}

func (this *Producer) Post(type_name string, data string) map[string]string {
	//获取type_id,判断是否合法
	Db.AndWhere("=", "name", type_name)
	Db.Field("id, name")
	Db.Select("mq_type")
	rows := Db.QueryRow()
	var id int
	var name string
	returnData := make(map[string]string)
	err := rows.Scan(&id, &name);
	if err != nil {
		returnData["status"] = "false"
		returnData["msg"] = "mq_type error"
	} else {
		t := time.Now()
		//保存mysql
		var add_mq_list mq_list
		add_mq_list.Type_name = type_name
		add_mq_list.Data = data
		add_mq_list.Create_time = t.Format("2006-01-02 15:04:02")
		add_mq_list.Update_time = t.Format("2006-01-02 15:04:02")
		add_mq_list.Status_id = 0
		Db.Save(&add_mq_list)
		//保存redis
		MQ.Push(type_name, strconv.Itoa(add_mq_list.Id))
		returnData["status"] = "true"
		returnData["msg"] = strconv.Itoa(add_mq_list.Id)
		//起对应协程
		//MQC.ProcessOne(type_name)
	}
	return returnData
}

//获取所有未成功的任务重新导入队列
func (this *Producer) AllFailPost(status_id string, type_name string) map[string]string {
	//获取type_id,判断是否合法
	if status_id != "" {
		Db.AndWhere("=", "status_id", status_id)
	}
	if type_name != "" {
		Db.AndWhere("=", "type_name", type_name)
	}
	Db.Field("id, type_name")
	Db.Select("mq_list")
	rows, _ := Db.Query()
	for _, item := range rows {
		MQ.Push(Db.GetValue(&item, "type_name"), Db.GetValue(&item, "id"))
		//起对应协程
		//MQC.ProcessOne(Db.GetValue(&item, "type_name"))
	}
	returnData := make(map[string]string)
	returnData["status"] = "true"
	returnData["msg"] = "success"
	return returnData
}

//根据ID单独提交到队列
func (this *Producer) OnePush(id string) map[string]string {
	if id != "" {
		Db.AndWhere("=", "id", id)
	}
	Db.Field("id, type_name")
	Db.Select("mq_list")
	rows, _ := Db.Query()
	for _, item := range rows {
		MQ.Push(Db.GetValue(&item, "type_name"), Db.GetValue(&item, "id"))
		//起对应协程
		//MQC.ProcessOne(Db.GetValue(&item, "type_name"))
	}
	returnData := make(map[string]string)
	returnData["status"] = "true"
	returnData["msg"] = "success"
	return returnData
}

//根据ID单独提交到队列
func (this *Producer) Jump(id string) map[string]string {
	if id != "" {
		Db.AndWhere("=", "id", id)
	}
	Db.Field("id, type_name")
	Db.Select("mq_list")
	rows, _ := Db.Query()
	for _, item := range rows {
		MQ.Jump(Db.GetValue(&item, "type_name"), Db.GetValue(&item, "id"))
		//起对应协程
		//MQC.ProcessOne(Db.GetValue(&item, "type_name"))
	}
	returnData := make(map[string]string)
	returnData["status"] = "true"
	returnData["msg"] = "success"
	return returnData
}