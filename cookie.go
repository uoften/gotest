package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)

}

//1 get url response html
func getUrlRespHtml(url string) string {
	fmt.Printf("\ngetUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	httpReq, err := http.NewRequest("GET", url, nil)
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("\nhttp get url=%s response error=%s\n", url, err.Error())
	}
	fmt.Printf("\nhttpResp.Header=%s", httpResp.Header)
	fmt.Printf("\nhttpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("\nget response for url=%s got error=%s\n", url, errReadAll.Error())
	}
	//全局保存
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)

	respHtml = string(body)
	return respHtml
}

//2
func getUrlRespHtmlWithHeader(url, headers string) string {
	fmt.Printf("\ngetUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	httpReq, err := http.NewRequest("GET", url, nil)
	AddHeaders(httpReq, headers)
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("\nhttp get url=%s response error=%s\n", url, err.Error())
	}
	fmt.Printf("\nhttpResp.Header=%s", httpResp.Header)
	fmt.Printf("\nhttpResp.Status=%s", httpResp.Status)
	fmt.Printf("\nhttpResp.cookies=%s", httpResp.Cookies())

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("\nget response for url=%s got error=%s\n", url, errReadAll.Error())
	}
	//全局保存
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)

	respHtml = string(body)
	return respHtml
}

//3
func PostUrlRespHtmlWithHeader(url, headers, formdata string) string {
	fmt.Printf("\ngetUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	httpReq, err := http.NewRequest("POST", url, ioutil.NopCloser(strings.NewReader(formdata)))
	AddHeaders(httpReq, headers)
	httpReq.Header.Set("ContentType", "application/x-www-form-urlencoded")
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("\nhttp get url=%s response error=%s\n", url, err.Error())
	}
	fmt.Printf("\nhttpResp.Header=%s", httpResp.Header)
	fmt.Printf("\nhttpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("\nget response for url=%s got error=%s\n", url, errReadAll.Error())
	}
	//全局保存
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)

	respHtml = string(body)
	return respHtml
}

func dbgPrintCurCookies() {
	var cookieNum int = len(gCurCookies)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		fmt.Printf("\n\n\n\n------ Cookie [%d]------", i)
		fmt.Printf("\n\tName=%s", curCk.Name)
		fmt.Printf("\n\tValue=%s", curCk.Value)
		fmt.Printf("\n\tPath=%s", curCk.Path)
		fmt.Printf("\n\tDomain=%s", curCk.Domain)
		fmt.Printf("\n\tExpires=%s", curCk.Expires)
		fmt.Printf("\n\tRawExpires=%s", curCk.RawExpires)
		fmt.Printf("\n\tMaxAge=%d", curCk.MaxAge)
		fmt.Printf("\n\tSecure=%t", curCk.Secure)
		fmt.Printf("\n\tHttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("\n\tRaw=%s", curCk.Raw)
		fmt.Printf("\n\tUnparsed=%s", curCk.Unparsed)
	}
}

func AddHeaders(req *http.Request, headers string) *http.Request {
	//将传入的Header分割成[]ak和[]av
	a := strings.Split(headers, "\n")
	ak := make([]string, len(a[:]))
	av := make([]string, len(a[:]))
	//要用copy复制值；若用等号仅表示指针，会造成修改ak也就是修改了av
	copy(ak, a[:])
	copy(av, a[:])
	//fmt.Println(ak[0], av[0])
	for k, v := range ak {
		i := strings.Index(v, ":")
		j := i + 1
		ak[k] = v[:i]
		av[k] = v[j:]
		//设置Header
		req.Header.Set(ak[k], av[k])
	}
	return req
}
curl 'https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf' \
-H 'authority: www.lindress.com' \
-H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
-H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
-H 'cache-control: max-age=0' \
-H 'cookie: client_id=1669700317931191; _c_id=1669700317931391753; _identity_cart=b421fa35-839e-4adf-89bc-88f96ffcc6d0; store_locale=en-US; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22184c1e4f7dc660-039093a1231da9-26021e51-3686400-184c1e4f7ddaad%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22%24device_id%22%3A%22184c1e4f7dc660-039093a1231da9-26021e51-3686400-184c1e4f7ddaad%22%7D; _fbp=fb.1.1669700321439.1375841203; _identity_popups=51d46e33-865a-4990-a32a-e83673d398611669700344; _identity_popups_bundle=955c9b47-96f7-451b-814a-f46c5da7861a1669700344; session_id=1670583491799396; __cf_bm=a77dkb_LR0io19.hwfWa0J6IGr5VhrYp0_tVaL_uq3U-1670583522-0-ARxk4kn5/5J6UKOJ995ye5DNsc7f+IM1DA3LwTmctUomIu6V/BgQtBr2sfgQBIeEljn1G6tb8nXioxmC6Ns1cVc=' \
-H 'sec-ch-ua: "Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"' \
-H 'sec-ch-ua-mobile: ?1' \
-H 'sec-ch-ua-platform: "Android"' \
-H 'sec-fetch-dest: document' \
-H 'sec-fetch-mode: navigate' \
-H 'sec-fetch-site: none' \
-H 'sec-fetch-user: ?1' \
-H 'upgrade-insecure-requests: 1' \
-H 'user-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36' \
--compressed
func main() {
	initAll()
	//var headers2 = `Accept: text/html, application/xhtml+xml, */*`
	//   fmt.Printf("====== step 1：get Cookie ======")
	//   var MainUrl string = "https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf"
	//   fmt.Printf("\nMainUrl=%s", MainUrl)
	//   fmt.Println(getUrlRespHtmlWithHeader(MainUrl, headers2))
	//   dbgPrintCurCookies()

	fmt.Printf("\n\n\n====== step 2：get Cookie ======")
	var headers2 = `Accept: text/html, application/xhtml+xml, */*
Referer: www.lindress.com
Accept-Language: en-US
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko
Content-Type: application/x-www-form-urlencoded
Host: 192.168.132.80
Content-Length: 258
Pragma: no-cache
Cookie: client_id=1669700317931191; _c_id=1669700317931391753; _identity_cart=b421fa35-839e-4adf-89bc-88f96ffcc6d0; store_locale=en-US; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22184c1e4f7dc660-039093a1231da9-26021e51-3686400-184c1e4f7ddaad%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22%24device_id%22%3A%22184c1e4f7dc660-039093a1231da9-26021e51-3686400-184c1e4f7ddaad%22%7D; _fbp=fb.1.1669700321439.1375841203; _identity_popups=51d46e33-865a-4990-a32a-e83673d398611669700344; _identity_popups_bundle=955c9b47-96f7-451b-814a-f46c5da7861a1669700344; sw_session=6392da6e0a018; __cf_bm=Lru31KFw9EvFsV4tVr02HC23ocg9jjy0x2NahUHZPwc-1670568558-0-AfBlJQX/HwZ/T37pBOmEKXyD8c0hh5jSFBJHndm3aoNfqZHIi5r4L08qvZ4u8+/CHnofHjQjNy+9xsdqIZQHbFw=; session_id=1670568559249280`

	var formdata = `loginfile=%2Fwui%2Ftheme%2Fecology7%2Fpage%2Flogin.jsp%3FtemplateId%3D6%26logintype%3D1%26gopage%3D&logintype=1&fontName=%CE%A2%C8%ED%D1%C5%BA%DA&message=&gopage=&formmethod=post&rnd=&serial=&username=&isie=true&loginid=admin&userpassword=1234&submit=`
	var getapiUrl string = "https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf"
	PostUrlRespHtmlWithHeader(getapiUrl, headers2, formdata)
	dbgPrintCurCookies()

	fmt.Printf("\n\n\n====== step 3：use the Cookie ======")
	var headers3 = `Host: www.lindress.com
Pragma: no-cache
Cache-Control: no-cache
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Language: en-US,en;q=0.8`
	var getapiUrl3 string = "https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf"
	fmt.Println(getUrlRespHtmlWithHeader(getapiUrl3, headers3))
	dbgPrintCurCookies()
}