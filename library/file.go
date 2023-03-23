package library

import (
	"net/url"
	"os"
	"path"
	"strings"
)

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// UriFilterExcludeQueryString 过滤url中的参数,并补齐https:
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

// GetFilenameFromUrl 截取url中的图片名
func GetFilenameFromUrl(uri string) (filename string) {
	clearUri := UriFilterExcludeQueryString(uri)
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(clearUri, "/")
	// 切出来
	filename = clearUri[lastIndex+1:]
	// 纳秒时间戳解决重名(出现图片重名时的方法)
	//timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	//filename = timePrefix + "_" + filename

	return
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