package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 下载图片信息
func downLoad(base string, url string) error {
	pic := base
	idx := strings.LastIndex(url, "/")
	if idx < 0 {
		pic += "/" + url
	} else {
		pic += url[idx+1:]
	}
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		return err
	}
	err = ioutil.WriteFile(pic, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return err
	}
	return nil
}

func main() {
	savePath := "./"
	url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	err := downLoad(savePath, url)
	if err != nil {
		fmt.Println("Download pic file failed!", err)
	} else {
		fmt.Println("Download file success.")
	}
	return
}