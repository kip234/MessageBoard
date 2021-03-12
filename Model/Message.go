package Model

import (
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name string	`binding:"required"`
}

type Message struct{
	Like uint							  										//点赞数
	Content string	`gorm:"string not null",binding:"required"`						//内容
	ID uint								  										//自己的身份标识
	Pid uint							  										//上一级ID
	Kids []Message	`gorm:"-"`	 												 //子级
}

func Save(m *Message,db *gorm.DB) (err error) {
	err=db.Create(&m).Error
	return
}
func Like(id uint,db *gorm.DB) (err error) {
	var tmp Message
	if err=db.Where("id=?",id).Find(&tmp).Error;err!=nil{
		return
	}
	tmp.Like+=1
	err=db.Save(&tmp).Error
	return
}

//用于处理数据库中的数据，
//var wg sync.WaitGroup
//产生对应ID的子级信息列表
func kidContent(id uint,db *gorm.DB) (result []Message) {
	db.Where("Pid=?",id).Find(&result)
	return result
}

//根据id生成一颗树
func buildTree(id uint,root *Message,db *gorm.DB) {
	var wg sync.WaitGroup
	db.Where("id=?",id).Find(root)
	root.Kids=kidContent(id,db)
	for i,_:=range root.Kids{
		wg.Add(1)
		func() {//=================递归并发===好想吹一手
			defer wg.Done()
			buildTree(root.Kids[i].ID, &root.Kids[i], db)
		}()
	}
	wg.Wait()
}
//生成一个...不知道怎么描述
func GetContent(db *gorm.DB) (result []Message) {
	db.Where("Pid=?",new(int)).Find(&result)
	for i,_:=range result{
		buildTree(result[i].ID,&result[i],db)
	}
	//wg.Wait()
	return
}