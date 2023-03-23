package main

import (
	"fmt"
	"regexp"
	"strings"
)
var cmdPageHintMap = func() map[string]string {return map[string]string {
	"1000": "M",
	"900":  "CM",
}}

func main()  {
	regCmd := regexp.MustCompile("more --")
	result := regCmd.ReplaceAllString("-- more --", "")
fmt.Println(result)

	fmt.Println(strings.Split("2|||1", "|||"))
	fmt.Println(len(strings.Split("2", "|||")))
	//fmt.Println(strings.Contains("aasdas dsa da ssss", "ssss"))
	//fmt.Println(cmdPageHintMap()["10100"])
	var in chan string
	var out chan string
	in = make(chan string, 5)
	out = make(chan string, 5)
	go func() {
		for i:=0;i<10;i++ {
			in<-"asd"
			//out<-"asdasd"
		}
	}()
	//close(in)
	//close(out)
	failed:=0
	for {
		//select {
		//case _, received := <-in:
		//	if received {
		//		fmt.Println(received)
		//	}
		//default:
		//	fmt.Println(1111)
		//	time.Sleep(10 * time.Millisecond)
		//	failed += 1
		//	if failed%10 == 0 {
		//		fmt.Println("failed:" + strconv.Itoa(failed))
		//	}
		//}
		select {
		case v, ok := <-in:
			if ok {
				fmt.Println(v)
				close(in)
			} else {
				in = nil
			}
		case _, ok := <-out:
			if ok {
				close(out)
			} else {
				out = nil
			}
		default:
			failed++
			break
		}
		if failed == 100 {
			break
		}else{
			fmt.Println("ko")
		}
	}
	sd := <- in
	fmt.Println(sd)
	fmt.Println("ok")
}