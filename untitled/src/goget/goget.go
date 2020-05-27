package goget

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)


func LuckAirShip(fv float64) (float64) {
	v:=fv
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.193:3306)/test?charset=utf8")
	checkErr(err)
	if v==0 {
		println("是零")
		//查询数据
		rows, err := db.Query("SELECT ticket_period ,award_nums FROM cp_ticket_award_record_8 ORDER BY ticket_period DESC LIMIT 0,1")
		checkErr(err)

		for rows.Next() {
			var ticket_period string
			var award_nums string
			err = rows.Scan(&ticket_period, &award_nums,)
			float,err := strconv.ParseFloat(ticket_period,64)
			v = float
			checkErr(err)
			//fmt.Println(ticket_period)
			//fmt.Println(award_nums)
			//fmt.Println(float)
		}
	}
	//return 0

	params := url.Values{}

	Url, err := url.Parse("https://luckylotterys.com/api/latest/getLotteryPksInfo.do?lotCode=10057")
	if err != nil {
		panic(err.Error())

	}
	params.Set("lotCode", "10057")
	params.Set("id", string("1"))
	////如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	println(urlPath)
	//resp, err := http.Get(urlPath)
	//defer resp.Body.Close()
	//s, err := ioutil.ReadAll(resp.Body)
	resp, err := http.Get("https://luckylotterys.com/api/latest/getLotteryPksInfo.do?lotCode=10057")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(s)), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		fmt.Println(dat["result"])
		if dat["result"]==nil {
			return fv
		}
		result,ok := dat["result"].(map[string]interface{})

		if ok {
			data,oks := result["data"].(map[string]interface{})
			if oks {
				inttype,okss := data["preDrawIssue"].(float64)
				fmt.Printf("%%.返回值0f:\t%.0f\n", inttype)
				if okss&& float64(inttype)<=v {
					if fv==0 {
						return fv
					}
					return v
				}


				//插入数据
				stmt, err := db.Prepare("INSERT INTO cp_ticket_award_record_8 SET ticket_id=?,ticket_period=?,award_nums=?,award_time=?,next_ticket_period=?")
				checkErr(err)

				res, err := stmt.Exec(8, data["preDrawIssue"], data["preDrawCode"], data["preDrawTime"],data["drawIssue"])
				checkErr(err)

				id, err := res.LastInsertId()
				checkErr(err)
				fmt.Println("id=:",id)
				draw,ok := data["preDrawIssue"].(float64)
				println(draw,ok)
				fmt.Printf("%%.0f:\t%.0f\n", draw)
				if ok {
					//log.Panicln("成功成功")
					println("chenggong")
					return draw
				}
			}

		}
		//log.Panicln(data)
	}
	db.Close()

	fmt.Printf(string(s))
	return fv
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}