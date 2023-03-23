package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"gotest/library"
	"sort"
	"strconv"
	"strings"
	"time"
)

var urls []string

func main() {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c2 := c1.Clone()

	//异步
	c2.Async = true
	//限速
	c2.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*.downxia.com/downinfo/*",
		Delay:        0,
		RandomDelay:  0,
		Parallelism:  1,
	})
	//采集器1，获取文章列表
	c1.OnHTML("div[class='bor well2']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			title := item.ChildText("div[class='top'] > a > strong")
			href := item.ChildAttr("div[class='bottom'] > a", "href")
			ctx := colly.NewContext()
			ctx.Put("href", href)
			ctx.Put("title", title)
			//通过Context上下文对象将采集器1采集到的数据传递到采集器2
			c2.Request("GET", href, nil, ctx, nil)
		})
	})
	//采集器2，获取文章详情
	c2.OnHTML("ul[class='list-down fix']", func(e *colly.HTMLElement) {
		//href := e.Request.Ctx.Get("a[uri]")
		//title := e.Request.Ctx.Get("h3[class='down-item-tit']")
		//summary := e.Request.Ctx.Get("div[class='content']")
		content := e.Request.Ctx.Get("div[class='txt']")
		e.ForEach("a[uri]", downloadImages)
		//detail := e.Text

		//fmt.Println("----------" + title + "----------")
		//fmt.Println(href)
		//fmt.Println(summary)
		//fmt.Println(detail)
		fmt.Println(content)
	})

	c1.OnRequest(func(r *colly.Request) {
		//fmt.Println("c1爬取页面：", r.URL)
	})

	c2.OnRequest(func(r *colly.Request) {
		//fmt.Println("c2爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("http://www.downxia.com/downlist/s_37_1.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	c2.Wait()
	fmt.Println(len(urls))
}

func downloadImages(_ int, e *colly.HTMLElement) {
	srcRef := e.Attr("uri")
	if !inArray(srcRef,urls) {
		urls = append(urls,srcRef)
		filename := spiderGetFilename(srcRef)
		spiderDownloadFile(srcRef, filename)
	}
}

func spiderDownloadFile(srcRef,filename string) {
	for i:=0;i<7;i++ {
		ok := library.DownloadFile(getRandUrl(i)+srcRef, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
			break
		} else {
			fmt.Printf("%s 下载失败\n", filename)
			continue
		}
	}
}

func spiderGetFilename(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]

	// 时间戳解决重名
	if !strings.Contains(filename,".") {
		filename = strconv.Itoa(int(time.Now().UnixNano()))
	}
	return
}

func inArray(target string, str_array []string) bool {

	sort.Strings(str_array)

	index := sort.SearchStrings(str_array, target)

	if index < len(str_array) && str_array[index] == target {

		return true

	}

	return false

}

func getRandUrl(num int) string{
	consumerArr := [8]string{
		"http://down1.downxia.com/",
		"http://zj.downxia.com/",
		"http://sd.downxia.com/",
		"http://js.downxia.com/",
		"http://cnc.downxia.com/",
		"http://lnlt.downxia.com/",
		"http://yd.downxia.com/",
		"http://tt.downxia.com/",
	}
	//rand.Seed(time.Now().UnixNano())
	//num := rand.Intn(10)
	return consumerArr[num]
}