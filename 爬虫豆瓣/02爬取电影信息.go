package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"reptile_douban/models"		//会执行init函数，创建数据库
	"strings"
	"time"
)

//主要用途是 从队列1中读取url并获得电影信息，存储写入redisvisit队列

func main() {

	for {

		url,_ := models.RpopURL()	//每次从队列1获取一个url

		isVisit := models.SismemberURL(url)		//判断url是否已经访问过，为真则跳过
		if isVisit {		//判断url是否访问过
			fmt.Println("该url访问过，本次跳过:",url)
			continue
		}
		req := httplib.Get(url)		//发起请求
		str,err := req.String()		//str则是html内容
		check(err)
		if str == ""{
			fmt.Println("请求的url为空，跳过")
			continue
		}
		movieInfo := getMovieInfo(str)	//获得电影信息
		num,_ := models.AddMonivInfo(&movieInfo)		//电影信息入库
		fmt.Println("电影信息已经入库：",num)
		err03 := models.SaddURL(url)			//该url已经使用过，可以作为记录了。
		check(err03)

		q1,q2,q3 := models.GetQueueNum()
		fmt.Println("队列1：",q1,"\t队列2：",q2,"\t队列3：",q3)
		time.Sleep(time.Second * 1)		//每一秒分析一部电影信息

	}

}
//传入html并返回电影信息结构体
func getMovieInfo(movieHTML string) models.MovieInfo  {
	var movieInfo models.MovieInfo		//声明这个类，接下来获取的信息，赋值到该类

	//导演
	regDirector := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	resultDirector := regDirector.FindAllStringSubmatch(movieHTML,-1)
	director := string(resultDirector[0][1])
	movieInfo.Director = director
	//fmt.Println(len(result[0]))
	//fmt.Println(result[0][0])
	//fmt.Println(result[0][1])
	//fmt.Println(result[0][1])		//表示导演
	//fmt.Println("导演director:",director)

	//电影名字
	regMovieName := regexp.MustCompile(`<span property="v:itemreviewed">(.*)</span>`)
	resultMovieName := regMovieName.FindAllStringSubmatch(movieHTML,-1)
	movieName := resultMovieName[0][1]
	//fmt.Println("电影名字movieName:",movieName)
	movieInfo.Name = movieName


	regScore := regexp.MustCompile(`<strong.*?property="v:average">(.*)</strong>`)
	resultGrade := regScore.FindAllStringSubmatch(movieHTML,-1)
	movieScore := resultGrade[0][1]
	//fmt.Println("电影评分movieScore:",movieScore)
	movieInfo.Score = movieScore

	regLanguage := regexp.MustCompile(`<span.*?>语言:</span>(.*)`)
	resultLanguage := regLanguage.FindAllStringSubmatch(movieHTML,-1)
	movieLanguage := strings.Trim(resultLanguage[0][1],"<br/>")
	//fmt.Println("电影语言movieLanguage:",strings.Trim(movieLanguage,"<br/>"))
	movieInfo.Language = strings.TrimSpace(movieLanguage)		//删除两侧空白

	regMins := regexp.MustCompile(`<span property="v:runtime" content=.*?>(.*)</span>`)
	resultMins := regMins.FindAllStringSubmatch(movieHTML,-1)
	movieMins := resultMins[0][1]
	//fmt.Println("电影时间movieMins:",movieMins)
	movieInfo.Mins = movieMins

	regOrigin := regexp.MustCompile(`<span class="pl">制片国家/地区:</span>(.*)`)
	resultOrigin := regOrigin.FindAllStringSubmatch(movieHTML,-1)
	movieOrigin := strings.Trim(resultOrigin[0][1],"<br/>")
	//fmt.Println("电影产地movieFabricate:",strings.Trim(movieFabricate,"<br/>"))
	movieInfo.Origin =  strings.TrimSpace(movieOrigin)

	//短评数量
	regEvaluate := regexp.MustCompile(`<a href="https://movie.douban.com/subject.*?">全部 (.*) 条</a>`)
	resultEvaluate := regEvaluate.FindAllStringSubmatch(movieHTML,-1)
	movieEvaluate := resultEvaluate[0][1]
	//fmt.Println("电影评价数量movieEvaluate:",movieEvaluate)
	movieInfo.CommentCount = movieEvaluate

	//电影类型
	//regType := regexp.MustCompile(`<span property="v:genre">(.*)</span> / .*?`)
	//resultType := regType.FindAllStringSubmatch(movieHTML,-1)
	//movieType := resultType[0][1]
	//fmt.Println("电影类型movieType:",movieType)


	return movieInfo
}























func check(err error)  {
	if err != nil{
		fmt.Println(err)
		return
	}
}
