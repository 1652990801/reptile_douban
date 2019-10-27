package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
const (

	URL_QUEUE = "url_queue"
	URL_QUEUE02 = "url_queue02"
	URL_VISIT_SET = "url_visit_set"
)
//队列2，存储每次获得的url
func LpushURL02(urlList []string)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	for _,v := range urlList {
		status,_ := conn.Do("LPUSH",URL_QUEUE02,v)
		fmt.Println("LpushURL02执行状态：",status)
	}
	lenth,_ := conn.Do("llen",URL_QUEUE02)
	fmt.Println("当前队列2数量：",lenth)
}
//队列2，读取队列2中的url,并获得N个url
func RpopURL02() (url string,err error)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	resultUrl,err01 := redis.String(conn.Do("RPOP",URL_QUEUE02))
	return resultUrl,err01
}



//存储url
func LpushURL(urlList []string)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	for _,v := range urlList{
		status,_ := conn.Do("LPUSH",URL_QUEUE,v)
		fmt.Println("LpushURL执行状态：",status)
	}
	lenth,_ := conn.Do("llen",URL_QUEUE)
	fmt.Println("当前队列1数量：",lenth)
}

//读取url
func RpopURL() (url string,err error)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	resultUrl,err01 := redis.String(conn.Do("RPOP",URL_QUEUE))
	return resultUrl,err01
}

//存储使用过的url
func SaddURL(url string)(error)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	status,err := conn.Do("SADD",URL_VISIT_SET,url)
	fmt.Println("saddurl执行状态：",status)

	num,_ := conn.Do("llen",URL_VISIT_SET)
	fmt.Println("当前url_visit_set数量：",num)
	return err
}

//判断url是否包含在其中
func SismemberURL(url string)bool  {
	var isVisit bool
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	status,_ := conn.Do("sismember",URL_VISIT_SET,url)	//存在status则为1，否则为2 int64类型
	//fmt.Println("sismemberURL执行状态：",status)
	//fmt.Println("sismemeberurl错误信息：",err)
	//fmt.Println("status:",status,"\t status_type:",reflect.TypeOf(status))
	//fmt.Println("err:",err,"\t err_type:",reflect.TypeOf(err))

	if status == int64(0) {		//如果存在则为true
		//fmt.Println("不存在")
		isVisit = false
	}else {
		//fmt.Println("存在")
		isVisit = true
	}
	return isVisit
}

func GetQueueNum() (interface{},interface{},interface{})  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()


	queue01,_ :=conn.Do("llen",URL_QUEUE)
	queue02,_ :=conn.Do("llen",URL_QUEUE02)
	queue03,_ :=conn.Do("llen",URL_VISIT_SET)

	return queue01,queue02,queue03



}