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

func main() {
    //fmt.Println("输入debug端口")
    //fmt.Scan(&post)
    //fmt.Println("输入服务端口")
    //fmt.Scan(&iPost)
    http.HandleFunc("/", handler)
    fmt.Println("请访问"+iPost+"端口")
    log.Fatal(http.ListenAndServe("0.0.0.0:"+iPost, nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
    remote, err := url.Parse("http://localhost:"+post)
    if err != nil {
        panic(err)
    }
    r.Host = "localhost:"+iPost
    proxy := httputil.NewSingleHostReverseProxy(remote)
    proxy.ServeHTTP(w, r)
}
