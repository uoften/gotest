package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gotest/common/helper"
	"strings"
)

func main()  {
	//md5加密
	userInfo := map[string]string{"name":"asd"}
	str,_ := json.Marshal(userInfo)
	fmt.Println(string(str))
	md5Str := helper.Md5(string(str))
	fmt.Println(md5Str)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("【批量下载】SMART_DP01_V01.00.01_00.00.00.00等")))
	strBase,err:=base64.StdEncoding.DecodeString("Y2tlZGl0b3JfNC4xNy4xXzVmY2NkY2Q1ZTdiZS56aXA=")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(strBase))


	//自定义加解密：base64+混淆
	strMap := map[string]string{
		"exptime":"asdasd",
		"uid":"asd",
	}
	fmt.Printf("加密前：%#v\n",strMap)
	key := "uoasdasdasfteasdasdn"
	fmt.Printf("密钥：%#v\n",key)
	//加密后的字符串
	fmt.Printf("加密后的字符串：%v\n",baseEncode(strMap,key))
	//解密后的字符串
	strdata := baseDecode("euyoJalsedHaBs0daaWs1fltIejaosidMaTsYd2nNDI0NjA5NiIsImZpbGVuYW1lIjoi44CQ5om56YeP5LiL6L2944CRU01BUlRfRFAwMV9WMDEuMDAuMDFfMDAuMDAuMDAuMDDnrYkoNCkuemlwIn0O0O0O","uoasdasdasfteasdasdn")
	fmt.Printf("解密后的json字符串：%v\n",strdata)
	start := strings.Index(strdata, "{")
	end := strings.Index(strdata, "}")
	jsonStr := strdata[start:end+1]
	fmt.Println(jsonStr)
	//fileInfo := FlieInfo{}
	//err = json.Unmarshal([]byte(jsonStr), &fileInfo)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func baseEncode(jsonMap map[string]string,key string) string {
	jsonStr,_ := json.Marshal(jsonMap)
	baseEncodedStr := base64.StdEncoding.EncodeToString(jsonStr)
	baseEncodedByte := []byte(baseEncodedStr)
	var retStr string
	for kk,vv:=range []byte(key){
		retStr = retStr+string(baseEncodedByte[kk])+string(vv)
	}
	baseEncodedByteLast := string(baseEncodedByte[len([]byte(key)):])
	return retStr+baseEncodedByteLast
}

func baseDecode(baseEncodedStr string,key string) string {
	var retStr string
	baseEncodedByte := []byte(baseEncodedStr)
	ii:=0
	for _,_=range []byte(key){
		retStr = retStr + string(baseEncodedByte[ii:ii+1])
		ii=ii+2
	}
	baseEncodedByteLast := string(baseEncodedByte[ii:])
	baseDecodedByte,_ := base64.StdEncoding.DecodeString(retStr+baseEncodedByteLast)
	return string(baseDecodedByte)
}