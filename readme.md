# 内容


> 知乎留言板。(实现评论、回复、点赞、发布言论等功能，5月22日截止)。

## 总览


```mermaid

```

### 参数列表

| 参数名 | 含义 |
| ------ | ---- |
|        |      |



## 数据结构


### 用户

```go
type User struct{
    ID uint		`json:`//用户唯一标识
    Name string	`json`//用户名
    Pwd string	`json`//用户密码
}
```



###发言

```go
type Speech struct{
    User			//发言的用户
    Like uint		//点赞数
    Content string	//内容
    Data string		//日期
    Sid uint		//该发言的ID
    Fid uint		//如果该发言是另一条发言的评论则有意义，否则小于0
}
```



## API 列表

### 评论

| 方法 |  路由   |
| :--: | :-----: |
| POT  | comment |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
| Fid  |       |             |

**form-data**

|   KEY   | VALUE | DESCRIPTION |
| :-----: | :---: | :---------: |
|         |       |             |
|         |       |             |
| Content |       |             |

> - - - - - - + - - - - - >

### 	点赞

| 方法 | 路由 |
| :--: | :--: |
| POT  | like |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |

**form-data**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |

> - - - - - - - - - - - - >

### 回复

| 方法 | 路由  |
| :--: | :---: |
| POT  | reply |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |

**form-data**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |

>- - - - - - - - - - - - >

### 发布言论

| 方法 |  路由   |
| :--: | :-----: |
| POT  | publish |

**Query Params**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |

**form-data**

| KEY  | VALUE | DESCRIPTION |
| :--: | :---: | :---------: |
|      |       |             |