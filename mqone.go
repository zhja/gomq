package main

import (
	"github.com/zhja/gonoc"
	"./config"
	"flag"
	//"fmt"
	//"os"
)

func main() {
	config.CInit()
	gonoc.Open(config.Con["db_type"], config.Con["db_config"])
	gonoc.CreateMQ(config.Con["mq_type"], config.Con["mq_host"], config.Con["mq_pwd"])
	gonoc.MQW.Init()
	key := flag.String("key", "", "mq type")
	flag.Parse()
	gonoc.MQC.BranchOne(*key)
}