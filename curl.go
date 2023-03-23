package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {


	fmt.Println(strconv.FormatInt(time.Now().Unix(),10))
	fmt.Println("------------------------------------------")
	ns := time.Now().UnixNano() // 获得当前单位为纳秒的时间戳

	fmt.Println("时间戳（秒秒）：", ns/1e9)		// 输出：时间戳（秒） ： 1665807442
	fmt.Println("时间戳（毫秒）：", ns/1e6)	// 输出：时间戳（毫秒）： 1665807442207
	fmt.Println("时间戳（微秒）：", ns/1e3)	// 输出：时间戳（微秒）： 1665807442207974
	fmt.Println("时间戳（纳秒）：", ns)		// 输出：时间戳（纳秒）： 1665807442207974500
	//pageStr,err := getPageStr("https://blog.csdn.net/Javatutouhouduan/article/details/128150631?spm=1001.2100.3001.7377&utm_medium=distribute.pc_feed_blog.none-task-blog-personrec_tag-4-128150631-null-null.nonecase&depth_1-utm_source=distribute.pc_feed_blog.none-task-blog-personrec_tag-4-128150631-null-null.nonecase")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf(pageStr)

	////生成client 参数为默认
	//client := &http.Client{}
	//
	////生成要访问的url
	//url := "http://www.baidu.com"
	//
	////提交请求
	//reqest, err := http.NewRequest("GET", url, nil)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	////处理返回结果
	//response, _ := client.Do(reqest)
	//
	////将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//stdout := os.Stdout
	//_, err = io.Copy(stdout, response.Body)
	//
	////返回的状态码
	//status := response.StatusCode
	//
	//fmt.Println(status)
}
// 抽取根据url获取内容
func getPageStr(url string) (pageStr string,err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "",errors.New(fmt.Sprintf("访问商品链接出错：%v",err.Error()))
	}else{
		defer res.Body.Close()  // 关闭连接流
		body, _ := ioutil.ReadAll(res.Body)
		pageStr = string(body)
		return pageStr,nil
	}
}