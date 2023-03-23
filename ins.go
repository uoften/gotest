package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Instargam struct {
	Data struct {
		User struct {
			EdgeOwnerToTimelineMedia struct {
				Edges []struct {
					Node struct {
						DisplayURL            string `json:"display_url"`
						VideoURL              string `json:"video_url"`
						EdgeSidecarToChildren struct {
							Edges []struct {
								Node struct {
									DisplayURL string `json:"display_url"`
									VideoURL   string `json:"video_url"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_sidecar_to_children"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"data"`
}

func GetHtml(Insurl string) (html string) {

	// 解析代理地址
	proxy, err := url.Parse("http://127.0.0.1:1087")
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	request, err := http.NewRequest("GET", Insurl, nil)
	if err != nil {
		log.Println(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36") //模拟浏览器User-Agent
	res, err := httpClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	//判断是否成功访问，如果成功访问StatusCode应该为200
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	content, _ := ioutil.ReadAll(res.Body)
	return string(content)
}

func GetDownloadUrl(txt string, jsonurl string) {
	respHtml := GetHtml(jsonurl)
	var ins Instargam
	json.Unmarshal([]byte(respHtml), &ins)
	fmt.Println()
	for _, v := range ins.Data.User.EdgeOwnerToTimelineMedia.Edges {

		var content string
		if v.Node.DisplayURL != "" {
			fmt.Println(v.Node.DisplayURL)
			content = v.Node.DisplayURL + "\n"
		}
		if v.Node.VideoURL != "" {
			fmt.Println(v.Node.VideoURL)
			content += v.Node.VideoURL + "\n"
		}
		for _, v1 := range v.Node.EdgeSidecarToChildren.Edges {
			if v1.Node.VideoURL != "" {
				if v1.Node.DisplayURL != "" {
					fmt.Println(v1.Node.DisplayURL)
					content += v1.Node.DisplayURL + "\n"
				}
				if v1.Node.VideoURL != "" {
					fmt.Println(v1.Node.VideoURL)
					content += v1.Node.VideoURL + "\n"
				}
			}
		}
		WirteText(txt, content)
	}
}

func WirteText(savefile string, txt string) {
	f, err := os.OpenFile(savefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString(txt)
	bw.Flush()
}

func main() {
	var url string
	for ; ; {
		fmt.Println("输入地址：")
		fmt.Scanln(&url)
		url = strings.TrimSpace(url)
		GetDownloadUrl("save.txt", url)
	}

}