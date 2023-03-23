package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func main() {
	f, err := os.Open("huawei.conf")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer f.Close()

	readTxt(f)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("content:", content)
}

func readTxt(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)

	l := make([]string, 0, 64)

	// 按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		fmt.Printf("\n%#v",strings.Trim(string(line), " "))
		//l = append(l, strings.Trim(string(line), " "))
	}

	return l, nil
}