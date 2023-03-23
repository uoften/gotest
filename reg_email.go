package main

import (

	"fmt"
	"regexp"
)

const text = `
My email is 8899666@qq.com
email1 is abc@def.org
as@asd
32@asd
  $@@.as
#f@sdf
email2 is kkk@qq.com
email3 is ddd@abc.com.cn
asdasd
asdas@as.as
asd
asd
@.as
@as.asasd choul.zimathk2@outlook.com.cn      choul_zimathk@outlook.com        choulzim-athk1@outlook.com     emaleepajsik@outlook.com              choulzimathk2@outlook.com   
choulzimathk@outlook.com      choulzimathk1@outlook.com        emaleepajsik@outlook.com     choulzimathk2@outlook.com   

`

const reEmail = "(\\.|\\w|\\-)+@\\w+(\\.\\w+)+"
//const reEmail = "([a-zA-Z0-9\\.\\_]+)@([a-zA-Z0-9]+)(\\.[a-zA-Z0-9.]+)"
func main() {
	imageReg := regexp.MustCompile(reEmail)
	resList := imageReg.FindAllStringSubmatch(text, -1)
	for _,v:=range resList{
		if len(v)>0 {
			fmt.Println(v[0])
		}
	}

	////确定要寻找的目标及返回需要的字符段
	//re  := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	////返回二维数组 函数的作用是得到字符段并按要求返回需要的单个字符串
	//match := re.FindAllStringSubmatch(text,-1)
	////循环打印每一段字符
	//for _, m := range match{
	//	fmt.Println(m)
	//}
}