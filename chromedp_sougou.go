package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/chromedp"

	"log"
)

type Keyword struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Level        int    `json:"level"`
	ArticleCount int    `json:"article_count"`
	LastTime     int64  `json:"last_time"` //上次执行时间
}

type Article struct {
	Id          int64  `json:"id"`
	KeywordId   int64  `json:"keyword_id"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	OriginUrl   string `json:"origin_url"`
	Status      int    `json:"status"`
	CreatedTime int    `json:"created_time"`
	UpdatedTime int    `json:"updated_time"`
	Content     string `json:"content"`
	ContentText string `json:"-"`
}

func main() {
	word := Keyword{
		Name:         "golang",
	}

	result := CollectArticleFromWeixin(&word)

	for i, v := range result {
		log.Println(i, v.Title, len(v.Content), v.OriginUrl)
		log.Println("纯内容：", v.ContentText)
	}
}

func CollectArticleFromWeixin(keyword *Keyword) []*Article {
	timeCtx, cancel := context.WithTimeout(GetChromeCtx(false), 30*time.Second)
	defer cancel()

	var collectLink string
	err := chromedp.Run(timeCtx,
		chromedp.Navigate(fmt.Sprintf("https://weixin.sogou.com/weixin?p=01030402&query=%s&type=2&ie=utf8", keyword.Name)),
		chromedp.WaitVisible(`//ul[@class="news-list"]`),
		chromedp.Location(&collectLink),
	)
	fmt.Println(collectLink)
	if err != nil {
		log.Println("读取搜狗搜索列表失败1：", keyword.Name, err.Error())
		return nil
	}
	log.Println("正在采集列表：", collectLink)
	var aLinks []*cdp.Node
	if err := chromedp.Run(timeCtx, chromedp.Nodes(`//ul[@class="news-list"]//h3//a`, &aLinks)); err != nil {
		log.Println("读取搜狗搜索列表失败2：", keyword.Name, err.Error())
		return nil
	}

	var articles []*Article
	for i := 0; i < len(aLinks); i++ {
		href := aLinks[i].AttributeValue("href")
		href, _ = joinURL("https://weixin.sogou.com/", href)
		article := &Article{}
		err = chromedp.Run(timeCtx,
			chromedp.Navigate(href),
			chromedp.WaitVisible(`#js_article`),
			chromedp.Location(&article.OriginUrl),
			chromedp.Text(`#activity-name`, &article.Title, chromedp.NodeVisible),
			chromedp.InnerHTML("#js_content", &article.Content, chromedp.ByID),
			chromedp.Text("#js_content", &article.ContentText, chromedp.ByID),
		)
		if err != nil {
			log.Println("读取搜狗搜索列表失败3：", keyword.Name, err.Error())
		}
		log.Println("采集文章：", article.Title, len(article.Content), article.OriginUrl)
		articles = append(articles, article)
	}

	return articles
}

// 重组url
func joinURL(baseURL, subURL string) (fullURL, fullURLWithoutFrag string) {
	baseURL = strings.TrimSpace(baseURL)
	subURL = strings.TrimSpace(subURL)
	baseURLObj, _ := url.Parse(baseURL)
	subURLObj, _ := url.Parse(subURL)
	fullURLObj := baseURLObj.ResolveReference(subURLObj)
	fullURL = fullURLObj.String()
	fullURLObj.Fragment = ""
	fullURLWithoutFrag = fullURLObj.String()
	return
}

//检查是否有9222端口，来判断是否运行在linux上
func checkChromePort() bool {
	addr := net.JoinHostPort("", "9222")
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// ChromeCtx 使用一个实例
var ChromeCtx context.Context
func GetChromeCtx(focus bool) context.Context {
	if ChromeCtx == nil || focus {
		allocOpts := chromedp.DefaultExecAllocatorOptions[:]
		allocOpts = append(allocOpts,
			chromedp.DisableGPU,
			chromedp.Flag("blink-settings", "imagesEnabled=false"),
			chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36`),
			chromedp.Flag("accept-language", `zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6`),
		)

		if checkChromePort() {
			// 不知道为何，不能直接使用 NewExecAllocator ，因此增加 使用 ws://127.0.0.1:9222/ 来调用
			c, _ := chromedp.NewRemoteAllocator(context.Background(),  "ws://127.0.0.1:9222/")
			ChromeCtx, _ = chromedp.NewContext(c)
		} else {
			c, _ := chromedp.NewExecAllocator(context.Background(), allocOpts...)
			ChromeCtx, _ = chromedp.NewContext(c)
		}
	}

	return ChromeCtx
}