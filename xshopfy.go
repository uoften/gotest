package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type Product struct {
	PostTitle string `json:"post_title"`
	Sku string `json:"sku"`
	Price interface{} `json:"price"`
	RegularPrice interface{} `json:"regular_price"`
	SalePrice interface{} `json:"sale_price"`
	MinPrice interface{} `json:"min_price"`
	MaxPrice interface{} `json:"max_price"`
	Image string `json:"image"`
	Gallery []GalleryData `json:"gallery"`
	variants [][]SkuAttrs
	PostContent string `json:"post_content"`
}

type GalleryData struct {
	ID int `json:"ID"`
	MediaContentType string `json:"media_content_type"`
	Thumbnail string `json:"thumbnail"`
	Url string `json:"url"`
}

type SkuAttrs struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Code string `json:"code"`
}

// ClearStrSpace 去除换行符空格
func clearStrSpace(str string) string{
	// 去除空格
	str = strings.TrimSpace(str)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func main() {
	pageStr := `

<link rel="stylesheet" href="https://static.hwshopy.com/liquid/buyer/public/css/styles.min.css?t=20221130103442">

<link rel="stylesheet"
    href="https://static.hwshopy.com/liquid/buyer/public/css/brooklyn.comm.min.css?t=20221130103442">

<script type="text/javascript" src="https://static.hwshopy.com/liquid/buyer/public/js/init.js"></script>
<script src="https://static.hwshopy.com/liquid/buyer/public/js/plug/jquery.js"></script>
    <script defer="defer" src="https://www.paypalobjects.com/api/checkout.min.js"></script>
    <!-- <script src="https://www.paypal.com/sdk/js?components=buttons"></script> -->

<script type="text/javascript" src="https://static.hwshopy.com/liquid/buyer/public/js/plug/Swiper.js"></script>
<script type="text/javascript" src="https://static.hwshopy.com/liquid/buyer/public/js/plug/fingerprint2.min.js"></script>
<script type="text/javascript" src="https://static.hwshopy.com/liquid/buyer/public/js/plug/vendor.min.js?t=20221130103442"></script>
`

	regPageStr := regexp.MustCompile("data-product-variants")
	pageStr = regPageStr.ReplaceAllString(pageStr, "class=\"data-product-variants\"")
	fmt.Println(pageStr)
}

func main11() {
	pageStr := `
<div data-product-variants>
            <div class="types-item J-Attribute">
                <p class="font14">
                    <span class="J-Type" data-type="Color">
                        Color
                    </span>
                    <span class="types-item-size J-TypeVal" data-type-val="sliver">
                        sliver
                    </span>
                </p>
                <ul class="product-types-list J-Value clearfloat">
                    
                    
                        <li data-title="sliver" data-title-code="c2xpdmVy"
                            class="product-types-item J-ProductTypes
                            
                                active
                            
                            
                            np-ui-subhead-title"
                        >
                            sliver
                        </li>
                    
                        <li data-title="black" data-title-code="YmxhY2s"
                            class="product-types-item J-ProductTypes
                            
                            
                            np-ui-subhead-title"
                        >
                            black
                        </li>
                    
                        <li data-title="beige" data-title-code="YmVpZ2U"
                            class="product-types-item J-ProductTypes
                            
                            
                            np-ui-subhead-title"
                        >
                            beige
                        </li>
                    
                </ul>
            </div>
            <div class="types-item J-Attribute">
                <p class="font14">
                    <span class="J-Type" data-type="Package includes">
                        Package includes
                    </span>
                    <span class="types-item-size J-TypeVal" data-type-val="1 pcs">
                        1 pcs
                    </span>
                </p>
                <ul class="product-types-list J-Value clearfloat">
                        <li data-title="1 pcs" data-title-code="MSBwY3M"
                            class="product-types-item J-ProductTypes
                                active
                            np-ui-subhead-title"
                        >
                            1 pcs
                        </li>
                        <li data-title="2 pcs(save 4 usd)" data-title-code="MiBwY3Moc2F2ZSA0IHVzZCk"
                            class="product-types-item J-ProductTypes
                            
                            
                            np-ui-subhead-title"
                        >
                            2 pcs(save 4 usd)
                        </li>
                        <li data-title="4 pcs( Best prices)" data-title-code="NCBwY3MoIEJlc3QgcHJpY2VzKQ"
                            class="product-types-item J-ProductTypes
                            np-ui-subhead-title"
                        >
                            4 pcs( Best prices)
                        </li>
                </ul>
            </div>
</div>
	`
	var skus [][]SkuAttrs
	regPageStr := regexp.MustCompile("data-product-variants")
	pageStr = regPageStr.ReplaceAllString(pageStr, "class=\"data-product-variants\"")
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	tt := d.Find(".data-product-variants .J-Attribute")
	for i := 0; i < tt.Length(); i++ {
		var sku []SkuAttrs
		var attr SkuAttrs
		tStr,_ := tt.Eq(i).Html()
		dKey, _ := goquery.NewDocumentFromReader(strings.NewReader(tStr))
		attrKey,_ := dKey.Find(".J-Type").Attr("data-type")

		dd, _ := goquery.NewDocumentFromReader(strings.NewReader(tStr))
		t := dd.Find("option")
		if t.Text()!="" {
			for i := 0; i < t.Length(); i++ {
				attr.Name = attrKey
				attrValue := t.Eq(i).Text()
				attr.Value = clearStrSpace(attrValue)
				attr.Code,_ = t.Eq(i).Attr("data-title-code")
				sku = append(sku,attr)
			}
		}else{
			t = dd.Find("li")
			for i := 0; i < t.Length(); i++ {
				attr.Name = attrKey
				attr.Value,_ = t.Eq(i).Attr("data-title")
				attr.Code,_ = t.Eq(i).Attr("data-title-code")
				sku = append(sku,attr)
			}
		}
		v := reflect.ValueOf(sku)
		fmt.Println(v.Interface().(interface{}))
		skus = append(skus,sku)
	}
}

//节点匹配获取属性对应价格和图片
func main3() (){
	pageStr := `
<div class="asd">
<div style="display: none" class="J-SkuList">
    
        <span data-image-url="https://img.cdncloud.top/uploader/c5e22c96b023635ba588571da7b9883a3b6b0bcc.jpg"
            banner-id="125862198"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="c2xpdmVyMSBwY3M"
            data-id="263945018"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/c5e22c96b023635ba588571da7b9883a3b6b0bcc.jpg"
            banner-id="125862198"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="c2xpdmVyMiBwY3Moc2F2ZSA0IHVzZCk"
            data-id="263945019"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/c5e22c96b023635ba588571da7b9883a3b6b0bcc.jpg"
            banner-id="125862198"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="c2xpdmVyNCBwY3MoIEJlc3QgcHJpY2VzKQ"
            data-id="263945020"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/a15082a33a4c9c92f14dea9afcded43beaf7c711.jpg"
            banner-id="125862200"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmxhY2sMSBwY3M"
            data-id="263945021"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/a15082a33a4c9c92f14dea9afcded43beaf7c711.jpg"
            banner-id="125862200"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmxhY2sMiBwY3Moc2F2ZSA0IHVzZCk"
            data-id="263945022"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/a15082a33a4c9c92f14dea9afcded43beaf7c711.jpg"
            banner-id="125862200"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmxhY2sNCBwY3MoIEJlc3QgcHJpY2VzKQ"
            data-id="263945023"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/2ce81931363f42f366d589f27cec3c46fb0e494a.jpg"
            banner-id="125862213"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmVpZ2UMSBwY3M"
            data-id="263945024"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/2ce81931363f42f366d589f27cec3c46fb0e494a.jpg"
            banner-id="125862213"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmVpZ2UMiBwY3Moc2F2ZSA0IHVzZCk"
            data-id="263945025"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    
        <span data-image-url="https://img.cdncloud.top/uploader/2ce81931363f42f366d589f27cec3c46fb0e494a.jpg"
            banner-id="125862213"
            data-compare-at-price="0"
            data-orig-compare-at-price="0.00"
            data-inventory-quantity="100000"
            data-price="0"
            data-orig-price="0.00"
            data-title-code="YmVpZ2UNCBwY3MoIEJlc3QgcHJpY2VzKQ"
            data-id="263945026"
            data-limitoffer-open=""
            data-limitoffer-num="0"
        >
        </span>
    </div>
</div>
`

	dd, _ := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	t,_ := dd.Find(".J-SkuList").Html()
	if t!="nil" {
		tt := dd.Find(".J-SkuList").Find("span")
		for i := 0; i < tt.Length(); i++ {
			imageUrl,_ := tt.Eq(i).Attr("data-image-url")
			keyCode,_ := tt.Eq(i).Attr("data-title-code")
			price,_ := tt.Eq(i).Attr("data-price")
			origPrice,_ := tt.Eq(i).Attr("data-orig-price")
			fmt.Println(keyCode)
			fmt.Println(imageUrl)
			fmt.Println(price)
			fmt.Println(origPrice)
		}
	}
	return
}

func findAttr(d *goquery.Document) (skus [][]SkuAttrs){
	tt := d.Find(".data-product-variants .J-Attribute")
	for i := 0; i < tt.Length(); i++ {
		var sku []SkuAttrs
		var attr SkuAttrs
		tStr,_ := tt.Eq(i).Html()
		dKey, _ := goquery.NewDocumentFromReader(strings.NewReader(tStr))
		attrKey,_ := dKey.Find(".J-Type").Attr("data-type")

		dd, _ := goquery.NewDocumentFromReader(strings.NewReader(tStr))
		t := dd.Find("option")
		if t.Text()!="" {
			tCode,_ := dd.Find("option").Attr("data-title-code")
			for i := 0; i < t.Length(); i++ {
				attrValue := t.Eq(i).Text()
				attr.Name = attrKey
				attr.Value = clearStrSpace(attrValue)
				attr.Code = tCode

				sku = append(sku,attr)

			}
		}else{
			t = dd.Find("li")
			for i := 0; i < t.Length(); i++ {
				attr.Name = attrKey
				attr.Value,_ = t.Eq(i).Attr("data-title")
				attr.Code,_ = t.Eq(i).Attr("data-title-code")

				sku = append(sku,attr)

			}
		}
		skus = append(skus,sku)
	}
	return
}


func findContent(d *goquery.Document) string {
	res, _ := d.Find(".slot-mobile .product-details").Html()
	return res
}

func FindDataByXshoppy(pageStr string) (currency,product string){
	//正则匹配获取当前货币
	currencyReg := regexp.MustCompile(";Shopline.currency=\"(.*)\";Shopline.themeId=\"")
	currencyRegMatch := currencyReg.FindStringSubmatch(pageStr)
	if len(currencyRegMatch) > 0 {
		currency = currencyRegMatch[1]
	}
	regPageStr := regexp.MustCompile("data-product-variants")
	pageStr = regPageStr.ReplaceAllString(pageStr, "class=\"data-product-variants\"")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		log.Fatal(err)
	}
	findAttr(doc)
	//var productData Product
	//var skus VariantsData
	//skus.Attrs = findAttr(doc)
	//productData.Variants = append(productData.Variants,skus)
	//fmt.Printf("%#v\n",productData.Variants)

	//正则匹配获取商品信息
	//productReg := regexp.MustCompile("__PRELOAD_STATE__.product=(.*);</script>")
	//productMatchData := productReg.FindStringSubmatch(pageStr)
	//var productMatchStr string
	//if len(productMatchData) > 0 {
	//	product = productMatchStr
	//}
	return
}

