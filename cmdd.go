package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

//go build -ldflags="-s -w -H windowsgui" -o offwin10.exe cmdd.go
func main() {
	resData := getCode()
	if resData["code"] == "S" {
		err := command("shutdown -s -t 10")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCode() (resData map[string]interface{}) {
	//get方式
	resp,err:=httpGet("/index/index/closewin1")
	if err != nil {
		resp,err=httpGet("/index/index/closewin1")
		if err != nil {
			resp,err=httpGet("/index/index/closewin1")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	resData = resp
	return

	//post方式
	//resp,err:=PostGetCode("/api/index/offwin10",nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//resData = resp
	//return
}

// 这里为了简化，我省去了stderr和其他信息
func command(cmd string) (err error) {
	//c := exec.Command("bash", "-c", cmd)
	//此处是windows版本
	//c := exec.Command("cmd", "/C", cmd)
	//err = c.Run()
	c := exec.Command("cmd", "/C", cmd)
	if runtime.GOOS == "windows" {
		c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	err = c.Run()
	return
}

func httpGet(url string) (resData map[string]interface{}, err error) {
	// get方式
	url = fmt.Sprintf("%s%s", "https://kaopuke.com/", url)
	resp, err := http.Get(url)
	// 关闭 resp.Body 的正确姿势
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resData = make(map[string]interface{})
	err = json.Unmarshal(body, &resData)
	if err != nil {
		body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
		err = json.Unmarshal(body, &resData)
		if err != nil {
			err = errors.New(fmt.Sprintf("调用接口%s失败：%v", url, err))
			return
		}
	}
	return
}

func PostGetCode(url string, data map[string]interface{}) (resData map[string]interface{}, err error) {
	url = fmt.Sprintf("%s%s", "http://127.0.0.1/", url)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	resData, err = HttpPost(url, data, header, map[string]int{}, map[string]string{})
	if err != nil {
		return
	}
	if resData["code"] != "S" {
		err = errors.New(resData["msg"].(string))
	}
	return
}

func HttpPost(url string, data interface{}, header map[string]string,
	other map[string]int, otherStr map[string]string) (resData map[string]interface{}, err error) {
	//超时时间
	var timeout int
	if value, ok := other["timeout"]; ok {
		timeout = value
	} else {
		timeout = 60
	}
	//
	resData = make(map[string]interface{})
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
		err = errors.New(fmt.Sprintf("调用接口%s失败：%d", url, statusCode))
		fmt.Print(err)
		return
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
		err = json.Unmarshal(body, &resData)
		if err != nil {
			return nil,err
		}
	}
	return
}