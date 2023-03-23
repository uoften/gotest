package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var configStr string
	// 获取配置文件
	err := readBlock("./log.txt", &configStr)
	if err!=nil {
		fmt.Errorf("%v",err)
	}
	//err := readBlock("./log.txt", &configStr)
	//if err!=nil {
	//	fmt.Errorf("%v",err)
	//}
	elapsed := time.Since(start)
	fmt.Println("执行完成耗时：", elapsed)
}
//流式读取：根据分隔符逐行读取，有分隔符的场景
func ReadBuf(fileName string, configStr *string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		//line, err := buf.ReadString('\n')
		line, _,err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
		//fmt.Println("--------------------------------------------------------------------")
		*configStr = string(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
//分片读取：每次读取一片，没有分隔符的场景
func readBlock(filePath string, configStr *string) error{
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// 设置每次读取字节数
	buffer := make([]byte, 1024)
	for {
		n, err := f.Read(buffer)
		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}
		// 如下代码打印出每次读取的文件块(字节数)
		line := string(buffer[:n])
		line = strings.TrimSpace(line)
		fmt.Println(line)
		//fmt.Println("--------------------------------------------------------------------")
		*configStr = line
	}
	return nil
}

func writeMain() {
	fileName := "./log.txt"
	strTest := "666666\n"
	var f *os.File
	var err error

	if CheckFileExist(fileName) { //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			fmt.Println("file open fail", err)
			return
		}
	} else { //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
	}
	//将文件写进去
	for i := 1; i <= 1; i++ {
		_, err1 := io.WriteString(f, strTest)
		if err1 != nil {
			fmt.Println("write error", err1)
			return
		}
	}
}
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}