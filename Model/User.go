package Model

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	Uid int `gorm:"primaryKey"`
	Name string	`gorm:"string not null"`//用户名
	Pwd string `gorm:"string not null"`//用户密码
}

func insertUser(db *sql.DB,m *User) (err error) {
	_,err=db.Exec("INSERT INTO `" +
		"Users`(`uid`,`Name`,`Pwd`) VALUES("+
		strconv.Itoa(m.Uid)+","+
		"\""+m.Name+"\""+","+
		"\""+m.Pwd+"\""+
		");")
	return
}

//查找符合条件的记录
func findUser(db *sql.DB,condition string,m *User) error {
	if nil == db || nil == m {
		return fmt.Errorf("params wrong")
	}
	k, err := db.Query("SELECT * FROM Users WHERE " + condition+";")
	defer k.Close()
	if ok:=k.Next();ok {
		k.Scan(&m.Uid, &m.Name, &m.Pwd)
	}else {
			return fmt.Errorf("database wrong")
		}
	return err
}

//创建用户
func (u *User)Save(db interface{}) (err error) {
	if value,ok:=db.(*gorm.DB);ok{
		err=value.Create(u).Error
	}else if value,ok:=db.(*sql.DB);ok {
		err=insertUser(value,u)
		findUser(value,"Name="+"\""+u.Name+"\"",u)//主要目的是为了返回新建用户的ID
	}else {
		err=fmt.Errorf("params wrong")
	}
	return
}

//根据提供的UID读取用户信息
//主要用于密码比对
func (u *User)Load(db interface{},uid int) (err error) {
	if value,ok:=db.(*gorm.DB);ok{
		err=value.Where("Uid=?",uid).Find(u).Error
	}else if value,ok:=db.(*sql.DB);ok {
		err=findUser(value,"Uid="+strconv.Itoa(uid),u)
	}else {
		err=fmt.Errorf("params wrong")
	}
	return
}

//判断密码是否正确，如果不正确返回false
func (u *User)PwdIsRight(db interface{}) bool {
	tmp:=User{}
	if valuer,ok:=db.(*gorm.DB);ok {
		valuer.Where("Uid=?", u.Uid).Find(&tmp)
	}else if valuer,ok:=db.(*sql.DB);ok{
		findUser(valuer,"Uid="+strconv.Itoa(u.Uid),&tmp)
	}
	return tmp.Pwd==u.Pwd
}

//判断是否存在，如果不存在返回false
func (u *User)IsExist(db interface{}) bool {
	tmp:=User{}
	if value,ok:=db.(*gorm.DB);ok {
		value.Where("Uid=?", u.Uid).Find(&tmp)
		if nil == value.Error {
			return true
		}
	}else if value,ok:=db.(*sql.DB);ok{
		err:=findUser(value,"Uid="+strconv.Itoa(u.Uid),&tmp)
		if nil==err{
			return true
		}
	}
	return false
}
