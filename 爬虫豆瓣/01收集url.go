package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"reptile_douban/models"
	"strings"
	"time"
)
//用途，从入口文件开始爬取页面连接，并存储到队列1和队列2；队列1用于获得电影信息的；
//获取url并存储到队列中
var urlList []string = make([]string,0)
func main()  {
	var url string = "https://movie.douban.com/subject/1292722/?from=subject-page"		//爬虫入口路径
	req := httplib.Get(url)
	str,_ := req.String()		//获取的html内容
	getALL_URL(str)		//传入html，得到全部url并存储到队列2和队列1
	//首次运行会存储10+个url

	for  {
		time.Sleep(time.Second * 5)		//等待2秒钟
		url01,_ := models.RpopURL02()	//每次从队列2取一个url
		fmt.Println("从队列2取出的url是：",url01)
		req := httplib.Get(url01)
		str,_ := req.String()		//获取的html内容
		if str == ""{
			fmt.Println("本次请求的ur：",url01,"内容为空，正在跳过")
			continue
		}
		getALL_URL(str)		//传入html，得到全部url并存储到队列2和队列1
	}
}

//传入一个html得到N个url,并存储到队列2和队列1中；
func getALL_URL(movieHTML string){

	relUrl := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*)/?from=subject-page" >`)
	resultUrl := relUrl.FindAllStringSubmatch(movieHTML,-1)		//得到url切片
	var urlList []string = make([]string,0)
	for _,v := range resultUrl{
		//fmt.Println(v[0])
		urlTmp := strings.Split(v[0],"\"")
		url := urlTmp[1]		//得到标准的url
		//fmt.Println(url)
		//fmt.Println("0 = ",v[0],"\t",strings.Split(v[0]," "))
		//fmt.Println("1 = ",v[1])
		//fmt.Println("2 = ",v[2])
		//err01 := models.LpushURL(url)		//每次得到url就存储到队列中
		//fmt.Println(err01)
		fmt.Println("当前的urlList:",urlList)
		urlList = append(urlList, url)		//将每次得到的url追加到一个切片中


	}
	models.LpushURL02(urlList)		//将url列表存储到队列2中，
	models.LpushURL(urlList)		//将url列表存储到队列1中，
	fmt.Println("存入的urlList:",urlList)

	//return urlList
}

