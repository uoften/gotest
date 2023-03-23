package main

import (
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	taskCount = 1
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 下载图片，传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:/home/download" + filename
	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImg    = `//[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// myTest()
	// DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201909/10242/e5794daf58_4.jpg", "1.jpg")

	// 1.初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, taskCount)
	// 2.爬虫协程
	for i := 259105; i < 259106; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://aabubbletea.com/products/2020-newest-multi-purpose-reusable-jar-ziplock-bags")
	}
	// 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

// 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile2("https:"+url, filename)
		if ok {
			//fmt.Printf("%s 下载成功\n", filename)
		} else {
			//fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}
func DownloadFile2(url,fileName string) bool{
	//fmt.Println(url)
	//fmt.Println(fileName)
	return true
}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]

	// 时间戳解决重名
	//timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	//filename = timePrefix + "_" + filename
	return
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == taskCount {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

type Attrs struct {
	Name string `json:"name"`
	Value []string `json:"value"`
}

// 获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	reg, _ := regexp2.Compile("<title>(.*)</title>", 0)
	m, _ := reg.FindStringMatch(pageStr)
	str1 := fmt.Sprintf("%s",m)
	str1 = strings.Replace(str1,"<title>","",-1)
	str1 = strings.Replace(str1,"</title>","",-1)
	fmt.Println(str1)

	reg2, _ := regexp2.Compile(":variant_attrs=\"(.*)}]\"", 0)
	m2, _ := reg2.FindStringMatch(pageStr)
	str := fmt.Sprintf("%s",m2)
	str = strings.Replace(str,":variant_attrs=\"","",-1)
	str = strings.Replace(str,"}]\"","}]",-1)
	jsonStr := html.UnescapeString(str)
	var asd []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &asd); err != nil {
		fmt.Printf("json反序列化失败%v",err)
	}
	fmt.Println(asd)

	//reg3, _ := regexp2.Compile(":product='(.*)}' :", 0)
	//m3, _ := reg3.FindStringMatch(pageStr)
	//str2 := fmt.Sprintf("%s",m3)
	//str2 = strings.Replace(str2,":product='","",-1)
	//str2 = strings.Replace(str2,"}' :","}",-1)
	//jsonStr2 := html.UnescapeString(str2)
	//fmt.Println(jsonStr2)
	//var asd2 map[string]interface{}
	//if err := json.Unmarshal([]byte(jsonStr2), &asd2); err != nil {
	//	fmt.Printf("json反序列化失败%v",err)
	//}
	//fmt.Println(asd2)

	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}