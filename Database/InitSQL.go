package Database

import (
	"MessageBoard/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitSQL(SQL *config.Sql) *sql.DB {
	//sql.Register("mysql", &mysql.MySQLDriver{})
	//dsn := SQL.SqlUserName+":"+SQL.SqlUserPwd+"@tcp("+SQL.SqlAddr+")/"+SQL.SqlName+"?charset=utf8"
	dsn := SQL.SqlUserName+":"+SQL.SqlUserPwd+"@tcp("+SQL.SqlAddr+")/"+SQL.SqlName+"?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return nil
	}
	err=db.Ping()//测试链接是否有效
	if err!=nil {
		fmt.Println("failed to ping database:", err.Error())
		return nil
	}
	return db
}

