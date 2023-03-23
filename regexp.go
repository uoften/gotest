package main

import (
	"errors"
	"fmt"
	"github.com/dlclark/regexp2"
	"io/ioutil"
	"log"
	"regexp"
)
func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则MustCompile错误:" + err.Error())
	}
	return reg.ReplaceAllString(str, replace), nil
}

func main() {

	orgStr := `<p>直接点击<a href="https://www.yidaoerp.cn/user/login" target="_blank">立即注册</a>登陆使用</p>`
	dstStr, err := ReplaceStringByRegex(orgStr, "<[^a>]+>", "")
	log.Println(dstStr, err)

	submatchArr := regexp.MustCompile(`toutiao.weaoo.com/detail/(\d+)_\d+.html`).FindStringSubmatch("https://toutiao.weaoo.com/detail/1080888_1.html")
	newsId := submatchArr[1]
	fmt.Print(newsId)

	config_bytes, err := ioutil.ReadFile("../device-agent-go/logs/20220513/10.54.5.167_switch.txt")
	if err != nil {
		fmt.Printf("读取对应的配置文件失败")
		return
	}
	configStr := string(config_bytes)
	reg, _ := regexp2.Compile("(.*)---- More ----", 0)
	m, _ := reg.FindStringMatch(configStr)
	if m != nil {
		res:= m.String()
		regcmd := regexp.MustCompile(res)
		result := regcmd.ReplaceAllString(configStr, "")
		fmt.Println(result)
	}else{
		fmt.Println("asdasd")
	}
}