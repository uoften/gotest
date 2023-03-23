package main

import (
	"fmt"
	"io"
	"sync"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)



func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页的body内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF{
				break
			}else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}


//爬取网页
func spiderPage(url string) string {

	fmt.Println("正在爬取", url)
	//爬,将所有的网页内容爬取下来
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	//把内容写入到文件
	filename := strconv.Itoa(rand.Int()) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil{
		fmt.Println(err1)
	}
	//写内容
	f.WriteString(result)
	//关闭文件
	f.Close()
	return url + " 抓取成功"

}

func asyn_worker(page chan string, results chan string,wg *sync.WaitGroup){

	defer wg.Done()  //defer wg.Done()必须放在go并发函数内

	for{
		v, ok := <- page //显示的调用close方法关闭通道。
		if !ok{
			fmt.Println("已经读取了所有的数据，", ok)
			break
		}
		//fmt.Println("取出数据：",v, ok)
		results <- spiderPage(v)
	}

	//for n := range page {
	//  results <- spiderPage(n)
	//}
}

func doWork(start, end int,wg *sync.WaitGroup) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	//因为很有可能爬虫还没有结束下面的循环就已经结束了，所以这里就需要且到通道
	page := make(chan string,100)
	results := make(chan string,100)


	go sendResult(results,start,end)

	go func() {

		for i := 0; i <= 20; i++ {
			wg.Add(1)
			go asyn_worker(page, results, wg)
		}
	}()


	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		page <- url
		println("加入" + url + "到page")
	}
	println("关闭通道")
	close(page)

	wg.Wait()
	//time.Sleep(time.Second * 5)
	println(" Main 退出 。。。。。")
}


func sendResult(results chan string,start,end int)  {

	//for i := start; i <= end; i++ {
	//  fmt.Println(<-results)
	//}

	// 发送抓取结果
	for{
		v, ok := <- results
		if !ok{
			fmt.Println("已经读取了所有的数据，", ok)
			break
		}
		fmt.Println(v)

	}
}

func main() {
	start_time := time.Now().UnixNano()

	var wg sync.WaitGroup

	doWork(1,200, &wg)
	//输出执行时间，单位为毫秒。
	fmt.Printf("执行时间: %ds",(time.Now().UnixNano() - start_time) / 1000)

}