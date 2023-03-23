package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gotest/lib/response"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 防跨域
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Static("/static", "./static")
	// 加载模板
	r.LoadHTMLGlob("views/*")
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
		//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1")
	})
	r.GET("/back1", Back1)
	r.GET("/back2", Back2)
	r.GET("/upload", Upload1)
	r.GET("/download", Downlaod)
	r.POST("/UploadServlet", UploadServlet)
	r.Run()
	//监听端口默认为8080
	//r.RunTLS(":7272", "/www/server/panel/vhost/cert/file2.uoften.com/fullchain.pem", "/www/server/panel/vhost/cert/file2.uoften.com/privkey.pem")
}

func UploadServlet(ctx *gin.Context) {
	retData := make(map[string]interface{})
	forms, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("error", err)
	}
	files := forms.File["file"]
	for _, v := range files {
		if v.Size>1024*1024*50 {
			response.Fail(ctx,"文件暂时不能大于50M",retData)
			break
		}
		retData["src"] = v.Filename
		retData["size"] = v.Size
		v.Filename = strings.Replace(v.Filename, " ", "_", -1)
		filePath := time.Now().Format("20060102")
		retData["src"] = v.Filename
		retData["size"] = v.Size
		retData["path"] = filePath
		fileDir := path.Join("./file/", filePath)
		// 目录不存在则创建
		_, err := os.Stat(fileDir)
		if err != nil {
			if os.IsNotExist(err) {
				err := os.MkdirAll(fileDir, os.ModePerm)
				if err != nil {
					return
				}
			}
		}
		if err := ctx.SaveUploadedFile(v, fmt.Sprintf("%s%s", "./file/"+filePath+"/", v.Filename)); err != nil {
			response.Fail(ctx,"上传失败",retData)
			break
		}else{
			response.Success(ctx,"上传成功",retData)
			break
		}
	}
}

func Upload1(c *gin.Context) {
	// 定义模板变量
	data := make(map[string]interface{}, 1)

	c.HTML(http.StatusOK, "views/page.html", data)
}
func Back1(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	c.HTML(http.StatusOK, "views/back1.html", data)
}
func Back2(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	c.HTML(http.StatusOK, "views/back2.html", data)
}

func Downlaod(ctx *gin.Context) {
	token := ctx.Query("token")

	tokenStr:=bDecode(token,"uoasdasdasfteasdasdn")

	tokenStr, _ = url.QueryUnescape(tokenStr)
	start := strings.Index(tokenStr, "{")
	end := strings.Index(tokenStr, "}")
	jsonStr := tokenStr[start:end+1]

	fileInfo := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &fileInfo)
	if err != nil {
		ctx.String(http.StatusOK, "文件token失效")
		return
	}
	fmt.Println(fileInfo)
	exptime,_:=strconv.ParseInt(fileInfo["exptime"],0,64)
	if time.Now().Unix()>exptime{
		//ctx.String(http.StatusOK, "连接已过期,请回到下载页面重新下载")
		ctx.Redirect(http.StatusTemporaryRedirect, "/back1")
		return
	}
	file_name,err:= base64.StdEncoding.DecodeString(fileInfo["filename"])
	if err != nil {
		ctx.String(http.StatusOK, "文件名失效")
		return
	}
	filePath := string(file_name)
	//file_path := time.Now().Format("20060102")
	file_path := fileInfo["path"]
	//打开文件
	file, err := os.Open("./file/" + file_path + "/" + filePath)
	if err != nil {
		//ctx.String(http.StatusOK, "找不到文件,请回到下载页面重新下载")
		ctx.Redirect(http.StatusTemporaryRedirect, "/back2")
		return
	}
	defer file.Close()
	//获取文件的名称
	fileName := path.Base(filePath)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; url="+"./file/" + fileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.FileAttachment("./file/" + file_path + "/" + filePath,fileName)
}

// CleanInvalidTxt 文本无效字符清理
func CleanInvalidTxt(in []byte) []byte {
	var result []byte
	var last byte

	result = make([]byte, len(in))
	size := 0
	for _, t := range in {
		if (t <= 31 && t != 10 && t != 9 && t != 27 && t != 7) || t == 127 {
			continue
		} else {
			if last == 10 {
				if t == 10 || t == 32 {
					continue
				}
			}
			result[size] = t
			last = t
			size++
		}
	}
	return result[:size]
}

// Ascii字符清理
func trimHiddenCharacter(originStr string) string {
	srcRunes := []rune(originStr)
	dstRunes := make([]rune, 0, len(srcRunes))
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}
	return string(dstRunes)
}

func bEncode(jsonMap map[string]string,key string) string {
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

func bDecode(baseEncodedStr string,key string) string {
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

func checkUrl(u string) bool{
	_, err := url.ParseRequestURI(u)
	if err != nil {
		fmt.Println(err)
		return false
	}
	url, err := url.Parse(u)
	if err != nil || url.Scheme == "" || url.Host == "" {
		fmt.Println(err)
		return false
	}
	return true
}