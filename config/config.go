package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const ConfPath = "conf.json"//外部配置文件

type Sql struct {
	SqlName string		//数据库名
	SqlUserName string	//数据库登录用账户名
	SqlUserPwd string	//数据库登录用账户密码
	SqlAddr string		//数据库地址
}

type Conf struct {
	Sql
	Addr string	//服务器地址
}

func Init() (result Conf) {
	file,err:=os.Open(ConfPath)
	if err!=nil {
		panic(err)
	}
	buf,err:=ioutil.ReadAll(file)
	if err!=nil {
		panic(err)
	}
	json.Unmarshal(buf,&result)
	return
}

