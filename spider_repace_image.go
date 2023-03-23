package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gotest/library"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const (
	folder = "D:\\home"
)
var (
	reImage    = `(https://|//)[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)|(webp)|(svg)))+(\S*?)" `
	//reUrl    = `https://[^"]+?.html`
)
// 获取当前页图片链接
func getImgs(pageStr string) (urls []string) {
	re := regexp.MustCompile(reImage)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}
func RepImages(htmls string) {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(htmls, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
		fmt.Println(strconv.Itoa(i), out[i])
	}
}
func main() {
	//pageData,_ := library.GetPageStr("https://www.cnblogs.com/itxiaoshen/p/16993369.html")
	pageData,_ := library.GetPageStr("https://blog.csdn.net/uxuepai5g/article/details/108570894")
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	articleContent,err := d.Find("#content_views").Html()
	if err != nil {
		fmt.Println(err)
	}
	RepImages(articleContent)
	imgUrls := getImgs(articleContent)
	for _,v:=range imgUrls{
		fmt.Println(v)
		//dateDir := time.Now().Format("2006010215")
		//err = DownloadImages(v,library.GetFilenameFromUrl(v),dateDir)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//articleContent = strings.Replace(articleContent, v, "//images.kaopuke.com/file_images/"+dateDir+`/`+library.GetFilenameFromUrl(v)+`" `, -1)
	}
	//fmt.Println(articleContent)
}

func DownloadImages(url string,fireName,dateDir string) error {
	baseDir := path.Join(folder,"/"+dateDir)
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		// 先创建文件夹
		os.Mkdir(baseDir, 0777)
		// 再修改权限
		os.Chmod(baseDir, 0777)
	}
	fireUrl := path.Join(folder,"/"+dateDir,"/",fireName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fireUrl)
	_ = ioutil.WriteFile(fireUrl, body, 0755)
	return nil
}