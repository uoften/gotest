package main

import (
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var client = http.Client{Timeout: time.Second * 180}

var threadGroup = sync.WaitGroup{}

var packageSize int64

func init() {
	//每个线程下载文件的大小
	packageSize = 1048576 * 4
}

func Download(url, cachePath string, scheduleCallback func(schedule float64)) string {

	var localFileSize int64
	var file *os.File
	if info, e := os.Stat(cachePath); e != nil {
		if os.IsNotExist(e) {
			if createFile, err := os.Create(cachePath); err == nil {
				file = createFile
			} else {
				panic(err)
			}
		} else {
			panic(e)
		}
	} else {
		localFileSize = info.Size()
	}
	//HEAD 方法请求服务端是否支持多线程下载,并获取文件大小
	if request, e := http.NewRequest("HEAD", url, nil); e == nil {
		if response, i := client.Do(request); i == nil {
			defer response.Body.Close()
			//得到文件大小
			ContentLength := response.ContentLength
			if localFileSize == ContentLength {
				log.Warn("file exist~")
				return cachePath
			} else {
				//判断是否支持多线下载
				if strings.Compare(response.Header.Get("Accept-Ranges"), "bytes") == 0 {
					//支持 走下载流程
					if dispSliceDownload(file, ContentLength, url, scheduleCallback) == 0 {
						return cachePath
					} else {
						return ""
					}
				} else {
					panic("nonsupport ~")
				}
			}
		} else {
			panic(i)
		}
	} else {
		panic(e)
	}
	return ""
}

func dispSliceDownload(file *os.File, ContentLength int64, url string, scheduleCallback func(schedule float64)) int {
	defer file.Close()
	//文件总大小除以 每个线程下载的大小
	i := ContentLength / packageSize
	//保证文件下载完整
	if ContentLength%packageSize > 0 {
		i += 1
	}
	//下载总进度
	var schedule int64
	//分配下载线程
	for count := 0; count < int(i); count++ {
		//计算每个线程下载的区间,起始位置
		var start int64
		var end int64
		start = int64(int64(count) * packageSize)
		end = start + packageSize
		if int64(end) > ContentLength {
			end = end - (end - ContentLength)
		}
		//构建请求
		if req, e := http.NewRequest("GET", url, nil); e == nil {
			req.Header.Set(
				"Range",
				"bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10))
			//
			threadGroup.Add(1)
			go sliceDownload(req, file, &schedule, &ContentLength, scheduleCallback, start)
		} else {
			panic(e)
		}

	}
	//等待所有线程完成下载
	threadGroup.Wait()
	return 0
}

func sliceDownload(request *http.Request, file *os.File, schedule *int64, ContentLength *int64, scheduleCallback func(schedule float64),
	start int64) {
	defer threadGroup.Done()
	if response, e := client.Do(request); e == nil && response.StatusCode == 206 {
		defer response.Body.Close()
		if bytes, i := ioutil.ReadAll(response.Body); i == nil {
			i2 := len(bytes)
			//从我们计算好的起点写入文件
			file.WriteAt(bytes, start)
			atomic.AddInt64(schedule, int64(i2))
			val := atomic.LoadInt64(schedule)
			num := float64(val*1.0) / float64(*ContentLength) * 100
			scheduleCallback(float64(num))
		} else {
			panic(e)
		}
	} else {
		panic(e)
	}
}