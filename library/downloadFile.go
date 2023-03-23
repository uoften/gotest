package library

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	thread = 32
	folder = "D:\\home\\download"
	timeout = 300
)
var (
	fileTypeMap sync.Map
	rLock       sync.Mutex
)


func init() { //用于判断文件名的后缀
	fileTypeMap.Store("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif")  //GIF (gif)
	fileTypeMap.Store("49492a00227105008037", "tif")  //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	fileTypeMap.Store("41433130313500000000", "dwg")  //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", "htm")  //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", "css")  //css
	fileTypeMap.Store("696b2e71623d696b2e71", "js")   //js
	fileTypeMap.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	fileTypeMap.Store("38425053000100000000", "psd")  //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", "mdb")  //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", "ps")
	fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706d70", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("000001ba210001000180", "mpg") //
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("4d546864000000060001", "mid") //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("235468697320636f6e66", "ini")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")        //可执行文件
	fileTypeMap.Store("3c25402070616765206c", "jsp")        //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", "mf")         //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", "xml")        //xml文件
	fileTypeMap.Store("494e5345525420494e54", "sql")        //xml文件
	fileTypeMap.Store("7061636b616765207765", "java")       //java文件
	fileTypeMap.Store("406563686f206f66660d", "bat")        //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")         //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", "class")      //bat文件
	fileTypeMap.Store("49545346030000006000", "chm")        //bat文件
	fileTypeMap.Store("04000000010000001300", "mxp")        //bat文件
	fileTypeMap.Store("504b0304140006000800", "docx")       //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov")         //Quicktime (mov)
	fileTypeMap.Store("FF575043", "wpd")         //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", "pst")         //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	fileTypeMap.Store("E3828596", "pwl")         //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", "ram")         //Real Audio (ram)
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

func SafeMkdir(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, os.ModePerm)
	}
}

func GetBytesFile(filename string, bufferSize int) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer
}

type DownloadTask struct {
	customFunc func(params interface{}) // 执行方法
	paramsInfo interface{}              // 执行方法参数
}

type DownloadThreadController struct {
	TaskQueue              chan DownloadTask       // 用于接收下载任务
	TaskCount              chan int                // 用于记载当前任务数量
	Exit                   chan int                // 用于记载当前任务数量
	ThreadCount            int                     // 最大协程数
	WaitGroup              sync.WaitGroup          // 等待协程完成
	RangeStrs              map[int]string          // 所有需要下载的文件名
	FileUrl                string                  // 下载链接
	DownloadResultInfoChan chan DownloadFileParams // 下载任务响应通道
	DownloadFolder         string                  // 下载文件保存文件夹
	DownloadFileName       string                  // 下载文件保存文件名
	Filenames              []string                // 子文件名，有序
}

type DownloadFileParams struct {
	UrlStr       string
	RangeStr     string
	RangeIndex   int
	TempFilename string
	Successed    bool
}

func (controller *DownloadThreadController) Put(task DownloadTask) {
	// 用于开启单个协程任务，下载文件的部分内容
	defer func() {
		err := recover() //内置函数，可以捕捉到函数异常
		if err != nil {
			fmt.Println("Chnnel closed", err)
		}
	}()
	controller.WaitGroup.Add(1)  // 每插入一个任务，就需要计数
	controller.TaskCount <- 1    // 含缓冲区的通道，用于控制下载器的协程最大数量
	controller.TaskQueue <- task // 插入下载任务
	//go task.customFunc(task.paramsInfo)
}

func (controller *DownloadThreadController) DownloadFile(paramsInfo interface{}) {
	// 下载任务，接收对应的参数，负责从网页中下载对应部分的文件资源
	defer func() {
		controller.WaitGroup.Done() // 下载任务完成，协程结束
	}()
	switch paramsInfo.(type) {
	case DownloadFileParams:
		params := paramsInfo.(DownloadFileParams)
		params.Successed = false
		defer func() {
			err := recover() //内置函数，可以捕捉到函数异常
			if err != nil {
				// 如果任意环节出错，表明下载流程未成功完成，标记下载失败
				fmt.Println("Failed to download file：", err)
				params.Successed = false
			}
		}()
		//fmt.Println("Start to down load " + params.UrlStr + ", Content-type: " + params.RangeStr + " , save to file: " + params.TempFilename)
		urlStr := params.UrlStr
		rangeStr := params.RangeStr
		tempFilename := params.TempFilename
		os.Remove(tempFilename) // 删除已有的文件, 避免下载的数据被污染
		// 发起文件下载请求
		req, _ := http.NewRequest("GET", urlStr, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Add("Accept-Encoding", "gzip, deflate")
		req.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Host", "www.downxia.com")
		req.Header.Add("Referer", urlStr)
		req.Header.Add("Range", rangeStr)      // 测试下载部分内容
		res, err := http.DefaultClient.Do(req) // 发出下载请求，等待回应
		if err != nil {
			fmt.Println("Failed to connet " + urlStr)
			params.Successed = false // 无法连接, 标记下载失败
		} else if res.StatusCode != 206 {
			params.Successed = false
		} else { // 能正常发起请求
			// 打开文件，写入文件
			fileObj, err := os.OpenFile(tempFilename, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Failed to open file " + tempFilename)
				params.Successed = false // 无法打开文件, 标记下载失败
			} else {
				defer fileObj.Close()                 // 关闭文件流
				body, err := ioutil.ReadAll(res.Body) // 读取响应体的所有内容
				if err != nil {
					fmt.Println("Failed to read response body.")
					params.Successed = false
				} else {
					defer res.Body.Close()  // 关闭连接流
					fileObj.Write(body)     // 写入字节数据到文件
					params.Successed = true // 成功执行到最后一步，则表示下载成功
				}
			}
		}
		controller.DownloadResultInfoChan <- params // 将下载结果传入
	}
}

func (controller *DownloadThreadController) Run() {
	// 只需要将待下载的请求发送一次即可，成功了会直接剔除，不成功则由接收方重试
	for rangeIndex, rangeStr := range controller.RangeStrs {
		params := DownloadFileParams{UrlStr: controller.FileUrl, RangeStr: rangeStr, TempFilename: controller.DownloadFolder + "/" + rangeStr, RangeIndex: rangeIndex, Successed: true} // 下载参数初始化
		task := DownloadTask{controller.DownloadFile, params}
		controller.Put(task) // 若通道满了会阻塞，等待空闲时再下载
	}
}

func (controller DownloadThreadController) GetSuffix(contentType string) string {
	suffix := ""
	contentTypes := map[string]string{
		"image/gif":                    "gif",
		"image/jpeg":                   "jpg",
		"application/x-img":            "img",
		"image/png":                    "png",
		"application/json":             "json",
		"application/pdf":              "pdf",
		"application/msword":           "word",
		"application/octet-stream":     "rar",
		"application/x-zip-compressed": "zip",
		"application/x-msdownload":     "exe",
		"video/mpeg4":                  "mp4",
		"video/avi":                    "avi",
		"audio/mp3":                    "mp3",
		"text/css":                     "css",
		"application/x-javascript":     "js",
		"application/vnd.android.package-archive": "apk",
	}
	for key, value := range contentTypes {
		if strings.Contains(contentType, key) {
			suffix = value
			break
		}
	}
	return suffix
}

func (controller *DownloadThreadController) ResultProccess(trunkSize int) string {
	// 负责处理各个协程下载资源的结果， 若成功则从下载列表中剔除，否则重新将该任务Put到任务列表中；超过5秒便会停止
	MAX_RETRY_TIME := 100
	nowRetryTime := 0
	result_msg := ""
	for {
		select {
		case resultInfo := <-controller.DownloadResultInfoChan:
			<-controller.TaskCount    // 取出一个计数器，表示一个协程已经完成
			if resultInfo.Successed { // 成功下载该文件，清除文件名列表中的信息
				rLock.Lock()
				delete(controller.RangeStrs, resultInfo.RangeIndex) // 删除任务队列中的该任务（rangeStr队列）
				rLock.Unlock()
				fmt.Println("Progress -> " + strconv.FormatFloat((1.0-float64(len(controller.RangeStrs))/float64(trunkSize))*100, 'f', 2, 64) + "%")
				if len(controller.RangeStrs) == 0 {
					result_msg = "SUCCESSED"
					break
				}
			} else {
				nowRetryTime += 1
				if nowRetryTime > MAX_RETRY_TIME { // 超过最大的重试次数退出下载
					result_msg = "MAX_RETRY"
					break
				}
				task := DownloadTask{customFunc: controller.DownloadFile, paramsInfo: resultInfo} // 重新加载该任务
				go controller.Put(task)
			}
		case task := <-controller.TaskQueue:
			function := task.customFunc
			go function(task.paramsInfo)
		case <-time.After(timeout * time.Second):
			result_msg = "TIMEOUT"
			break
		}
		if result_msg == "MAX_RETRY" {
			fmt.Println("The network is unstable, exceeding the maximum number of redownloads.")
			break
		} else if result_msg == "SUCCESSED" {
			fmt.Println("Download file successeed!")
			break
		} else if result_msg == "TIMEOUT" {
			fmt.Println("Download timeout!")
			break
		}
	}

	close(controller.TaskCount)
	close(controller.TaskQueue)
	close(controller.DownloadResultInfoChan)
	return result_msg
}

func (controller *DownloadThreadController) Download(oneThreadDownloadSize int) bool {
	fmt.Println("Try to parse the object file...")
	length, rangeMaps, tempFilenames, contentType, err := TestDownload(controller.FileUrl, oneThreadDownloadSize)
	fmt.Println("File total size -> " + strconv.FormatFloat(float64(length)/(1024.0*1024.0), 'f', 2, 64) + "M")
	if err != nil {
		fmt.Printf("The file does not support multi-threaded download:%v\n",err)
		return false
	}
	fmt.Println("Parse the target file successfully, start downloading the target file...")
	controller.Init() // 初始化通道、分片等配置
	//oneThreadDownloadSize := 1024 * 1024 * 2 // 1024字节 = 1024bite = 1kb -> 2M
	//oneThreadDownloadSize = 1024 * 1024 * 4 // 1024字节 = 1024bite = 1kb -> 4M
	if length<1024 * 1024 * 2{
		oneThreadDownloadSize = length/2
	}
	filenames := []string{}
	for _, value := range tempFilenames {
		filenames = append(filenames, controller.DownloadFolder+"/"+value)
	}
	fileSuffix := controller.GetSuffix(contentType)
	filename := controller.DownloadFileName // 获取文件下载名
	controller.Filenames = filenames        //下载文件的切片列表
	rLock.Lock()
	controller.RangeStrs = rangeMaps        // 下载文件的Range范围
	rLock.Unlock()
	go controller.Run()                     // 开始下载文件
	proccessResult := controller.ResultProccess(len(rangeMaps))
	downloadResult := false // 定义下载结果标记
	if proccessResult == "SUCCESSED" {
		absoluteFilename := controller.DownloadFolder + "/" + filename + "." + fileSuffix
		downloadResult = controller.CombineFiles(filename + "." + fileSuffix)
		if downloadResult {
			newSuffix := GetFileType(GetBytesFile(absoluteFilename, 10))
			err = os.Rename(absoluteFilename, controller.DownloadFolder+"/"+filename+"."+newSuffix)
			if err != nil {
				downloadResult = false
				fmt.Println("Combine file successed, Rename file failed " + absoluteFilename)
			} else {
				fmt.Println("Combine file successed, rename successed, new file name is -> " + controller.DownloadFolder + "/" + filename + "." + newSuffix)
			}
		} else {
			fmt.Println("Failed to downlaod file.")
		}
	} else {
		fmt.Println("Failed to download file. Reason -> " + proccessResult)
		downloadResult = false
	}
	return downloadResult
}

func (controller *DownloadThreadController) CombineFiles(filename string) bool {
	os.Remove(controller.DownloadFolder + "/" + filename)
	goalFile, err := os.OpenFile(controller.DownloadFolder+"/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file ")
		return false
	}

	// 正确的话应按照初始计算的文件名顺序合并，并且无缺失
	for _, value := range controller.Filenames {
		retryTime := 3
		tempFileBytes := []byte{}
		for retryTime > 0 {
			tempFileBytes = ReadFile(value)
			time.Sleep(100) // 休眠100毫秒，看看是不是文件加载错误
			if tempFileBytes != nil {
				break
			}
			retryTime = retryTime - 1
		}
		goalFile.Write(tempFileBytes)
		os.Remove(value)
	}
	goalFile.Close()
	return true
}

func ReadFile(filename string) []byte {
	tempFile, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file " + filename)
		return nil
	}
	tempFileBytes, err := ioutil.ReadAll(tempFile)
	if err != nil {
		fmt.Println("Failed to read file data " + filename)
		return nil
	}
	tempFile.Close()
	return tempFileBytes
}

func TestDownload(urlStr string, perThreadSize int) (int, map[int]string, []string, string, error) {
	// 尝试连接目标资源，目标资源是否可以使用多线程下载
	length := 0
	rangeMaps := make(map[int]string)
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "www.downxia.com")
	req.Header.Add("Referer", urlStr)
	req.Header.Add("Range", "bytes=0-1") // 测试下载部分内容
	res, err := http.DefaultClient.Do(req)
	contentType := ""
	rangeIndex := 1
	filenames := []string{}
	if err != nil {
		rangeMaps[rangeIndex] = urlStr
		return length, rangeMaps, filenames, contentType, errors.New("Failed to connet " + urlStr)
	}
	if res.StatusCode != 206 {
		rangeMaps[rangeIndex] = urlStr
		return length, rangeMaps, filenames, contentType, errors.New("Http status is not equal to 206!")
	}
	// 206表示响应成功，仅仅返回部分内容
	contentLength := res.Header.Get("Content-Range")
	contentType = res.Header.Get("Content-Type")
	total_length, err := strconv.Atoi(strings.Split(contentLength, "/")[1])
	if err != nil {
		return length, rangeMaps, filenames, contentType, errors.New("Can't calculate the content-length form server " + urlStr)
	}
	now_length := 0 // 记录byte偏移量
	if total_length<1024*1024*2 {
		perThreadSize = total_length / 2
	}
	for {
		if now_length >= total_length {
			break
		}
		var tempRangeStr string // 记录临时文件名
		if now_length+perThreadSize >= total_length {
			tempRangeStr = "bytes=" + strconv.Itoa(now_length) + "-" + strconv.Itoa(total_length-1)
			now_length = total_length
		} else {
			tempRangeStr = "bytes=" + strconv.Itoa(now_length) + "-" + strconv.Itoa(now_length+perThreadSize-1)
			now_length = now_length + perThreadSize
		}
		rangeMaps[rangeIndex] = tempRangeStr
		filenames = append(filenames, tempRangeStr)
		rangeIndex = rangeIndex + 1
	}
	return total_length, rangeMaps, filenames, contentType, nil
}

func (controller *DownloadThreadController) Init() {
	taskQueue := make(chan DownloadTask, controller.ThreadCount)
	taskCount := make(chan int, controller.ThreadCount+1)
	exit := make(chan int)
	downloadResultInfoChan := make(chan DownloadFileParams)
	controller.TaskQueue = taskQueue
	controller.TaskCount = taskCount
	controller.Exit = exit
	controller.DownloadResultInfoChan = downloadResultInfoChan
	controller.WaitGroup = sync.WaitGroup{}
	rLock = sync.Mutex{}
	controller.RangeStrs = make(map[int]string)
	SafeMkdir(controller.DownloadFolder)
}

func DownloadFile(url,fileName string) bool{
	fmt.Println(url+"|||"+fileName)
	oneThreadDownloadSize := 1024 * 1024 * 2 // 一个线程下载文件的大小
	controller := DownloadThreadController{ThreadCount: thread, FileUrl: url, DownloadFolder: folder, DownloadFileName: fileName}
	return controller.Download(oneThreadDownloadSize)
}
