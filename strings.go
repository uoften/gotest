package main

import (
	"fmt"
	"strings"
)

func main(){
	//截取字符串
	str11 := "https://cloud.tencent.com/developer/article/2073969?from=article.detail.1432059&areaSource=106000.1&traceId=V3ffCGThCv0la7KT2_yqN"
	fmt.Println(str11[:strings.Index(str11,"?")])

	str111 := "75340 is your Facebook confirmation code"
	fmt.Println(str111[(strings.Index(str111,"is your Facebook")+len("is your Facebook")):])

	title := "987654321123"
	if len(title)>11{
		title = title[:11]
	}
	fmt.Println(title)
	fmt.Println(strings.TrimSpace("    aBc, ,    100a ")) //删除前后空格
	fmt.Println(strings.Replace("           aBc,  ,    100a","  "," ",2)) //替换第n个,-1为全部
	fmt.Println(strings.ReplaceAll("    aBc,  ,    100a","  "," ")) //全部替换

	s1 := " aBc"
	s2 := " 100a"
	s3 := s1 + s2
	fmt.Println(s3)
	fmt.Println(strings.HasPrefix(s3, "a")) //判断前缀
	fmt.Println(strings.HasSuffix(s3, "0")) //判断后缀
	fmt.Println(strings.Contains(s3, "9"))  //字符串包含关系
	fmt.Println(strings.Index(s3, "0"))     //判断子字符串或字符在父字符串中出现的位置（索引）
	fmt.Println(strings.LastIndex(s3, "4")) //最后出现位置的索引
	fmt.Println(strings.Replace(s3,"0","1",-1))//如果 n = -1 则替换所有字符串
	fmt.Println(strings.Count(s3,"0"))//出现的非重叠次数
	fmt.Println(strings.Repeat(s3,2))//重复字符串
	fmt.Println(strings.ToLower(s3))//修改字符串大小写
	fmt.Println(strings.ToUpper(s3))//修改字符串大小写
	fmt.Println(strings.TrimSpace(s3))//修剪字符串 去掉开头和结尾空格
	fmt.Println(strings.Trim(strings.TrimSpace(s3),"a"))//修剪字符串 去掉开头和结尾字符串
	fmt.Println(strings.Split(strings.TrimSpace(s3)," "))//分割字符串成数组
	arr := make([]string, 10)
	arr = []string{" a"," b"}
	fmt.Println(getArrayTrimSpace(arr))

	//获取字符串的某一段字符
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")
	fmt.Println(comma)

	fmt.Println(getTrimMoreByString("[100/2]"))

	asd:=strings.Split("any", " ")
	fmt.Printf("%#v\n",asd)

}

func getArrayTrimSpace(arr []string) (array []string) {
	for _,v:=range arr{
		array = append(array,strings.TrimSpace(v))
	}
	return array
}

func getTrimMoreByString(str string) (priority string) {
	left := strings.Index(str, "[")
	middle := strings.Index(str, "/")
	right := strings.Index(str, "[")
	if left!=1 && middle !=-1 &&  right !=-1 {
		return str[left+1:middle]
	}
	return "0"
}