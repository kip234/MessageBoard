package main

import (
	"MessageBoard/Database"
	"MessageBoard/Routers"
	"MessageBoard/config"
)

func main()  {
	conf:=config.Init()
	//db:=Database.InitGorm(&conf.Sql)
	db:=Database.InitSQL(&conf.Sql)

	server:=Routers.BuildRouter(db)
	server.Run(conf.Addr)
}
