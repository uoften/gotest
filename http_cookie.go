package main

import (
	"strings"
	"fmt"
	"io"
	"net/http"
)

//cookie设置方法一
func Cookie(w http.ResponseWriter,r *http.Request){
	ck:=&http.Cookie{
		Name:"myCookie",
		Value:"hello",
		Path:"/",
		Domain:"localhost",
		MaxAge:120,
	}

	http.SetCookie(w,ck)

	ck2,err:=r.Cookie("myCookie")

	if err!=nil{
		io.WriteString(w,err.Error())
		return
	}

	io.WriteString(w,ck2.Value)

}

//cookie设置方法二
func Cookie2(w http.ResponseWriter,r *http.Request){
	ck:=&http.Cookie{
		Name:"myCookie2",
		Value:"hello world",
		Path:"/",
		Domain:"localhost",
		MaxAge:120,
	}

	w.Header().Set("Set-Cookie",strings.Replace(ck.String()," ","%20",-1)) //http包中将空格视为非法，所以需要在此处添加空格

	ck2,err:=r.Cookie("myCookie2")

	if err!=nil{
		io.WriteString(w,err.Error())
		return
	}

	io.WriteString(w,ck2.Value)
}

func main(){

	http.HandleFunc("/",Cookie)
	http.HandleFunc("/2",Cookie2)

	fmt.Println("listen......")

	http.ListenAndServe(":8081",nil)
}