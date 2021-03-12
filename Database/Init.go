package Database

import (
	"MessageBoard/Model"
	"MessageBoard/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(sql *config.Sql) *gorm.DB {
	dsn := sql.SqlUserName+":"+sql.SqlUserPwd+"@tcp("+sql.SqlAddr+")/"+sql.SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	//"root:root@tcp(127.0.0.1:3306)/留言板?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Model.Message{},Model.User{})
	return db
}
