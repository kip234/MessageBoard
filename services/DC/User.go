package DC

import "fmt"

type User struct {
	Uid int `gorm:"primaryKey"`
	Name string	`gorm:"string not null"`//用户名
	Pwd string `gorm:"string not null"`//用户密码
}

//创建用户
func (u *User)Save(arg User,reply *User) (err error) {
	err= DB.Create(&arg).Error
	*reply=arg
	fmt.Println(err)
	return
}

//根据提供的UID读取用户信息
//主要用于密码比对
func (u *User)Load(uid int,reply *User) (err error) {
	err=DB.Where("Uid=?",uid).Find(reply).Error
	return
}

//判断密码是否正确，如果不正确返回false
func (u *User)PwdIsRight(arg User,reply *bool) error {
	tmp:= User{}
	DB.Where("Uid=?", arg.Uid).Find(&tmp)
	*reply=tmp.Pwd==arg.Pwd
	return DB.Error
}

//判断是否存在，如果不存在返回false
func (u *User)IsExist(arg User,reply *bool) error {
	tmp:= User{}
	DB.Where("Uid=?", arg.Uid).Find(&tmp)
	if nil == DB.Error {
		*reply=true
		return nil
	}
	*reply = false
	return DB.Error
}