// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
    "encoding/json"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

var post = "8088"
var iPost = "8777"
var proxyHost = "http://localhost:8088"

func main() {
    a := app.New()
    a.SetIcon(theme.FyneLogo())
    w := a.NewWindow("前端换源工具")

    hello := widget.NewLabel("Hello Fyne!")
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("启动", func() {
            hello.SetText("Welcome :)")
        }),
    ))

    w.ShowAndRun()
    //fmt.Println(LoadConfig())
    //init := Config{Name: "不知道你说什么东西"}
    //SaveConfig(&init)
}

type Config struct {
    Name   string  `json:"name"`
}
func FyneLogo() fyne.Resource {
    return fynelogo
}

var configPath = "./config.json"

func SaveConfig(config *Config){

    data,err:=json.Marshal(config)
    if err!=nil{
        log.Fatal(err)
    }
    err=ioutil.WriteFile(configPath,data,0660)
    if err!=nil{
        log.Fatal(err)
    }

}

func LoadConfig()(config *Config){
    data,err:=ioutil.ReadFile(configPath)
    if err!=nil{
        log.Fatal(err)
    }
    //for _, line := range strings.Split(string(data), "\n") {
    //    fmt.Println(line)
    //}
    config =&Config{}
    err=json.Unmarshal(data,&config)
    if err!=nil{
        log.Fatal(err)
    }
    return config
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

