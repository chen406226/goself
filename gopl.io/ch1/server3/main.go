// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", handlerUpload)
	log.Fatal(http.ListenAndServe("localhost:8777", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

type Data struct {
	Code int `json:"code"`
	Message string`json:"msg"`
}

//!-handler
func handlerUpload(w http.ResponseWriter, r *http.Request) {
	// Here the parameter is the size of the form data that should
	// be loaded in memory, the remaining being put in temporary
	// files
	fmt.Println("jinlaile")
	r.ParseMultipartForm(0)
	//r.ParseMultipartForm(0)
	if r.MultipartForm != nil {
		fmt.Println(r.MultipartForm.Value,"dsfffffffffffffffffffff")
	}
	fmt.Println(r.FormValue("file"),"fileddddddd")
	fmt.Println(r.FormValue("token"),"ddddddddddddddddddddddd")
	jData, err := json.Marshal(Data{200,"http://www.baidu.com/img"})
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
