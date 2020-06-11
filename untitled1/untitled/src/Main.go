package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
)


func main() {
	LuckAirShip()
	//openMysql()
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
type Profile struct {
	Name    string
	//Hobbies []string
	Hobbies  []interface{}
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
	
	obj := map[string]interface{}{"name":"ck","sex":"Man","age":18}
	//slice1 := [] interface{}{"不知道SB",34,"sdf"}
	//profile := Profile{"Alex", []interface{}{"snowboarding", "programming"}}
	//pro := Profile{"Alex", slice1}
	//jss, _ := json.Marshal(pro)
	res := Res{"200","成功","8"}

	fmt.Println("==========================",res)
	//js, err := json.Marshal(res)
	js, err := json.Marshal(obj)

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
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
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
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
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

	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}