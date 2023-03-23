package library

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ClearStrSpace 去除换行符空格
func ClearStrSpace(str string) string{
	// 去除空格
	str = strings.TrimSpace(str)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func HttpPost(url string, data interface{}, header map[string]string,
	other map[string]int, otherStr map[string]string) (resData []byte, err error) {
	//超时时间
	var timeout int
	if value, ok := other["timeout"]; ok {
		timeout = value
	} else {
		timeout = 30
	}
	//
	//resData = make(map[string]interface{})
	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	//
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	request, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return
	}
	//Auth
	if otherStr["auth_user"] != "" && otherStr["auth_pass"] != "" {
		request.SetBasicAuth(otherStr["auth_user"], otherStr["auth_pass"])
	}
	//
	for key, value := range header {
		request.Header.Add(key, value)
	}
	cookie := &http.Cookie{Name: "", Value: ""}
	request.AddCookie(cookie)
	res, err := client.Do(request)
	if err != nil {
		return
	}
	//
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	//
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	//
	statusCode := res.StatusCode
	if statusCode != 200 {
		err = errors.New(fmt.Sprintf("调用接口失败：%d", statusCode))
		fmt.Print(err)
		return
	}
	resData = body
	//err = json.Unmarshal(body, &resData)
	return
}
