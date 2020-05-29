package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/url"
	"strconv"
	"strings"
	"time"
	"untitled/src/goget"
	_ "untitled/src/goget"
)
type lotType struct {
	//openUrl string
	mysqlTable string
	lotKey string //
	nowDraw float64 //当前已存期数redis
	nextTime int64
}
var webs map[string]string
var firstRest bool
var db *sql.DB
var lotList map[string]*lotType

func main() {
	//var er error
	//db,er = sql.Open("mysql", "root:123456@tcp(192.168.0.193:3306)/test?charset=utf8")
	//db,er = sql.Open("mysql", "root:OmLgjkxS9l(6Ck@/test?charset=utf8")
	//checkErr(er)
	firstRest = true
	webs = make(map[string]string)
	lotList = make(map[string]*lotType)
	println("111111111111")
	//initLot()

	println("3333333333")

	//lotList["ssc"] = &lotType{"cp_ticket_award_record_1","ssc",0,1504111241}
	//lotList["jsk3"] = &lotType{"cp_ticket_award_record_2","jsk3",0,1504111241}
	//lotList["ahk3"] = &lotType{"cp_ticket_award_record_3","ahk3",0,1504111241}
	//lotList["bjk3"] = &lotType{"cp_ticket_award_record_4","bjk3",0,1504111241}
	//lotList["gdsyxw"] = &lotType{"cp_ticket_award_record_5","gdsyxw",0,1504111241}
	//lotList["jxsyxw"] = &lotType{"cp_ticket_award_record_6","jxsyxw",0,1504111241}
	//lotList["bjpk10"] = &lotType{"cp_ticket_award_record_7","bjpk10",0,1504111241}
	//lotList["xyft"] = &lotType{"cp_ticket_award_record_8","xyft",0,1504111241}
	//lotList["fsd"] = &lotType{"cp_ticket_award_record_fc3d","fsd",0,1504111241}

	webs["time"] = "1590481936"
	webs["passwd"] = "640e6ae0cb53c7ca3c6df174d3db8a22"

	//openMysql()
	if firstRest {
		//resetwebs()
		firstRest = false
	}
	fmt.Println("爬虫开始")
	//pachong()

	//go task()
	//go getdata()
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




func initLot()  {
	//查询数据
	rows, err := db.Query("SELECT tables_name,name_en FROM cp_ticket")
	checkErr(err)
	if err!=nil {
		return
	}
	fmt.Println(rows,"rows============================")
	for rows.Next() {
		var table_name string
		var name_en string
		err = rows.Scan(&table_name, &name_en,)
		checkErr(err)
		lotList[name_en] = &lotType{table_name,name_en,0,1504111241}

		//fmt.Println(ticket_period)
		//fmt.Println(award_nums)
		//fmt.Println(float)
	}
	fmt.Println("lotlist是：：::::::::::::::::::::::::::::::::::::::::::::",lotList)
	println("2222222222222222222222")
}




func resetwebs(){
	//ch := make(chan map[string]string)
	//go parseUrls("https://www.caipiaoapi.com/hall/hallk3/index/jsk3", ch)
	//webs = <-ch
	println("-----------------------------------------修改webs值-------------------------")
	parseUrl("https://www.caipiaoapi.com/hall/hallk3/index/jsk3")
}

func sddSql(mm map[string]interface{},ssc *lotType)  {
	if data,ok:=mm["data"].(map[string]interface{});ok {
		preDrawIssue,err := strconv.ParseFloat(data["preDrawIssue"].(string),64)
		fmt.Println(preDrawIssue,"float是多少last_no",data["preDrawIssue"],ssc.nowDraw)
		if int(ssc.nowDraw)==0 {
			//查询数据
			rows, err := db.Query("SELECT ticket_period ,award_nums,next_award_time FROM "+ssc.mysqlTable+" ORDER BY ticket_period DESC LIMIT 0,1")
			checkErr(err)
			for rows.Next() {
				var ticket_period string
				var award_nums string
				var next_award_time string
				err = rows.Scan(&ticket_period, &award_nums,&next_award_time)
				ticket,err := strconv.ParseFloat(ticket_period,64)
				ssc.nowDraw = ticket
				loc, _ := time.LoadLocation("Local")
				the_time, err := time.ParseInLocation("2006-01-02 15:04:05", next_award_time, loc)
				if err == nil {
					ssc.nextTime = the_time.Unix() //1504082441
					fmt.Println("重置下次获得时间:",data["drawTime"],ssc.nextTime)
				}
				checkErr(err)
			}
		}
		if float64(preDrawIssue)<=ssc.nowDraw {
			return
		}
		fmt.Println("ssc.nowDraw",ssc.nowDraw,"period",)
		ssc.nowDraw,_ = strconv.ParseFloat(data["preDrawIssue"].(string),64)
		fmt.Println("ssc.nowDraw",ssc.nowDraw,"period",data["preDrawIssue"])

		//插入数据
		stmt, err := db.Prepare("INSERT INTO "+ssc.mysqlTable+" SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
		checkErr(err)

		res, err := stmt.Exec(1, data["preDrawIssue"], data["preDrawCode"], data["preDrawTime"],data["drawIssue"],data["drawTime"])
		checkErr(err)
		if err==nil {
			loc, _ := time.LoadLocation("Local")
			the_time, err := time.ParseInLocation("2006-01-02 15:04:05", data["drawTime"].(string), loc)
			if err == nil {
				ssc.nextTime = the_time.Unix() //1504082441
				fmt.Println("重置下次获得时间:",data["drawTime"],ssc.nextTime)
			}
			//ssc.nextTime
		}


		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(ssc.mysqlTable,"插入数据======================================id:",id)
	}
}

func doasync(s *lotType) {
	mm:= doget(*s)
	if mm==nil { //避免过期过次重置
		if !firstRest {
			firstRest = true
			resetwebs()
			fmt.Println(webs)
		}
	}else{
		sddSql(mm,s)
	}
}

func pachong() {

	//ssc := lotType{"cp_ticket_award_record_1","ssc",0}
	//mm := doget(ssc)
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * *", func() {
		timestamp := time.Now().Unix()
		fmt.Println("进入task，不知道进没进循环:",time.Now())
		for key,s := range lotList {
			if s.nextTime<timestamp-10 {
				fmt.Println(time.Now(),"开始",key,time.Unix(s.nextTime, 0).Format("02/01/2006 15:04:05 PM"))
				go doasync(s)
				fmt.Println(time.Now(),"结束",key)
			}else{
				fmt.Println("进入task，没有进入循环:",time.Unix(s.nextTime, 0).Format("02/01/2006 15:04:05 PM"))
			}
		}
	})
	c.Start()





	//start := time.Now()
	//fmt.Println(time.Now(),"开始时间")
	//ch := make(chan map[string]string)
	//go parseUrls("https://www.caipiaoapi.com/hall/hallk3/index/jsk3", ch)
	//webs = <-ch
	//fmt.Println(time.Now(),"结束时间")
	//fmt.Println(webs,"参数获得了")
	//elapsed := time.Since(start)
	//fmt.Printf("Took %s", elapsed)
}

func doget(lottly lotType )(map[string]interface{}) {

	//params := url.Values{}
	//
	////Url, err := url.Parse("https://www.caipiaoapi.com/hall/hallajax/getLotteryList")
	//Url, err := url.Parse("https://www.caipiaoapi.com/hall/hallajax/getLotteryInfo")
	//if err != nil {
	//	panic(err.Error())
	//
	//}
	//params.Set("lotKey", lottly.lotKey)
	//params.Set("count", string("5"))
	//params.Set("time", webs["time"])
	//params.Set("passwd", webs["passwd"])
	//
	//////如果参数中有中文参数,这个方法会进行URLEncode
	//Url.RawQuery = params.Encode()
	//urlPath := Url.String()
	//println(urlPath)
	//resp, err := http.Get(urlPath)
	//
	//defer resp.Body.Close()
	//s, errs := ioutil.ReadAll(resp.Body)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.caipiaoapi.com/hall/hallajax/getLotteryInfo", nil)
	if err!=nil {
		return nil
	}
	q:=req.URL.Query()
	q.Add("lotKey", lottly.lotKey)
	q.Add("count", string("5"))
	q.Add("time", webs["time"])
	q.Add("passwd", webs["passwd"])
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Set("Referer", "https://www.caipiaoapi.com/")
	req.Header.Set("Host", "www.caipiaoapi.com")
	resp, err := client.Do(req)
	if err!=nil {
		return nil
	}
	defer resp.Body.Close()
	s, errs := ioutil.ReadAll(resp.Body)
	if errs!=nil {
		return nil
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(s)), &dat); err == nil&&errs == nil&&dat["result"]!=nil {
	fmt.Println("姓名",lottly.lotKey,dat["result"])
		result,ok := dat["result"].(map[string]interface{})
		if ok {
			return result
		}
		return nil
	}else{
		println(err,"页面获取失败，准备重新抓取")
		fmt.Println("页面获取失败================",err,errs)
		return nil

	}
	return nil
	//println(string(s))
	//println(err,"errrrrrrrrrrf")
}

func parseUrl(url string) {
	doc := fetch(url)
	fmt.Println("进来了进来了",doc)

	if doc==nil {
		firstRest = false
		return
	}
	pswValue ,_ :=doc.Find("input#ajax_passwd").Eq(0).Attr("value")
	timeValue ,_ :=doc.Find("input#ajax_time").Eq(0).Attr("value")
	webs["time"] = timeValue
	webs["passwd"] = pswValue
	firstRest = false

	fmt.Println("设置webs成功",webs)
}

func parseUrls(url string, ch chan map[string]string) {
	doc := fetch(url)
	fmt.Println("docc-------------------------------------------------",doc)
	//doc.Find("ol.grid_view li").Find(".hd").Each(func(index int, ele *goquery.Selection) {
	//	movieUrl, _ := ele.Find("a").Attr("href")
	//	fmt.Println(strings.Split(movieUrl, "/")[4], ele.Find(".title").Eq(0).Text())
	//})
	pswValue ,_ :=doc.Find("input#ajax_passwd").Eq(0).Attr("value")
	timeValue ,_ :=doc.Find("input#ajax_time").Eq(0).Attr("value")
	fmt.Println("时间-----------------------------------------------", pswValue)
	fmt.Println("密码-----------------------------------------------", timeValue)
	//ssc := lotType{"cp_ticket_award_record_1","ssc",0}

	//doget(ssc)
	//
	//resp, err := http.Get("https://www.caipiaoapi.com/hall/hallajax/getLotteryList?")
	//if err != nil {
	//	panic(err)
	//
	//}




	//time.Sleep(2 * time.Second)
	//ch <- true
	//ch <- timeValue,pswValue
	ch <- map[string]string{"time":timeValue,"passwd":pswValue}

}

func fetch(url string) *goquery.Document {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Set("Referer", "https://www.caipiaoapi.com/")
	req.Header.Set("Host", "www.caipiaoapi.com")
	resp, err := client.Do(req)
	if err != nil {
		//log.Fatal("Http get err:", err)
		println("Http get err:", err)
		return nil
	}
	if resp.StatusCode != 200 {
		return nil
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil
		log.Fatal(err)
	}
	return doc
}


func getdata()  {
	newPeriod := map[int]float64 {1:0,6:0,16:0,25:0,39:0}
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.193:3306)/test?charset=utf8")
	checkErr(err)
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(newPeriod)
		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://www.caim8.com/openApi/for_fresh", strings.NewReader("name=cjb"))
		if err != nil {
			// handle error
		}
		req.Header.Set("referer", "https://www.caim8.com/syxwjx/index.html")
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(string(body)), &dat); err == nil {
			fmt.Println("==============json str 转map=======================")
			fmt.Println(dat["data"])
			if dat["data"]==nil {

			}
			data,ok := dat["data"].([]interface{})
			fmt.Println(ok,len(data))
			checkErr(err)
			if ok&& len(data)>0 {
				for _,item := range data {
					if ite,oks := item.(map[string]interface{});oks {
						if st,_:= ite["last_open"].(string); len(st)<1 {
							println("没有开奖号码")
							fmt.Println(ite["czid"],"没有开奖号码",ite["title"],ite)
							continue
						}
						fl,_ := ite["czid"].(float64)
						mt := int(fl)
						switch mt {
						case 1: //重庆时时彩
							float,err := strconv.ParseFloat(ite["last_no"].(string),64)
							fmt.Println(float,"float是多少last_no",ite["last_no"].(string))
							if newPeriod[1]==0 {
								//查询数据
								rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_1 ORDER BY ticket_period DESC LIMIT 0,1")
								checkErr(err)
								fmt.Println(rows,"rows============================")
								for rows.Next() {
									var ticket_period string
									var award_nums string
									err = rows.Scan(&ticket_period, &award_nums,)
									ticket,err := strconv.ParseFloat(ticket_period,64)
									newPeriod[1] = ticket
									checkErr(err)
									//fmt.Println(ticket_period)
									//fmt.Println(award_nums)
									//fmt.Println(float)
								}
							}
							if float64(float)<=newPeriod[1] {
								continue
							}
							newPeriod[1],_ = strconv.ParseFloat(ite["last_no"].(string),64)

							//插入数据
							stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_1 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
							checkErr(err)

							res, err := stmt.Exec(8, ite["last_no"], ite["last_open"], ite["last_day_time"],ite["next_no"],ite["next_day_time"])
							checkErr(err)

							id, err := res.LastInsertId()
							checkErr(err)
							println("++++++++++++++++++++++++++++++++++")
							fmt.Println("id=:",id)
							fmt.Println("插入数据======================================:")
							fmt.Println(ite)
						case 6: //广东11选5
							float,err := strconv.ParseFloat(ite["last_no"].(string),64)
							fmt.Println(float,"float是多少last_no",ite["last_no"].(string))
							if newPeriod[6]==0 {
								//查询数据
								rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_5 ORDER BY ticket_period DESC LIMIT 0,1")
								checkErr(err)
								fmt.Println(rows,"rows============================")
								for rows.Next() {
									var ticket_period string
									var award_nums string
									err = rows.Scan(&ticket_period, &award_nums,)
									ticket,err := strconv.ParseFloat(ticket_period,64)
									newPeriod[6] = ticket
									checkErr(err)
									//fmt.Println(ticket_period)
									//fmt.Println(award_nums)
									//fmt.Println(float)
								}
							}
							if float64(float)<=newPeriod[6] {
								continue
							}
							newPeriod[6],_ = strconv.ParseFloat(ite["last_no"].(string),64)
							//插入数据
							stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_5 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
							checkErr(err)

							res, err := stmt.Exec(8, ite["last_no"], ite["last_open"], ite["last_day_time"],ite["next_no"],ite["next_day_time"])
							checkErr(err)

							_, err = res.LastInsertId()
							checkErr(err)
							fmt.Println("插入数据======================================:")
							fmt.Println(ite)
						case 16: //江西11x5
							float,err := strconv.ParseFloat(ite["last_no"].(string),64)
							fmt.Println(float,"float是多少last_no",ite["last_no"].(string))
							if newPeriod[16]==0 {
								//查询数据
								rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_6 ORDER BY ticket_period DESC LIMIT 0,1")
								checkErr(err)
								fmt.Println(rows,"rows============================")
								for rows.Next() {
									var ticket_period string
									var award_nums string
									err = rows.Scan(&ticket_period, &award_nums,)
									ticket,err := strconv.ParseFloat(ticket_period,64)
									newPeriod[16] = ticket
									checkErr(err)
									//fmt.Println(ticket_period)
									//fmt.Println(award_nums)
									//fmt.Println(float)
								}
							}
							if float64(float)<=newPeriod[16] {
								continue
							}
							newPeriod[16],_ = strconv.ParseFloat(ite["last_no"].(string),64)
							//插入数据
							stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_6 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
							checkErr(err)

							res, err := stmt.Exec(8, ite["last_no"], ite["last_open"], ite["last_day_time"],ite["next_no"],ite["next_day_time"])
							checkErr(err)

							_, err = res.LastInsertId()
							checkErr(err)
							fmt.Println("插入数据======================================:")
							fmt.Println(ite)
						case 25: //江苏k3
							float,err := strconv.ParseFloat(ite["last_no"].(string),64)
							fmt.Println(float,"float是多少last_no",ite["last_no"].(string))
							if newPeriod[25]==0 {
								//查询数据
								rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_2 ORDER BY ticket_period DESC LIMIT 0,1")
								checkErr(err)
								fmt.Println(rows,"rows============================")
								for rows.Next() {
									var ticket_period string
									var award_nums string
									err = rows.Scan(&ticket_period, &award_nums,)
									ticket,err := strconv.ParseFloat(ticket_period,64)
									newPeriod[25] = ticket
									checkErr(err)
									//fmt.Println(ticket_period)
									//fmt.Println(award_nums)
									//fmt.Println(float)
								}
							}
							if float64(float)<=newPeriod[25] {
								continue
							}
							newPeriod[25],_ = strconv.ParseFloat(ite["last_no"].(string),64)
							//插入数据
							stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_2 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
							checkErr(err)

							res, err := stmt.Exec(8, ite["last_no"], ite["last_open"], ite["last_day_time"],ite["next_no"],ite["next_day_time"])
							checkErr(err)

							_, err = res.LastInsertId()
							checkErr(err)
							fmt.Println("插入数据======================================:")
							fmt.Println(ite)
						case 39: //安徽k3
							float,err := strconv.ParseFloat(ite["last_no"].(string),64)
							fmt.Println(float,"float是多少last_no",ite["last_no"].(string))
							if newPeriod[39]==0 {
								//查询数据
								rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_3 ORDER BY ticket_period DESC LIMIT 0,1")
								checkErr(err)
								fmt.Println(rows,"rows============================")
								for rows.Next() {
									var ticket_period string
									var award_nums string
									err = rows.Scan(&ticket_period, &award_nums,)
									ticket,err := strconv.ParseFloat(ticket_period,64)
									newPeriod[39] = ticket
									checkErr(err)
									//fmt.Println(ticket_period)
									//fmt.Println(award_nums)
									//fmt.Println(float)
								}
							}
							if float64(float)<=newPeriod[39] {
								continue
							}
							newPeriod[39],_ = strconv.ParseFloat(ite["last_no"].(string),64)
							//插入数据
							stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_3 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?,next_award_time=?")
							checkErr(err)

							res, err := stmt.Exec(8, ite["last_no"], ite["last_open"], ite["last_day_time"],ite["next_no"],ite["next_day_time"])
							checkErr(err)

							_, err = res.LastInsertId()
							checkErr(err)
							fmt.Println("插入数据======================================:")
							fmt.Println(ite)
						default:
							continue
						}
					}
				}
			}

			//db.Close()

		}
		fmt.Println(err)

	})
	c.Start()





	//fmt.Println(string(body))
}

type Profile struct {
	Name    string `json:"name"` //key转换成小写
	//Hobbies []string
	Hobbies  []interface{} `json:"hobbies"`
}



type Res struct {
	//string
	Code string
	Message string
	Obj interface{}
	//body interface{}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("Content-Type:",r.Header.Get("Content-Type"))  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	/*

	fmt.Println("key:",r.FormValue("key"))
	或者，效果一样 。方式 x-www-form-urlencoded 方式
		// 根据请求body创建一个json解析器实例
		decoder := json.NewDecoder(r.Body)

		// 用于存放参数key=value数据
		var params map[string]string

		// 解析参数 存入map
		decoder.Decode(&params)
		fmt.Println(params)
	*/
	r.ParseMultipartForm(0)
	if r.MultipartForm != nil {
		fmt.Println(r.MultipartForm.Value)
	}

	fmt.Println("==========================")

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("==========================")

	w.Header().Set("Server","A Go Web Server")
	w.WriteHeader(200)

	fmt.Println("==========================")


	//fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	//var bdss [] interface
	
	//obj := map[string]interface{}{"name":"ck","sex":"Man","age":18}
	slice1 := [] interface{}{"不知道SB",34,"sdf"}
	//profile := Profile{"Alex", []interface{}{"snowboarding", "programming"}}
	pro := Profile{"Alex", slice1}
	//js, _ := json.Marshal(pro)
	//js, err := json.Marshal(pro)
	js, err := json.Marshal(pro)

	res := Res{"200","成功","8"}

	fmt.Println("==========================",res)
	//js, err := json.Marshal(res)
	//js, err := json.Marshal(obj)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(string(js)))
	w.Write(js)
	//st := Map{}
}


func openMysql()  {
	//db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	//db, err := sql.Open("mysql", "root:OmLgjkxS9l(6Ck@tcp(47.98.141.227:3306)/test?charset=utf8")
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.193:3306)/test?charset=utf8")
	checkErr(err)

	//插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	name := "userinfo_1"
	query :="INSERT INTO "+name+" SET username=?,department=?,created=?"
	stmt, err := db.Prepare(query)
	checkErr(err)

	res, err := stmt.Exec("新增", "傻逼部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=:",id)
	//更新数据
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("astaxieupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println("affect",affect)

	//查询数据
	rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record ORDER BY ticket_period DESC LIMIT 0,1")
	checkErr(err)

	for rows.Next() {
		var ticket_period string
		var award_nums string
		err = rows.Scan(&ticket_period, &award_nums,)
		float,err := strconv.ParseFloat(ticket_period,64)
		checkErr(err)
		fmt.Println(ticket_period)
		fmt.Println(award_nums)
		fmt.Println(float)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println("affect2",affect)

	defer db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func task(){
	fmt.Println(time.Now())

	println("taskfsdkdlsjfkdfslk")
	c := cron.New(cron.WithSeconds())
	var v float64 = 0
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Every hour on the half hour")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		//log.Panicln("当前期是：：",v)
		fmt.Printf("%%.当前期0f:\t%.0f\n", v)

		var ns = goget.LuckAirShip(v)
		if ns!=v { //最后一期或者获得成功之后推迟
			c.Stop()
			if int64(ns)%1000==180 { //最后一期刷新
				time.AfterFunc(time.Second * 280, func() {
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"sdfffffffffff")
					c.Start()
				})
			}else{
				time.AfterFunc(time.Minute * 470, func() {
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"sdfffffffffff")
					c.Start()
				})
			}
		}
		if ns==v&&int64(ns)%1000==180 { //最后一期休息时间
			c.Stop()
			time.AfterFunc(time.Minute*5, func() {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"sdfffffffffff")
				c.Start()
			})
		}
		v = ns
		fmt.Printf("%%.0被更改了f:\t%.0f\n", ns)

	})
	c.Start()
	//select{}
	//c.Stop()
	//time.AfterFunc(time.Second * 4, func() {
	//	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"sdfffffffffff")
	//	c.Start()
	//})
}


