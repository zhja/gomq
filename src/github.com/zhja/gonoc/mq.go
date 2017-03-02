package gonoc

import (
	//"strconv"
	//"fmt"
)

type mq_list struct {
	Id int `PK`
	Type_name string
	Data string
	Create_time string
	Update_time string
	Status_id int
	Callback_data string
}

type mq_list_done struct {
	Id int `PK`
	Type_name string
	Data string
	Create_time string
	Update_time string
	Status_id int
	Callback_data string
}

var MQ MessageQueue

type MessageQueue struct {
	RedisPool
}

func CreateMQ(types string, server string, password string) {
    if types == "redis" {
    	MQ.Pool = RedisDb.CreateRedisPool(server, password)
    }
}

func (this *MessageQueue) Push(key string, id string) {
	err := this.Rpush(key, id)
	CheckError(err)
}

func (this *MessageQueue) Pop(key string) (rs *map[string]string, err error) {
	value, _ := this.Lpop(key)
	if value != "" {
		rs, err = this.Info(value)
	}
	return
}

func (this *MessageQueue) Info(value string) (*map[string]string, error) {
	Db.AndWhere("=", "id", value)
	Db.Field("id, type_name, data, create_time, status_id, callback_data")
	Db.Select("mq_list")
	rows, err := Db.QueryRows()
	return &rows, err
}

func (this *MessageQueue) SimplePop(key string) string {
	value, _ := this.Lpop(key)
	return value
}

//插队
func (this *MessageQueue) Jump(key string, value string) {
	_ = this.Lrem(key, value)
	err := this.Lpush(key, value)
	CheckError(err)
}