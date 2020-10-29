package main

import (
	"fmt"
	"net/http"
	"html/template"
)



func sayhello(w http.ResponseWriter,r *http.Request) {
	//1. 定义模板

	//2. 解析模板
	t,err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err = ",err)
		return
	}

	//3. 渲染模板
	err = t.Execute(w,"小王子")
	if err != nil {
		fmt.Println("render template failed,err = ",err)
		return
	}
}

func main() {
	http.HandleFunc("/",sayhello)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println("http serve start failed,err = ",err)
		return
	}
}
