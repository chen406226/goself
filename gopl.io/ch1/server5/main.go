// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

var post = "8088"
var iPost = "8777"
var proxyHost = "http://localhost:8088"

func main() {
    //fmt.Println("输入debug端口")
    //fmt.Scan(&post)
    //fmt.Scan(&iPost)
    fmt.Println("输入服务端口")
    fmt.Println("EP: http://10.10.0.123:8088")
    fmt.Println(">> Enter启动")
    fmt.Scan(&proxyHost)
    http.HandleFunc("/", handler)
    //fmt.Println("请访问"+iPost+"端口")
    fmt.Println("-------------")
    fmt.Println("运行中不要关闭本窗口")
    fmt.Println("关闭本窗口重启以更换服务地址")
    fmt.Println("-------------")
    log.Fatal(http.ListenAndServe("0.0.0.0:9876", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
    remote, err := url.Parse(proxyHost)
    if err != nil {
        panic(err)
    }
    r.Host = "localhost:"+iPost
    proxy := httputil.NewSingleHostReverseProxy(remote)
    proxy.ServeHTTP(w, r)
}
