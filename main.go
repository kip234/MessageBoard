package main

import (
	"MessageBoard/Database"
	"MessageBoard/Routers"
	"MessageBoard/config"
)

func main()  {
	conf:=config.Init()//获取服务器配置
	db:=Database.InitGorm(&conf.Sql)//使用gorm
	//db:=Database.InitSQL(&conf.Sql)//使用原生sql接口

	server:=Routers.BuildRouter(db)
	server.Run(conf.Addr)
}
