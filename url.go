package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func main() {
	url1 := "//cdn.shoplazza.com/3915e6d02eaef4d234cc6370a37593fe.png?with=1024&asd=1212"
	url := UriFilterExcludeQueryString(url1)

	fmt.Println(url)

	fmt.Println(GetFilenameFromUrl(url))

	fmt.Println(GetFileNameFromUrlNoSuffix(url))

	fmt.Println(GetFileSuffixFromUrl(url))
}

// 截取url中的图片后缀
func GetFileSuffixFromUrl(url string) string{
	//获取文件名称带后缀
	fileNameWithSuffix:=path.Base(url)
	//获取文件的后缀(文件类型)
	fileType:=path.Ext(fileNameWithSuffix)
	return fileType
}

// 截取url中的图片名不带后缀
func GetFileNameFromUrlNoSuffix(url string) string{
	//获取文件名称带后缀
	fileNameWithSuffix:=path.Base(url)
	//获取文件的后缀(文件类型)
	fileType:=path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	fileNameOnly:=strings.TrimSuffix(fileNameWithSuffix, fileType)
	return fileNameOnly
}

// 截取url中的图片名
func GetFilenameFromUrl(uri string) (filename string) {
	URL, _ := url.Parse(uri)

	clearUri := strings.ReplaceAll(uri, URL.RawQuery, "")

	clearUri = strings.TrimRight(clearUri, "?")

	clearUri = strings.TrimRight(clearUri, "/")
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(clearUri, "/")
	// 切出来
	filename = clearUri[lastIndex+1:]
	// 纳秒时间戳解决重名(出现图片重名时的方法)
	//timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	//filename = timePrefix + "_" + filename

	return
}

func UriFilterExcludeQueryString(uri string) string {
	URL, _ := url.Parse(uri)

	clearUri := strings.ReplaceAll(uri, URL.RawQuery, "")

	clearUri = strings.TrimRight(clearUri, "?")

	clearUri = strings.TrimRight(clearUri, "/")

	if !strings.HasPrefix(clearUri, "https://") {
		clearUri = "https:"+clearUri
	}
	return clearUri
}
