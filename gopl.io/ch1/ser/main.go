package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var count int
var sum int
var msg = ""
func main() {
	http.HandleFunc("/time", handler)
	http.HandleFunc("/love", counter)
	http.HandleFunc("/reset", reset)
	msg = LoadMsg()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func LoadMsg() string {
	data,err:=ioutil.ReadFile("/home/server/lulu/lu.txt")
	if err!=nil{
		//log.Println(err)
		return ""
	}

	return string(data)
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("定时器开始---",count, "结束；；；；；；")
}
// handler echoes the Path component of the requested URL.
func reset(w http.ResponseWriter, r *http.Request) {
	msg = LoadMsg()
	log.Println("更新了MSG------")
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["phone"]
	sum++
	fmt.Println("来了---", sum, "结束；；；；；；")
	if !ok || len(keys) < 1 {
		log.Println("Url Param 'phone' is missing")
		return
	}
	if len(msg) == 0 {
		msg = LoadMsg()
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	if key !="15057402742" {
		fmt.Fprintf(w, "手机号错误")
		return
	}
	count++
	log.Println("查看了---",count, "结束；；；；；；")
	fmt.Fprintln(w, msg)

	//fmt.Sprintf("Count %d\n", count)
	//fmt.Fprintf(w, "Count %d\n", count)
}

