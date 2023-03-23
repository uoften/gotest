package main

import (
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

func main() {

	client, _ := cloudscraper.Init(false, false)
	//res, _ := client.Post("https://www.lindress.com/collections/mens-coats-jackets/products/leather-and-fur-vest-imtw-ovqg-k750-1hz2-5bx1-r7jr-zjsc-dmva-lq55-1flt-ppcs-jnc1-3enf", make(map[string]string), "")
	res2, _ := client.Get("http://t.zoukankan.com/liangxiaofeng-p-5109889.html", make(map[string]string), "")
	//print(res.Body)
	print("-------------------------------------------------")
	print(res2.Body)
}