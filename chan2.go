package main

import (
	"fmt"
	"time"
)

func main() {
	personList := []string{"1","2","3","4"}
	chanP := make(chan string,1)
	for _, p := range personList {
		go func() {
			time.Sleep(time.Second*1)
			chanP <- p
		}()
		//等待阻塞读取管道
		str := <- chanP
		fmt.Println(str)
	}
	fmt.Println("end")
}