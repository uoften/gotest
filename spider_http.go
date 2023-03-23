package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

// 负责抓取页面的源代码(html)
// 通过http包实现
func fetch(url string) string {

	// 得到一个客户端
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Mobile Safari/537.36")
	request.Header.Add("Cookie", "test_cookie=CheckForPermission; expires=Tue, 30-Aug-2022 01:04:32 GMT; path=/; domain=.doubleclick.net; Secure; HttpOnly; SameSite=none")

	// 客户端发送请求，并且获取一个响应
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error: ", err)
		return ""
	}

	// 如果状态码不是200，就是响应错误
	if response.StatusCode != 200 {
		log.Println("Error: ", response.StatusCode)
		return ""
	}

	defer response.Body.Close() // 关闭

	// 读取响应体中的所有数据到body中，这就是需要的部分
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error: ", err)
		return ""
	}

	// 转换为字符串(字节切片 --> 字符串)
	return string(body)
}

var waitGroup sync.WaitGroup

// 解析页面源代码
func parseURL(body string) {

	// 将body(响应结果)中的换行替换掉，防止正则匹配出错
	html := strings.Replace(body, "\n", "", -1)
	// 正则匹配
	re_Img_div := regexp.MustCompile(`<div class="img_wrapper">(.*?)</div>`)

	img_div := re_Img_div.FindAllString(html, -1) // 得到<div><img/></div>

	for _, v := range img_div {

		// img正则
		re_link := regexp.MustCompile(`src="(.*?)"`)
		// 找到所有的图片链接
		links := re_link.FindAllString(v, -1) // 得到所有图片链接

		// 遍历links，切掉不必要的部分src="和最后的"
		for _, v := range links {

			src := v[5 : len(v)-1]
			src = "http:" + src

			waitGroup.Add(1)
			go downLoad(src)
		}
	}

}

// 下载
func downLoad(src string) {

	fmt.Println("================================", src)

	// 取一个文件名
	filename := string(src[len(src)-8 : len(src)])
	fmt.Println(filename)

	response, _ := http.Get(src)
	picdata, _ := ioutil.ReadAll(response.Body)

	image, _ := os.Create("D:\\home\\download\\" + filename)
	image.Write(picdata)

	defer func() {
		image.Close()
		waitGroup.Done()
	}()
}

func main() {

	url := "https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf"

	body := fetch(url)
	// fmt.Println(body)
	parseURL(body)

	waitGroup.Wait()
}