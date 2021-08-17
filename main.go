package main

import (
	"MessageBoard/Routers"
	"MessageBoard/config"
	"MessageBoard/services/DC"
)

func main()  {
	conf:=config.Init()//获取服务器配置
	DC.InitGorm(&conf.Sql)
	go DC.Run()
	DCClient:=DC.NewClient()
	server:=Routers.BuildRouter(DCClient)
	server.Run(conf.Addr)
}