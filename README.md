![image-20221006194319209](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061943411.png)

Gorm是Go语言的一个orm框架，类似Java中的JPA的实现（Hibernate、EclipseLink等）。

# 本文目的

本文就是按照官网官方说明，自己动手尝试下，增加记忆，仅此而已。

# 安装

![image-20221006194651682](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061946829.png)

进入项目目录安装gorm

```shell
➜  Go-Gorm-Notes git:(main) ✗ go get -u gorm.io/gorm
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added gorm.io/gorm v1.23.10
```

我用的数据库是mysql，所以还需要引入Mysql驱动

```shell
➜  Go-Gorm-Notes git:(main) ✗ go get -u gorm.io/driver/mysql
go: added github.com/go-sql-driver/mysql v1.6.0
go: added gorm.io/driver/mysql v1.3.6
```

# 创建连接

下面代码是创建一个Mysql的gorm.DB，之后的操作都要使用这个DB。

```
package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db = initDb()

//初始化数据库连接
func initDb() *gorm.DB {
	// Mysql的链接字符串,我电脑上用户名是root，密码是qwe!@#123
	dsn := "root:qwe!@#123@tcp(127.0.0.1:3306)/go-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// Mysql驱动有自己的配置选项，可以通过 mysql.New(mysql.Config{})配置。具体可看mysql.Config。
	Db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return Db
}
```

# Model

对于这种ORM框架，Model是特别重要的，也是值得去深入学习的地方，因为一个小问题，可能就会引发数据库的问题。

# 初识

我会创建一个简单的Model，并使用gorm的自动创建数据库表功能，来看看如何做到通过Model自动创建表。

创建一个Author结构体

```go
package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string
}
```

main方法中使用来自动创建数据库表

```go
package main

import "Go-Gorm-Notes/models"

func main() {
	// 使用AutoMigrate自动创建数据库表
	models.Db.AutoMigrate(&models.Author{})
}
```

运行后会发现数据库中多了一个叫做authors的表。请注意这个名字，这是gorm自动创建的，能否自己指定呢？肯定是可以的，看下自动创建的表的结构如下：

```sql
create table authors
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    name       longtext    null
);

create index idx_authors_deleted_at
    on authors (deleted_at);

```

说明：id，created_at，updated_at，deleted_at都是因为Author结构体继承了gorm.Model。这是gorm自身提供的，我们也可以不使用它，如果不用就要自己定义主键字段。

另外name字段是自己加的，但是name字段的类型是longtext，这可能不是我们想要的，我们可能想要的是name是varchar(200)这样的类型，也是可以自定义的。官网都有说明，接下来我一一尝试下，并记录下来。

# 自定义表名

自定义表名需要实现gorm.schema下的TableName方法。

```go
package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string
}

// TableName 自定义表名
func (Author) TableName() string  {
	return "author"
}
```

重新运行看看是否会自动创建author表

![image-20221007102438864](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071024108.png)

# 模型定义

## 字段标签

可以在结构体字段名后面使用`gorm:xxx`的机构来配置标签，从而达到自定义数据库列信息的效果。

声明 model 时，tag 是可选的，GORM 支持以下 tag： tag 名大小写不敏感，但建议使用 `camelCase` 风格

| 标签名                 | 说明                                                         |
| :--------------------- | :----------------------------------------------------------- |
| column                 | 指定 db 列名                                                 |
| type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：`not null`、`size`, `autoIncrement`… 像 `varbinary(8)` 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：`MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT` |
| serializer             | 指定将数据序列化或反序列化到数据库中的序列化器, 例如: `serializer:json/gob/unixtime` |
| size                   | 定义列数据类型的大小或长度，例如 `size: 256`                 |
| primaryKey             | 将列定义为主键                                               |
| unique                 | 将列定义为唯一键                                             |
| default                | 定义列的默认值                                               |
| precision              | specifies column precision                                   |
| scale                  | specifies column scale                                       |
| not null               | specifies column as NOT NULL                                 |
| autoIncrement          | specifies column auto incrementable                          |
| autoIncrementIncrement | auto increment step, controls the interval between successive column values |
| embedded               | embed the field                                              |
| embeddedPrefix         | column name prefix for embedded fields                       |
| autoCreateTime         | track current time when creating, for `int` fields, it will track unix seconds, use value `nano`/`milli` to track unix nano/milli seconds, e.g: `autoCreateTime:nano` |
| autoUpdateTime         | track current time when creating/updating, for `int` fields, it will track unix seconds, use value `nano`/`milli` to track unix nano/milli seconds, e.g: `autoUpdateTime:milli` |
| index                  | create index with options, use same name for multiple fields creates composite indexes, refer [Indexes](https://gorm.io/zh_CN/docs/indexes.html) for details |
| uniqueIndex            | same as `index`, but create uniqued index                    |
| check                  | creates check constraint, eg: `check:age > 13`, refer [Constraints](https://gorm.io/zh_CN/docs/constraints.html) |
| <-                     | set field’s write permission, `<-:create` create-only field, `<-:update` update-only field, `<-:false` no write permission, `<-` create and update permission |
| ->                     | set field’s read permission, `->:false` no read permission   |
| -                      | ignore this field, `-` no read/write permission, `-:migration` no migrate permission, `-:all` no read/write/migrate permission |
| comment                | add comment for field when migration                         |

### column

定义列的名字，比如如下代码将Name字段对应的数据库列名为t_name。

```go
Name string `gorm:"column:t_name"`
```

### type

列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：`not null`、`size`, `autoIncrement`… 像 `varbinary(8)` 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：`MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT`

比如我继续设置Name字段为varchar(200) not null default ''

```go
Name string `gorm:"column:name;type:varchar(200) not null default ''"`
```

![image-20221007105105657](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071051801.png)

可以看到类型等信息已经正确设置。

### size

定义列数据类型的大小或长度，例如 `size: 256`

```go
Sex  string `gorm:"size:10"`
```

这里我没有使用name字段，是因为Name字段制定了type，类型varchar(200)，再指定size无效。

### primarykey

上面的结构体我使用了gorm.Model，gorm自动给我生成了主键，现在我要自己定义主键，重新定义下结构体

```go
package models

type Author struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"column:name;type:varchar(200) not null default '';"`
	Sex  string `gorm:"size:10"`
}

// TableName 自定义表名
func (Author) TableName() string {
	return "author"
}
```

我使用

```go
ID   uint   `gorm:"primarykey"`
```

定义ID是主键。

![image-20221007111644223](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071116337.png)

### unique

使用unique就会设置唯一索引。

```go
Name string `gorm:"column:name;type:varchar(200) not null default '';unique"`
```

结果如下：

![image-20221007111945164](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071119339.png)

### default

设置默认值

```go
Sex  string `gorm:"size:10;default:'男'"`
```

![image-20221007112257219](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071122338.png)

### index

```go
Name string `gorm:"column:name;type:varchar(200) not null default '';index"`
```

![image-20221007113923909](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071139052.png)

创建了一个普通索引。

### uniqueIndex

创建唯一索引

```go
// 身份信息唯一
Identify string `gorm:"size:100;uniqueIndex"`
```

需要注意的是上面的size:100;不能去掉，使用它可以是数据库类型变为varchar，否则类型是longtext，这个类似不能加唯一索引的。会报错。

![image-20221007114355511](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071143628.png)

### not null

指定列不能为空

```go
Identify string `gorm:"size:100;uniqueIndex;not null;"`
```

![image-20221007114808725](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071148855.png)

### autoCreateTime

创建记录时自动填充时间，取值nano或者milli。

![image-20221007121408205](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071214319.png)

插入一条记录

```go
models.Db.Save(&models.Author{Name: "张飞", Identify: "001"})
```

结果是：

| id   | name | sex  | identify | ct\_time\_nano          | ct\_time\_milli         | ct\_nano            | ct\_milli     |
| :--- | :--- | :--- | :------- | :---------------------- | :---------------------- | :------------------ | :------------ |
| 1    | 张飞 | 男   | 001      | 2022-10-07 12:11:47.854 | 2022-10-07 12:11:47.854 | 1665115907854000000 | 1665115907854 |

### autoUpdateTime

跟autoCreateTime类似。

### embedded

内嵌字段，将一个结构体嵌入进来

```go
package models

import "time"

type Author struct {
   ID   uint   `gorm:"primarykey"`
   Name string `gorm:"column:name;type:varchar(200) not null default '';index"`
   Sex  string `gorm:"size:10;default:'男'"`
   // 身份信息唯一
   Identify    string    `gorm:"size:100;uniqueIndex;not null;"`
   CtTimeNano  time.Time `gorm:"autoCreateTime:nano"` //nano/milli
   CtTimeMilli time.Time `gorm:"autoCreateTime:milli"`
   CtNano      int       `gorm:"autoCreateTime:nano"`
   CtMilli     int       `gorm:"autoCreateTime:milli"`
   ConcatWay   `gorm:"embedded"`
}

// ConcatWay 联系方式
type ConcatWay struct {
   Address string
   Phone   string
   Email   string
}

// TableName 自定义表名
func (Author) TableName() string {
   return "author"
}
```

运行AutoMigrate后，数据库表结构如下

![image-20221007122018677](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071220792.png)

增加了Address结构体下的字段。

### embeddedPrefix

对于内嵌的结构体字段，增加前缀，默认是空

```go
ConcatWay   `gorm:"embedded;embeddedPrefix:cw_"`
```

![image-20221007122758254](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071227397.png)

### comment

字段的备注信息

```go
ID   uint   `gorm:"primarykey;comment:主键ID"`
```

![image-20221007122846298](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071228411.png)

### serializer

指定将数据序列化或反序列化到数据库中的序列化器, 例如: `serializer:json/gob/unixtime`

```go
ByteJson      []byte    `gorm:"serializer:json"`
ConcatWayJSON ConcatWay `gorm:"serializer:json"`
ConcatWayGob  ConcatWay `gorm:"serializer:gob"`
TimeUnixtime  int64     `gorm:"serializer:unixtime;type:time"` //将int64的内容转化为ddatetime存储
```

![image-20221007123649845](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071236095.png)

存储一条数据

```go
models.Db.Create(&models.Author{
   Name:          "张飞",
   Identify:      "001",
   ByteJson:      []byte("字节切片"),
   ConcatWayJSON: models.ConcatWay{Address: "https://itlab1024.com", Email: "itlab1024@163.com"},
   ConcatWayGob:  models.ConcatWay{Address: "https://itlab1024.com", Email: "itlab1024@163.com"},
   TimeUnixtime:  12,
})
```

| id   | name | sex  | identify | ct\_time\_nano          | ct\_time\_milli         | ct\_nano            | ct\_milli     | cw\_address | cw\_phone | cw\_email | byte\_json         | concat\_way\_json                                            | concat\_way\_gob                                             | time\_unixtime      |
| :--- | :--- | :--- | :------- | :---------------------- | :---------------------- | :------------------ | :------------ | :---------- | :-------- | :-------- | :----------------- | :----------------------------------------------------------- | :----------------------------------------------------------- | :------------------ |
| 1    | 张飞 | 男   | 001      | 2022-10-07 14:27:01.859 | 2022-10-07 14:27:01.859 | 1665124021859000000 | 1665124021859 |             |           |           | "5a2X6IqC5YiH54mH" | {"Address":"https://itlab1024.com","Phone":"","Email":"itlab1024@163.com"} | 0x37FF8103010109436F6E63617457617901FF82000103010741646472657373010C00010550686F6E65010C000105456D61696C010C0000002DFF82011568747470733A2F2F69746C6162313032342E636F6D021169746C616231303234403136332E636F6D00 | 1970-01-01 08:00:12 |

### autoIncrement 

设置列自增，需要与type标签联合使用

### autoIncrementIncrement

设置列自增步长，需要与type标签联合使用

### unique

设置唯一键

### check

设置约束，比如设置name的值不能等于abc

```go
Name  string `gorm:"check:name <> 'abc'"`
```

# 关联

## Belongs To

`belongs to` 会与另一个模型建立了一对一的连接。 这种模型的每一个实例都“属于”另一个模型的一个实例。就是谁属于谁，比如一个人属于一个国家。

看如下如下两个结构体User和Country

```go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	CountryId uint
	// 这里不能使用匿名
	Country   Country
}
type Country struct {
	*gorm.Model
	Name string
}
```

自动创建表后，可以得到如下结构

![image-20221007155420333](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071554609.png)

用户表的country_id和country表的id关联了起来。一对一的关系。

## Has One

`has one` 与另一个模型建立一对一的关联，但它和一对一关系有些许不同。 这种关联表明一个模型的每个实例都包含或拥有另一个模型的一个实例。

例如，您的应用包含 user 和 credit card 模型，且每个 user 只能有一张 credit card。

```go
// User 有一张 CreditCard，UserID 是外键
type User struct {
   gorm.Model
   CreditCard CreditCard
}

type CreditCard struct {
   gorm.Model
   Number string
   UserID uint
}
```



![image-20221007155958693](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071559841.png)

可以看到users表没有credit的相关字段。

再看下credit_cards表

![image-20221007160034368](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071600525.png)

该表的userId关联到了user表的ID字段。

## Has Many

类似于has one，只不过这里要使用的是切片

```go
// User 有多张 CreditCard，UserID 是外键
type User struct {
   gorm.Model
   CreditCards []CreditCard
}

type CreditCard struct {
   gorm.Model
   Number string
   UserID uint
}
```

创建出来的表如下：

users表:

![image-20221007160358864](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071603000.png)

credit_cards表如下：

![image-20221007160504811](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071605944.png)

跟has one创建的表结构是一样的。

## Many To Many

Many to Many 会在两个 model 中添加一张连接表，可以通过标签many2manay设置关联表的名字。

```go
// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
   gorm.Model
   Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
   gorm.Model
   Name string
}
```

会创建user_languages中间表，并且关联表中的user_id跟User表的ID关联，language_id跟languages表的id关联。

![image-20221007164121248](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210071641394.png)

上面的都是使用默认的情况，比如外键名称等，如果想更换名字等信息，就得重写外键，这里我就一一说明了。
# 原生SQL
待更新...
