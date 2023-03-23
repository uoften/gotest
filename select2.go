package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {

	var hostListCh chan string
	hostListCh = make(chan string,5)
	go func() {
		for i:=0;i< 100;i++ {
			//fmt.Println("in"+strconv.Itoa(i))
			hostListCh <- "asd"+strconv.Itoa(i)
		}
	}()
	//fmt.Println(cap(hostListCh))
	//for i:=0;i< cap(hostListCh);i++ {
	//	hostListCh <- "asd"
	//}
	//优雅关channel
	reCount:=0
	for {
		select {
		case ss,ok := <-hostListCh :
			if ok {
				fmt.Println(ss)
			}else{

			}
		default:
			time.Sleep(10 * time.Millisecond)
			reCount += 1
		}
		if reCount == 5 {
			break
		}
	}
	fmt.Println("the end")
}