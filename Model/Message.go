package Model

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type Message struct{
	Like int//点赞数
	Content string `gorm:"string not null",binding:"required"`//内容
	Mid int `gorm:"primaryKey",binding:"required`//自己的身份标识
	Pid int//上一级ID
	Kids []Message `sql:"-",gorm:"-"`//子级
	Uid int //归属的用户
}

func insertMessage(db *sql.DB,m *Message) (err error) {
	if nil==db||nil==m {
		err=fmt.Errorf("insertMessage param wrong")
		return
	}
	_,err=db.Exec("INSERT INTO `" +
		"messages`(`like`,`content`,`mid`,`pid`,`uid`) VALUES("+
		strconv.Itoa(m.Like)+","+
		"\""+m.Content+"\""+","+
		strconv.Itoa(m.Mid)+","+
		strconv.Itoa(m.Pid)+","+
		strconv.Itoa(m.Uid)+
		");")
	return
}

//查找符合条件的记录
func findMessage(db *sql.DB,condition string,m *Message) error {
	if nil == db || nil == m {
		return fmt.Errorf("findMessage params wrong")
	}
	k, err := db.Query("SELECT * FROM messages WHERE " + condition+";")
	if nil!=err {
		fmt.Println("findMessage db.Query error:",err.Error())
		return err
	}
	if nil==k{
		fmt.Println("findMessage db.Query error: k=nil")
		return err
	}

	if ok:=k.Next();ok {
		k.Scan(&m.Like,&m.Content, &m.Mid, &m.Pid,&m.Uid)
	}
	defer k.Close()
	return err
}

//查找符合条件的记录
func findMessageAll(db *sql.DB,condition string,m *[]Message) error {
	if nil == db || nil == m {
		return fmt.Errorf("findMessageAll params wrong")
	}
	k, err := db.Query("SELECT * FROM messages WHERE " + condition+";")
	defer k.Close()
	for k.Next() {
		tmp:=Message{}
		k.Scan(&tmp.Like, &tmp.Content,  &tmp.Mid, &tmp.Pid,&tmp.Uid)
		*m=append(*m,tmp)
	}
	return err
}

func (m *Message)Save(db interface{}) (err error) {
	if value,ok:=db.(*gorm.DB);ok{//当前使用Gorm
		err=value.Create(m).Error
	}else if value,ok:=db.(*sql.DB);ok{
		err=insertMessage(value,m)
	}else {
		 err=fmt.Errorf("paramter wrong")
	}
	return
}

func Like(id int,db interface{}) (err error) {
	var tmp Message
	if value,ok:=db.(*gorm.DB);ok {//当前使用Gorm
		if err = value.Where("Mid=?", id).Find(&tmp).Error; err != nil {
			return
		}
		tmp.Like += 1
		err = value.Save(&tmp).Error
	}else if value,ok:=db.(*sql.DB);ok{
		if err=findMessage(value,"Mid="+strconv.Itoa(id),&tmp);err!=nil{
			return
		}
		tmp.Like+=1
		err=insertMessage(value,&tmp)
	}else {
		err=fmt.Errorf("paramter wrong")
	}
	return
}

//用于处理数据库中的数据，
//var wg sync.WaitGroup
//产生对应ID的子级信息列表
func kidContent(id int,db interface{}) (result []Message) {
	if value,ok:=db.(*gorm.DB);ok{//当前使用Gorm
		value.Where("Pid=?", id).Find(&result)
	}else if value,ok:=db.(*sql.DB);ok{
		findMessageAll(value,"Pid="+strconv.Itoa(id),&result)
	}
	return result
}

//根据id生成一颗树
func buildTree(id int,root *Message,db interface{}) {
	var wg sync.WaitGroup
	if value,ok:=db.(*gorm.DB);ok{//当前使用Gorm
		value.Where("Mid=?", id).Find(root)
		root.Kids = kidContent(id, db)
	}else if value,ok:=db.(*sql.DB);ok{
		findMessage(value,"Mid="+strconv.Itoa(id),root)
		root.Kids = kidContent(id, db)
	}
	for i,_:=range root.Kids{
		wg.Add(1)
		/*go*/func() {//=================递归并发===好想吹一手
			defer wg.Done()
			buildTree(root.Kids[i].Mid, &root.Kids[i], db)
		}()
	}
	wg.Wait()
}
//生成一个...不知道怎么描述
func GetContent(db interface{}) (result []Message) {
	if value,ok:=db.(*gorm.DB);ok {//当前使用Gorm
		value.Where("Pid=?",new(int)).Find(&result)
	}else if value,ok:=db.(*sql.DB);ok {
		findMessageAll(value,"Pid="+strconv.Itoa(0),&result)
	}
	for i,_:=range result{
		buildTree(result[i].Mid,&result[i],db)
	}
	//wg.Wait()
	return
}