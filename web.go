package main

import (
	"github.com/zhja/gonoc"
	"./router"
	"./config"
	"fmt"
)

func main() {
	config.CInit()
	router.Init()
	gonoc.Open(config.Con["db_type"], config.Con["db_config"])
	fmt.Println("Mysql success")

	gonoc.CreateMQ(config.Con["mq_type"], config.Con["mq_host"], config.Con["mq_pwd"])
	fmt.Println("MQ success")

	gonoc.MQW.Init()
	fmt.Println("MQ Worker success")

	fmt.Println("WEB success")
    gonoc.Run(config.Con["static_file"], config.Con["static_route"], config.Con["server_host"])
}
