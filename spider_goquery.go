package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/panjf2000/ants/v2"
	"gotest/common/helper"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

type GetUserActivitiesReq struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	IsOwner bool `json:"isOwner"`
	PageNumber int `json:"pageNumber"`
	PageSize int `json:"pageSize"`
	Uid int `json:"uid"`
}

type GetHotReq struct {
	PageNumber int `json:"pageNumber"`
	SortType string `json:"sortType"`
	CategoryId int `json:"categoryId"`
}

type RespData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data Data `json:"data"`
}

type Data struct {
	Total int `json:"total"`
	List []List `json:"list"`
}
type List struct {
	ActionInfo ActionInfo `json:"action_info"`
}
type ActionInfo struct {
	ArticleId int `json:"article_id"`
	FollowedUid int `json:"followed_uid"`
}

func main() {
	pageStr, err := helper.GetPageStr("https://developer.aliyun.com/article/970538")
	if err != nil {
		fmt.Println(err)
	}
	d, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`"larkContent": '(.*)',`)
	platformRegRegMatch := re.FindStringSubmatch(pageStr)
	var articleContent string
	if len(platformRegRegMatch) >0 {
		articleContent = platformRegRegMatch[0]
	}
	if articleContent=="" {
		fmt.Println(111)
		fmt.Println(articleContent)
	}

	//articleContent,err := d.Find(".J-articleContent").Html()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(articleContent)

	//desReg := regexp.MustCompile("<meta name=\"description\" content=\"(.*)\"/>")
	//desRegRegMatch := desReg.FindStringSubmatch(pageStr)
	//if len(desRegRegMatch) > 0 {
	//	fmt.Println(desRegRegMatch[1])
	//}
	//keyReg := regexp.MustCompile("<meta name=\"keywords\" content=\"(.*)\"/>")
	//keyRegRegMatch := keyReg.FindStringSubmatch(pageStr)
	//if len(keyRegRegMatch) > 0 {
	//	fmt.Println(keyRegRegMatch[1])
	//}

	//keyReg := regexp.MustCompile("<meta name=\"keywords\" content=\"(.*)\"/><meta name=\"subject-time\"")
	//keyRegRegMatch := keyReg.FindStringSubmatch(pageStr)
	//if len(keyRegRegMatch) > 0 {
	//	fmt.Println(len(keyRegRegMatch[1]))
	//	fmt.Println(keyRegRegMatch[1])
	//	if len(keyRegRegMatch[1])>1000 {
	//		fmt.Println(keyRegRegMatch[1][:1000])
	//	}
	//}
	//
	//desReg := regexp.MustCompile("<meta name=\"description\" content=\"(.*)\"/><meta name=\"viewport\"")
	//desRegRegMatch := desReg.FindStringSubmatch(pageStr)
	//if len(desRegRegMatch) > 0 {
	//	fmt.Println(len(desRegRegMatch[1]))
	//	fmt.Println(desRegRegMatch[1])
	//	if len(desRegRegMatch[1])>1000 {
	//		fmt.Println(desRegRegMatch[1][:1000])
	//	}
	//}

	//articleCategoryStr := d.Find("a").Find(".col-tag")

	//articleCategoryStr := d.Find("a[class=c-pages-item]").Last().Text()
	//fmt.Println(articleCategoryStr)
	//
	//reqData := GetUserActivitiesReq{}
	//reqData.Payload.IsOwner = false
	//reqData.Payload.PageNumber =3
	//reqData.Payload.PageSize =20
	//reqData.Payload.Uid =2871589
	//header := make(map[string]string)
	//header["Content-Type"] = "application/json"
	////header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	//res, err := library.HttpPost("https://cloud.tencent.com/developer/services/ajax/user-center?action=GetUserActivities", reqData, header, map[string]int{}, map[string]string{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//re := regexp.MustCompile("\"articleId\":([0-9]+?),")
	//productMatchData := re.FindAllString(string(res), -1)
	//fmt.Println()
	//if len(productMatchData)>0 {
	//	for _,v:=range productMatchData{
	//		re := regexp.MustCompile("[0-9]+")
	//		vv := re.FindAllString(v, -1)
	//		if len(vv)>0 {
	//			fmt.Println(vv[0])
	//		}
	//	}
	//}

	//for i := 0; i < articleCategoryStr.Length(); i++ {
	//	articleCategory := articleCategoryStr.Eq(i).Text()
	//	fmt.Println(articleCategory)
	//}

	t := d.Find("li[class=panel-cell]")
	var urls []string
	if html, _ := t.Html(); html != "" {
		for i := 0; i < t.Length(); i++ {
			tStr, _ := t.Eq(i).Html()
			dd, err := goquery.NewDocumentFromReader(strings.NewReader(tStr))
			if err != nil {
				fmt.Println(err)
			}else{
				dd.Find("a").Each(func(i int, selection *goquery.Selection) {
					href, ok := selection.Attr("href")
					if ok {
						link, err := url.Parse(href)
						if err==nil{
							urls = append(urls,link.String())
						}
					}
				})
			}
		}
	}
	fmt.Println(len(urls))
	defer ants.Release()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(32, func(i interface{}) {
		str := i.(string)
		// todo
		fmt.Println(str)
		wg.Done()
	})
	defer p.Release()
	for _,v := range urls{
		wg.Add(1)
		_ = p.Invoke(v)
	}
	wg.Wait()
}
