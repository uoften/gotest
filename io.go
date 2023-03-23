package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("at.txt")
	defer f.Close()
	f.Write([]byte("Go是一种令人愉悦的编程语言")) //写入字节流
	f.Seek(0, os.SEEK_SET)            //将指针重置
	p := make([]byte, 2) // 读取 2 byte( len(buf)=2 )
	if _, err := f.Read(p); err != nil {
		log.Fatal("[F]", err)
	}
	fmt.Printf("读取字符 \"%s\", 长度为 %d byte\n", p, len(p))

	p = make([]byte, 50)
	if _, err := f.Read(p); err != nil {
		if err != io.EOF { //忽略 EOF 错误
			log.Fatal("[F]", err)
		}
	}
	fmt.Printf("读取字符 \"%s\", 长度为 %d byte\n", p, len(p))
}
