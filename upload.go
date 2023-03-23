package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
const (
	uploadPath = "./Files/"
)
func main() {
	http.HandleFunc("/upload", uploadHandle)
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/Files/", http.StripPrefix("/Files", fs))
	log.Fatal(http.ListenAndServe(":8037", nil))
}
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	filePath := uploadPath + head.Filename
	fW, err := os.Create(filePath)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	io.WriteString(w, "save to "+filePath)
}