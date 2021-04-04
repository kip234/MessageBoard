> 前情提要：由于某次阶段作业内容与考核作业类似，且有不同要求。由于已经实现了部分代码，于是有部分更改
>
> 详情查看[更该部分](#更改部分)

# 内容


> 知乎留言板。(实现评论、回复、点赞、发布言论等功能，5月22日截止)。

## 目录

> - [总览](#总览)
>- [关系图解](#关系图解)
>   
>- [数据库设计](#数据库设计)
>   
>- [架构模式](#架构模式)
> - [模型](#模型)
> 	- [User](#User)
>	- [Message](#Message)
>- [配置](#配置)
> - [API列表](#API列表)
>	- [访问主页](#访问主页)
>	- [发布言论](#发布言论)
>	- [评论](#评论)
>	- [点赞](#点赞)
>	- [回复](#回复)
>	- [注册](#注册)
>	- [登录](#登录)

## 总览

### 关系图解

> 额…虽然图是这么画的，不过在写代码的时候 评论和回复 都没有对象的限制


```mermaid
graph TD

free2[行为]==>free1((对象))

a((言论))
b((评论))
c((回复))

发布言论==>a

评论==>a

点赞==>a
点赞==>b
点赞==>c

回复==>b
回复==>c
```

### 数据库设计

> 全搁在一个表里面

### 架构模式

> 额，好吧，我不知道属于哪类

```mermaid
graph TD
all(pkgs)
all==>pkg1(config)==>处理JSON配置
all==>pkg2(Database)==>1(初始化数据库<br>好像流行用dao)
all==>pkg3(Handler)==>处理器
all==>pkg4(Model)==>模型
all==>pkg5(Routers)==>路由设置
```



## 模型

> 经过多次尝试最终决定存在一个表里面
>
> 也许是我太菜了吧。

### User

```go
type User struct {
	Uid uint `gorm:"primaryKey",binding:"required"`
	Name string	`binding:"required"`//用户名
	Pwd string `binding:"required"`//用户密码
}
```

### Message

```go
type Message struct{
	Like uint//点赞数
	Content string `gorm:"string not null",binding:"required"`//内容
	Mid uint `gorm:"primaryKey",binding:"required`//自己的身份标识
	Pid uint//上一级ID
	Kids []Message `gorm:"-"`//子级
	Uid uint `binding:"required`//归属的用户
}
```

## 配置

```json
{
  "SqlName":"留言板",
  "SqlUserName": "root",
  "SqlUserPwd": "root",
  "SqlAddr":"127.0.0.1:3306",
  "Addr":"localhost:8080"
}
```

| KEY         | DESCRIPTION            |
| ----------- | ---------------------- |
| SqlName     | 数据库名               |
| SqlUserName | 数据库登录用用户名     |
| SqlUserPwd  | 数据库登录用用户名密码 |
| SqlAddr     | 数据库地址             |
| Addr        | 服务器地址             |



## API列表

### 访问主页

| 方法 | 路由 |
| :--: | :--: |
| GET  |  /   |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

### 发布言论

| 方法 |   路由   |
| :--: | :------: |
| POST | /publish |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

|   KEY   | VALUE  | DESCRIPTION |
| :-----: | :----: | :---------: |
| Content | string |  发言内容   |

> - - - - - - - - - - - - >

### 评论

| 方法 |   路由   |
| :--: | :------: |
| POST | /comment |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

|   KEY   | VALUE  | DESCRIPTION |
| :-----: | :----: | :---------: |
| Content | string | 评论的内容  |
|   Pid   |  uint  |    上级     |

> - - - - - - + - - - - - >

### 点赞

| 方法 | 路由  |
| :--: | :---: |
| POST | /like |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| Pid  | uint  |   上一级    |

> - - - - - - - - - - - - >

### 回复

| 方法 |  路由  |
| :--: | :----: |
| POST | /reply |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

|   KEY   | VALUE  | DESCRIPTION |
| :-----: | :----: | :---------: |
|   Pid   |  uint  |   上一级    |
| Content | string | 回复的内容  |

###注册

| 方法 |   路由    |
| :--: | :-------: |
| POST | /register |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

| KEY  | VALUE  | DESCRIPTION |
| :--: | :----: | :---------: |
| Pwd  | string |  用户密码   |
| Name | string |   用户名    |

### 登录

| 方法 |  路由  |
| :--: | :----: |
| POST | /login |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| nil  |  nil  |     nil     |

**form-data**

| KEY  | VALUE  | DESCRIPTION |
| :--: | :----: | :---------: |
|Uid|int|用户ID|
| Pwd  | string |  用户密码   |

## 更改部分

路径：

```
/Model/Message.go
/Model/User.go
/Middleware/*
/Handler/*
```

> 使用空接口来适配\*gorm.DB与\*sql.DB两种数据库

```
/main.go
```

> 声明了两种数据库供使用

```
/Database/InitSQL.go
```

> 使用原始接口

```
/Model/Message.go 128行
```

> 在使用gorm时可使用并发,如果使用原生接口则会由于链接过多而报错

吐槽一下：

> 原生接口的scan传入的参数的顺序应和 结构定义时的成员顺序一致/desc命令输出的一致