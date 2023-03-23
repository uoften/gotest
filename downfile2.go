package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "http://example.com/file.txt"
	err := DownloadFile("./example.txt", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	contentType = resp.Header.Get("Content-Type")

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if contentType == "application/octet-stream" {
		// Create the file
		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return err
	} else {
		fmt.Println("Requested URL is not downloadable")
		return nil
	}
}