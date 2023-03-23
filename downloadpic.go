package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"gotest/library"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	//src := "https://upload-images.jianshu.io/upload_images/15604519-7648c594d4b3972c.png?imageMogr2/auto-orient/strip|imageView2/2/w/398/format/webp"
	src := "https://img-va.myshopline.com/image/store/3400003311/1627028228094/294ff9272bc94dc99aaabf87982e21f2.webp?w=640&h=853"
	resp, body, errs := gorequest.New().Set("referer", src).Timeout(30 * time.Second).Get(src).EndBytes()
	if errs == nil {
		//处理
		contentType := resp.Header.Get("content-type")
		fmt.Println(contentType)
		if contentType == "image/jpeg" || contentType == "image/jpg" || contentType == "image/png" || contentType == "image/gif" {
			//获取宽高
			bufFile := &bytes.Buffer{}
			bufFile.Write(body)
			img, imgType, err := image.Decode(bufFile)
			if err != nil {
				//无法获取图片尺寸
				fmt.Println("无法获取图片尺寸")
			}
			bufFile.Reset()
			bufFile.Write(body)
			imgType = strings.ToLower(imgType)
			width := img.Bounds().Dx()
			height := img.Bounds().Dy()
			fmt.Println(height)
			//只允许上传jpg,jpeg,gif,png
			if imgType != "jpg" && imgType != "jpeg" && imgType != "gif" && imgType != "png" {
				fmt.Printf("不支持的图片格式：%s。", imgType)
			}
			if imgType == "jpeg" {
				imgType = "jpg"
			}
			//获取文件的MD5，检查数据库是否已经存在，存在则不用重复上传
			md5hash := md5.New()
			_, err = io.Copy(md5hash, bufFile)
			if err != nil {
				fmt.Println("检查文件错误")
			}
			bufFile.Reset()
			bufFile.Write(body)
			md5Str := hex.EncodeToString(md5hash.Sum(nil))


			bufFile.Reset()
			bufFile.Write(body)
			//如果图片宽度大于800，自动压缩到800, gif 不能处理
			resizeWidth := 800
			if resizeWidth == 0 {
				//默认800
				resizeWidth = 800
			}
			buff := &bytes.Buffer{}

			if width > resizeWidth && imgType != "gif" {
				newImg := library.Resize(img, resizeWidth, 0)
				width = newImg.Bounds().Dx()
				height = newImg.Bounds().Dy()
				if imgType == "jpg" {
					// 保存裁剪的图片
					_ = jpeg.Encode(buff, newImg, nil)
				} else if imgType == "png" {
					// 保存裁剪的图片
					_ = png.Encode(buff, newImg)
				}
			} else {
				_, _ = io.Copy(buff, bufFile)
				fmt.Println("文件已存在或者不支持的格式")
			}

			tmpName := md5Str[8:24] + "." + imgType
			filePath := strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/"

			//将文件写入本地
			basePath := "./file/"
			//先判断文件夹是否存在，不存在就先创建
			_, err = os.Stat(basePath + filePath)
			if err != nil && os.IsNotExist(err) {
				err = os.MkdirAll(basePath+filePath, os.ModePerm)
				if err != nil {
					fmt.Println("无法创建目录")
				}
			}

			originFile, err := os.OpenFile(basePath+filePath+tmpName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				//无法创建
				fmt.Println("无法创建文件")
			}

			defer originFile.Close()

			_, err = io.Copy(originFile, buff)
			if err != nil {
				//文件写入失败
				fmt.Println("文件写入失败")
			}

			//生成宽度为250的缩略图
			thumbName := "thumb_" + tmpName

			newImg := library.ThumbnailCrop(0, 0, img, 0)
			if imgType == "jpg" {
				_ = jpeg.Encode(buff, newImg, nil)
			} else if imgType == "png" {
				_ = png.Encode(buff, newImg)
			} else if imgType == "gif" {
				_ = gif.Encode(buff, newImg, nil)
			}

			thumbFile, err := os.OpenFile(basePath+filePath+thumbName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				//无法创建
				fmt.Println("无法创建")
			}

			defer thumbFile.Close()

			_, err = io.Copy(thumbFile, buff)
			if err != nil {
				//文件写入失败
				fmt.Println("文件写入失败")
			}
			fmt.Println("图片下载成功")
		} else {
			fmt.Println("不支持的图片格式")
		}
	}

}