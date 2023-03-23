package library

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string,err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", url)
	//c1 := http.Cookie{
	//	Name: "client_id",
	//	Value: strconv.FormatInt(time.Now().UnixMicro(),10),
	//	Path: "/",
	//	Domain: "lindress.com",
	//}
	//req.Header.Add("Set-Cookie", c1.String())
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

func PostPageStr(url string) (pageStr string,err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", url)
	//c1 := http.Cookie{
	//	Name: "client_id",
	//	Value: strconv.FormatInt(time.Now().UnixMicro(),10),
	//	Path: "/",
	//	Domain: "lindress.com",
	//}
	//req.Header.Add("Set-Cookie", c1.String())
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

