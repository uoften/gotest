//package main
//
//import (
//	"github.com/Danny-Dasilva/CycleTLS/cycletls"
//	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
//)
//
//func main() {
//
//	client, _ := cloudscraper.Init(false, false)
//	options := cycletls.Options{
//		Headers:         map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"},
//		//Body:            "",
//		//Proxy:           "proxy.company.com",
//		Timeout:         20,
//		DisableRedirect: true,
//	}
//	res, _ := client.Do("https://www.koulb.com/products/men's-t-shirt-tee-hot-stamping-graphic-letter-crew-neck-casual-daily-print-short-sleeve-tops-lightweight-fashion-muscle-big-and-tall-white-black-gray-summer-1", options, "PUT")
//	print(res.Body)
//}

package main

import (
"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

func main() {

	client, _ := cloudscraper.Init(false, false)
	//res, _ := client.Post("https://www.koulb.com/products/men's-t-shirt-tee-hot-stamping-graphic-letter-crew-neck-casual-daily-print-short-sleeve-tops-lightweight-fashion-muscle-big-and-tall-white-black-gray-summer-1", make(map[string]string), "")
	res2, _ := client.Get("https://www.baidu.com/", make(map[string]string), "")
	//print(res.Body)
	print(res2.Body)
}