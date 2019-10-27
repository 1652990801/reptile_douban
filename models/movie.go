package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/Go-SQL-Driver/MySQL"
)

//表的名字是movie_info
type MovieInfo struct {
	Id     int 	`orm:"pk;auto;column(Id)"` //主键，自增id
	Name	string	//电影名字
	Score	string	//电影评分
	Director	string	//电影导演
	Language	string	//电影语言
	Mins	string	//电影片长
	Origin	string	//产地
	CommentCount	string	//电影评论数量
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:666666@tcp(127.0.0.1:3306)/reptile_douban?charset=utf8")
	orm.RegisterModel(new(MovieInfo))		//创建数据库
	orm.RunSyncdb("default",false,true)

}

func AddMonivInfo(movieInfo *MovieInfo)(result int64,err error) {
	db := orm.NewOrm()	//创建一个orm对象
	fmt.Println("要插入的电影信息是movieInfo:",movieInfo)
	num,err := db.Insert(movieInfo)		//插入数据，这里接受的是内存地址
	return num,err
}