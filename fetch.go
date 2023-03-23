// fetch.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

// 建立正则
var RE = regexp.MustCompile("\\w+\\.\\w+$")

func main() {
	urls := []string {
		"http://www.qq.com",
		"http://www.163.com",
		"http://www.sina.com",
	}

	// 建立channel
	ch := make(chan string)

	// 开始时间
	start := time.Now()

	for _, url := range urls {
		// 开启一个goroutine
		go fetch(url, ch)
	}

	for range urls {
		// 打印channel中的信息
		fmt.Println(<-ch)
	}

	// 总消耗的时间
	elapsed := time.Since(start).Seconds()

	fmt.Printf("%.2fs elapsed\n", elapsed)
}

// 根据URL获取资源内容
func fetch(url string, ch chan<- string) {
	start := time.Now()

	// 发送网络请求
	res, err := http.Get(url)

	if err != nil {
		// 输出异常信息
		ch <- fmt.Sprint(err)
		os.Exit(1)
	}

	// 读取资源数据
	body, err := ioutil.ReadAll(res.Body)

	// 关闭资源
	res.Body.Close()

	if err != nil {
		// 输出异常信息
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		os.Exit(1)
	}

	// 写入文件
	ioutil.WriteFile(getFileName(url), body, 0644)

	// 消耗的时间
	elapsed := time.Since(start).Seconds()

	// 输出单个URL消耗的时间
	ch <- fmt.Sprintf("%.2fs %s", elapsed, url)
}

// 获取文件名
func getFileName(url string) string {
	// 从URL中匹配域名部分
	return RE.FindString(url) + ".txt"
}