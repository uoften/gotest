package main

import (
	"bytes"
	"fmt"
	"gotest/library"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pageStr, err := library.GetCloudflare("https://developer.aliyun.com/article/857040")
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`"larkContent": '(.*)',`)
	platformRegRegMatch := re.FindStringSubmatch(pageStr)
	str:= Stripslashes(platformRegRegMatch[0])
	v, _ := UnescapeUnicode([]byte(str))
	fmt.Println(string(v))
}

func Stripslashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}