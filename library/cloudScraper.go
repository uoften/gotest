package library

import (
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

//绕过CloudFlare安全检查
func GetCloudflare(url string) (pageStr string,err error) {
	client, _ := cloudscraper.Init(false, false)
	header := make(map[string]string)
	res, err := client.Get(url, header, "")
	if err != nil {
		return "error",err
	}
	pageStr = res.Body
	return pageStr,nil
}

//绕过CloudFlare安全检查
func PostCloudflare(url,body string) (pageStr string,err error) {
	client, _ := cloudscraper.Init(false, false)
	header := make(map[string]string)
	res, err := client.Post(url, header, body)
	if err != nil {
		return "error",err
	}
	pageStr = res.Body
	return pageStr,nil
}