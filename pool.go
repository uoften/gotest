package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func baidu(id int) {
	//生成要访问的url
	url := "http://www.baidu.com?id="+id
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	return resp

}

func main() {
	for id := 2; id < 9202; id++ {
		wg.Add(1)
		go baidu(id)
		if (id % 300) == 0 {
			time.Sleep(time.Duration(2000) * time.Millisecond)
		}
	}
	wg.Wait()
}