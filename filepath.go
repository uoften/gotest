package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir,filename:=filepath.Split("https://studygolang.com/dl/golang/go1.13.4.linux-amd64.tar.gz")
	fmt.Println(dir)
	fmt.Println(filename)
}
