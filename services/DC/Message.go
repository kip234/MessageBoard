package DC

import (
	"fmt"
	"sync"
)

type Message struct{
	Likes int       //点赞数
	Content string `gorm:"string not null" binding:"required"`//内容
	Mid int        `gorm:"primaryKey" binding:"required`//自己的身份标识
	Pid int        //上一级ID
	Kids []Message `gorm:"-"` //子级
	Uid int        //归属的用户
}

//保存
func (m *Message)Save(arg Message,reply *Message) (err error) {
	err= DB.Create(&arg).Error
	return
}

//点赞
func (m *Message)Like(id int,resp *Message) (err error) {
	var tmp Message
	if	err = DB.Where("Mid=?", id).Find(&tmp).Error; err != nil {
		fmt.Printf("1>Like:%s\n",err.Error())
		return
	}
	tmp.Likes += 1
	err = DB.Save(&tmp).Error
	if err != nil{
		fmt.Printf("1>Like:%s\n",err.Error())
	}
	return
}

//用于处理数据库中的数据，
//var wg sync.WaitGroup
//产生对应ID的子级信息列表
func kidContent(id int) (result []Message) {
	DB.Where("Pid=?", id).Find(&result)
	return result
}

//根据id生成一颗树
func buildTree(id int,root *Message) {
	var wg sync.WaitGroup
	DB.Where("Mid=?", id).Find(root)
		root.Kids = kidContent(id)
	for i,_:=range root.Kids{
		wg.Add(1)
		/*go*/ func() {//=================递归并发===好想吹一手
			//47.3048ms gorm不并发
			//64.1708ms gorm并发
			//32.9072ms 原生不并发
			//49.657ms 原生并发
			//database/sql/driver 级别的Conn,Stmt都是非并发安全的
			//database/sql 级别的DB,Stmt都是并发安全的，因为内部实现了连接池和锁
			defer wg.Done()
			buildTree(root.Kids[i].Mid, &root.Kids[i])
		}()
	}
	wg.Wait()
}
//生成一个...不知道怎么描述
func (m *Message)GetContent(r Message,reply *[]Message) error {
	DB.Where("Pid=?",new(int)).Find(reply)
	for i,_:=range *reply{
		buildTree((*reply)[i].Mid,&(*reply)[i])
	}
	return DB.Error
}