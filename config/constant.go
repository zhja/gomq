package config

import (
	//"fmt"
	//"strconv"
)

var Con map[string]string

func CInit() {
	rs := make(map[string]string)
	rs["db_type"] = "mysql"
	rs["db_config"] = "user:password@tcp(xxx.xxx.xxx.xxx:8000)/database?charset=utf8"
	rs["mq_type"] = "redis"
	rs["mq_host"] = "xxx.xxx.xxx.xxx:6379"
	rs["mq_pwd"] = ""
	rs["static_file"] = "/usr/local/gonoc/static"
	rs["static_route"] = "/static/"
	rs["server_host"] = ":8010"
	Con = rs
}
